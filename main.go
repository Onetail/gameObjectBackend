package main

import (
	"gameObjectBackend/app"
	"gameObjectBackend/app/controller"
)

// @title GameObjectBackend
// @version 1.0.0
// @description Gin swagger api docs

// @contact.name WayneChu
// @contact.url https://github.com/Onetail

// @license.name Nginx
// @license.url https://www.nginx.com/

// @host localhost:7011
// schemes http
func main() {

	// Initializing application
	app := app.App{}
	app.Init()

	user := controller.Users{}
	user.Init(app.HTTPServer)

	// Start application
	app.Run()
}
