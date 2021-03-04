package main

import (
	"gameObjectBackend/app"
	"gameObjectBackend/app/controller"
)

func main() {

	// Initializing application
	app := app.App{}
	app.Init()

	user := controller.Users{}
	user.Init(app.HTTPServer)

	// Start application
	app.Run()
}
