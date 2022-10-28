package helper

import (
	"os"

	"github.com/pusher/pusher-http-go"
)

func NewPusherClient() pusher.Client {
	pusherClient := pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
		Secure:  true,
	}

	return pusherClient
}
