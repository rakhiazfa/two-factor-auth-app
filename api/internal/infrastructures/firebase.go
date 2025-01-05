package infrastructures

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

func NewFirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile(viper.GetString("firebase.service_account_file"))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("Failed to initialize firebase application: ", err)
	}

	return app
}
