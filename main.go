package main

import (
	"os"

	"github.com/krifik/groupchat-be/app"
	"github.com/krifik/groupchat-be/exception"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}
func main() {
	appl := app.InitializedApp()
	// app.InitializedPusher()
	// Start App
	err := appl.Listen(getPort())
	exception.PanicIfNeeded(err)
}
