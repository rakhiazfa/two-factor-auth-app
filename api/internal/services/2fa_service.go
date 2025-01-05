package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/pusher/pusher-http-go/v5"
	"github.com/rakhiazfa/gin-boilerplate/constants"
	"github.com/rakhiazfa/gin-boilerplate/internal/dtos"
	"github.com/rakhiazfa/gin-boilerplate/internal/entities"
	"github.com/rakhiazfa/gin-boilerplate/internal/repositories"
	"github.com/rakhiazfa/gin-boilerplate/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type TwoFactorAuthService struct {
	db                                  *gorm.DB
	validator                           *utils.Validator
	firebase                            *firebase.App
	pusher                              *pusher.Client
	twoFactorAuthSessionRepository      *repositories.TwoFactorAuthSessionRepository
	twoFactorAuthNumberOptionRepository *repositories.TwoFactorAuthNumberOptionRepository
}

func NewTwoFactorAuthService(
	db *gorm.DB,
	validator *utils.Validator,
	firebase *firebase.App,
	pusher *pusher.Client,
	twoFactorAuthSessionRepository *repositories.TwoFactorAuthSessionRepository,
	twoFactorAuthNumberOptionRepository *repositories.TwoFactorAuthNumberOptionRepository,
) *TwoFactorAuthService {
	return &TwoFactorAuthService{
		db:                                  db,
		validator:                           validator,
		firebase:                            firebase,
		pusher:                              pusher,
		twoFactorAuthSessionRepository:      twoFactorAuthSessionRepository,
		twoFactorAuthNumberOptionRepository: twoFactorAuthNumberOptionRepository,
	}
}

func (s *TwoFactorAuthService) Create2FASession(ctx context.Context, req *dtos.Create2FASessionReq) (*entities.TwoFactorAuthSession, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	var twoFactorAuthSession entities.TwoFactorAuthSession

	if err := copier.Copy(&twoFactorAuthSession, req); err != nil {
		return nil, err
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return s.twoFactorAuthSessionRepository.WithTx(tx).Create(&twoFactorAuthSession)
	})
	if err != nil {
		return nil, err
	}

	return &twoFactorAuthSession, nil
}

func (s *TwoFactorAuthService) Create2FANumberOptions(ctx context.Context, twoFactorAuthSession *entities.TwoFactorAuthSession) (string, error) {
	correctNumber, numberOptions := s.generateNumberOptions()

	hash, err := bcrypt.GenerateFromPassword([]byte(strconv.Itoa(correctNumber)), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	twoFactorAuthSession.CorrectNumber = utils.ToPointer(string(hash))

	var twoFactorAuthNumberOptions []entities.TwoFactorAuthNumberOption

	for _, number := range numberOptions {
		encryptedNumber, err := utils.AESEncrypt(strconv.Itoa(number))
		if err != nil {
			return "", err
		}

		twoFactorAuthNumberOptions = append(twoFactorAuthNumberOptions, entities.TwoFactorAuthNumberOption{
			TwoFactorAuthSessionId: twoFactorAuthSession.ID,
			Number:                 encryptedNumber,
		})
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := s.twoFactorAuthSessionRepository.WithTx(tx).Save(twoFactorAuthSession)
		if err != nil {
			return err
		}

		err = s.twoFactorAuthNumberOptionRepository.WithTx(tx).CreateMany(&twoFactorAuthNumberOptions)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return strconv.Itoa(correctNumber), nil
}

func (s *TwoFactorAuthService) FindOneById(ctx context.Context, sessionId uuid.UUID, relations ...string) (*entities.TwoFactorAuthSession, error) {
	twoFactorAuthSession, err := s.twoFactorAuthSessionRepository.WithContext(ctx).FindOneById(sessionId, relations...)
	if err != nil {
		return nil, err
	}
	if twoFactorAuthSession == nil {
		return nil, utils.NewHttpError(404, "Session not found", nil)
	}

	return twoFactorAuthSession, nil
}

func (s *TwoFactorAuthService) SendOptionNumbers2FA(ctx context.Context, sessionId uuid.UUID) error {
	twoFactorAuthSession, err := s.FindOneById(ctx, sessionId, "TwoFactorAuthNumberOptions")
	if err != nil {
		return err
	}

	var targetDevices []string

	if err := s.db.WithContext(ctx).Raw(
		`SELECT ud.token FROM user_devices AS ud  WHERE ud.user_id = ? AND ud.id != ?`,
		twoFactorAuthSession.UserId,
		twoFactorAuthSession.UserDeviceId,
	).Find(&targetDevices).Error; err != nil {
		return err
	}

	for index, numberOption := range twoFactorAuthSession.TwoFactorAuthNumberOptions {
		decrypted, err := utils.AESDecrypt(numberOption.Number)
		if err != nil {
			return err
		}

		twoFactorAuthSession.TwoFactorAuthNumberOptions[index].Number = decrypted
	}

	numberOptions, err := json.Marshal(&twoFactorAuthSession.TwoFactorAuthNumberOptions)
	if err != nil {
		return err
	}

	client, err := s.firebase.Messaging(ctx)
	if err != nil {
		return err
	}

	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "2FA Verification",
			Body:  string(numberOptions),
		},
		Data: map[string]string{
			"numberOptions": string(numberOptions),
		},
		Tokens: targetDevices,
	}

	response, err := client.SendEachForMulticast(ctx, message)
	if err != nil {
		return err
	}

	log.Printf("[FIREBASE] %d messages were sent successfully\n", response.SuccessCount)

	return nil
}

