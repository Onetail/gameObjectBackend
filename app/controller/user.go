package controller

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gameObjectBackend/app"
	"gameObjectBackend/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	u.Router.GET("/", u.GetUsers)
	u.Router.GET("/:userId", u.GetUserByUserId)
	u.Router.POST("/", u.PostUser)
	u.Router.PATCH("/:userId", u.UpdateUserByUserId)
	u.Router.DELETE("/:userId", u.DeleteUserByUserId)

}

// @tags users
// @Summary Get user list
// @Description 取得 user 列表
// @Accept  json
// @Produce  json
// @Success 200 {object} model.UserListResponseObject string "ok"
// @Failure 403 {object} string "err.Error()"
// @Router /api/v1/users/ [get]
func (u *Users) GetUsers(c *gin.Context) {
	var users []model.User

	db := u.app.Database.GetDb()

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &users,
	})

}

// @tags users
// @Summary Get user by userId
// @Description 取得單一 user
// @Accept  json
// @Produce  json
// @Param userId path string true "user id"
// @Success 200 {object} model.UserResponseObject string "ok"
// @Failure 404 {object}  string "record not found"
// @Router /api/v1/users/:userId [get]
func (u *Users) GetUserByUserId(c *gin.Context) {

	var user model.User
	userId := c.Params.ByName("userId")

	db := u.app.Database.GetDb()

	if err := db.Where("id=?", userId).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
	})

}

// @tags users
// @Summary create user
// @Description 新增 user
// @Accept  json
// @Produce  json
// @Param body body model.CreateUserBody true "參數"
// @Success 200 {object} model.User string "ok"
// @Router /api/v1/users/ [post]
func (u *Users) PostUser(c *gin.Context) {
	var user model.User
	var userLogin model.UserLogin
	var body model.CreateUserBody
	database := u.app.Database.GetDb()
	transaction := database.Begin()

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.Where("email=?", body.Email).First(&userLogin).Error
	if err == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "user email already exist"})
		return
	}

	err = database.Model(&user).Create(&model.User{Nickname: body.Nickname, Birthday: body.Birthday, PhoneCountryCode: body.PhoneCountryCode, PhoneNumber: body.PhoneNumber, Gender: body.Gender, Region: body.Region}).Error
	if err != nil {
		transaction.Rollback()
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	}

	h := hmac.New(sha256.New, []byte(viper.GetString("token_manager.secret")))
	h.Write([]byte(body.Password))
	password := hex.EncodeToString(h.Sum(nil))
	err = database.Model(&userLogin).Create(&model.UserLogin{Email: body.Email, Password: string(password[:])}).Error
	if err != nil {
		transaction.Rollback()
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	}

	transaction.Commit()

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
	})
}

// @tags users
// @Summary update user
// @Description 更新 user
// @Accept  json
// @Produce  json
// @Param userId path string true "user id"
// @Param body body model.UpdateUserBody true "參數"
// @Success 200 {object} model.RowsAffectedModel string "ok"
// @Router /api/v1/users/:userId [patch]
func (u *Users) UpdateUserByUserId(c *gin.Context) {
	var user model.User
	var body model.UpdateUserBody
	database := u.app.Database.GetDb()
	userId := c.Params.ByName("userId")

	err := database.Where("id = ?", userId).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateResult := database.Model(&user).Update(body)
	if updateResult.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": updateResult.RowsAffected,
	})

}

// @tags users
// @Summary delete user
// @Description 刪除 user
// @Accept  json
// @Produce  json
// @Param userId path string true "user id"
// @Success 200 {object} model.RowsAffectedModel "success"
// @Router /api/v1/users/:userId [delete]
func (u *Users) DeleteUserByUserId(c *gin.Context) {
	var user model.User
	database := u.app.Database.GetDb()
	userId := c.Params.ByName("userId")

	err := database.Where("id = ?", userId).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	result := database.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"data": result.RowsAffected,
	})

}
