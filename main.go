package main

import (
	"mangojek-backend/app"
	"mangojek-backend/exception"
)

func main() {
	appl := app.InitializedApp()
	// app.InitializedPusher()
	// Start App
	err := appl.Listen(":3000")
	exception.PanicIfNeeded(err)
}