func (s *TwoFactorAuthService) Verify2FANumber(ctx context.Context, authUserDevice *entities.UserDevice, sessionId uuid.UUID, req *dtos.Verify2FAOptionReq) error {
	if err := s.validator.Validate(req); err != nil {
		return err
	}

	twoFactorAuthSession, err := s.FindOneById(ctx, sessionId, "User")
	if err != nil {
		return err
	}

	if twoFactorAuthSession.Verified {
		return utils.NewHttpError(http.StatusForbidden, "session has been verified", nil)
	}

	selectedNumberOption, err := s.twoFactorAuthNumberOptionRepository.FindOneById(req.OptionId)
	if err != nil {
		return err
	}
	if selectedNumberOption == nil {
		return utils.NewHttpError(http.StatusBadRequest, "The option you selected is incorrect", nil)
	}

	decryptedSelectedNumber, err := utils.AESDecrypt(selectedNumberOption.Number)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*twoFactorAuthSession.CorrectNumber), []byte(decryptedSelectedNumber)); err != nil {
		return utils.NewHttpError(http.StatusBadRequest, "The option you selected is incorrect", nil)
	}

	twoFactorAuthSession.Verified = true
	twoFactorAuthSession.ExpiresAt = time.Now().Add(7 * 24 * time.Hour)
	twoFactorAuthSession.ApprovedBy = utils.ToPointer(authUserDevice.ID)

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return s.twoFactorAuthSessionRepository.WithTx(tx).Save(twoFactorAuthSession)
	})
	if err != nil {
		return err
	}

	refreshToken, err := utils.CreateRefreshToken(jwt.MapClaims{
		"sub": twoFactorAuthSession.ID,
	})
	if err != nil {
		return err
	}

	accessToken, err := utils.CreateAccessToken(jwt.MapClaims{
		"sub": twoFactorAuthSession.ID,
	})
	if err != nil {
		return err
	}

	s.pusher.Trigger(constants.AppChannel_Auth+"."+twoFactorAuthSession.ID.String(), constants.AppEvent_Verify2FA, map[string]string{
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
	})

	return nil
}

func (s *TwoFactorAuthService) generateNumberOptions() (selectedNumber int, numberOptions [3]int) {
	used := make(map[int]bool, 3)

	for i := 0; i < len(numberOptions); i++ {
		num := utils.RandRange(10, 99)

		for used[num] {
			num = utils.RandRange(10, 99)
		}

		numberOptions[i] = num
		used[num] = true
	}

	selectedNumber = numberOptions[utils.RandRange(0, len(numberOptions))]

	return selectedNumber, numberOptions
}
