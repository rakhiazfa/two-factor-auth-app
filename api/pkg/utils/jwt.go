package utils

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func CreateRefreshToken(claims jwt.MapClaims) (string, error) {
	privateKey, err := ConvertToPrivateKey(viper.GetString("refresh_token.private_key"))
	if err != nil {
		return "", err
	}

	exp := viper.GetInt("refresh_token.expires_in")
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(exp)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateAccessToken(claims jwt.MapClaims) (string, error) {
	privateKey, err := ConvertToPrivateKey(viper.GetString("access_token.private_key"))
	if err != nil {
		return "", err
	}

	exp := viper.GetInt("access_token.expires_in")
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(exp)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyRefreshToken(tokenString string) (*jwt.MapClaims, error) {
	publicKey, err := ConvertToPublicKey(viper.GetString("refresh_token.public_key"))
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return &claims, nil
}

func VerifyAccessToken(tokenString string) (*jwt.MapClaims, error) {
	publicKey, err := ConvertToPublicKey(viper.GetString("access_token.public_key"))
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return &claims, nil
}

func ExtractJwtExp(tokenString string) (*time.Time, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("exp claim not found or is invalid")
	}

	expTime := time.Unix(int64(exp), 0)

	return &expTime, nil
}

func ConvertToPrivateKey(key string) (*rsa.PrivateKey, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(key))
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func ConvertToPublicKey(key string) (*rsa.PublicKey, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(key))
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
