package controller

import (
	"gameObjectBackend/app"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Users struct {
	app    *app.App
	server *app.HTTPServer
	Router *gin.RouterGroup
}


func (u *Users) Init(server *app.HTTPServer) {

	u.server = server
	u.app = server.App
	u.Router = server.GetEngine().Group("/api/v1/users")
	
	
	u.Router.GET("/list", u.GetUsers)
}

func (u *Users) GetUsers(c *gin.Context) {


	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}