package account

import (
	"gameObjectBackend/app"
	"gameObjectBackend/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Accounts struct {
	app    *app.App
	server *app.HTTPServer
	Router *gin.RouterGroup
}

func (a *Accounts) Init(server *app.HTTPServer) {

	a.server = server
	a.app = server.App
	a.Router = server.GetEngine().Group("/api/v1/accounts")
	a.Router.POST("/signin", a.SignIn)
}

// @tags accounts
// @Summary get user token
// @Description 回傳 user token
// @Accept  json
// @Produce  json
// @Param body body model.PostSigninBody true "參數"
// @Success 200 {object} model.UserSignInResponseObject string "ok"
// @Failure 403 {object} string "err.Error()"
// @Router /api/v1/accounts/signin [post]
func (a *Accounts) SignIn(c *gin.Context) {
	var userLogin model.UserLogin
	var body model.PostSigninBody

	db := a.app.Database.GetDb()

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Where("email=?", body.Email).First(&userLogin).Error
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &userLogin,
	})

}
