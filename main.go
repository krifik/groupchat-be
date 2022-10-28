package main

import (
	"github.com/krifik/groupchat-be/app"
	"github.com/krifik/groupchat-be/exception"
)

func main() {
	appl := app.InitializedApp()
	// app.InitializedPusher()
	// Start App
	err := appl.Listen(":3000")
	exception.PanicIfNeeded(err)
}
