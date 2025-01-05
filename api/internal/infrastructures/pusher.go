package infrastructures

import (
	"github.com/pusher/pusher-http-go/v5"
	"github.com/spf13/viper"
)

func NewPusherClient() *pusher.Client {
	pusherClient := &pusher.Client{
		AppID:   viper.GetString("pusher.app_id"),
		Key:     viper.GetString("pusher.key"),
		Secret:  viper.GetString("pusher.secret"),
		Cluster: viper.GetString("pusher.cluster"),
		Secure:  viper.GetBool("pusher.secure"),
	}

	return pusherClient
}
