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

func (u *Users) GetUsers(c *gin.Context) {
	var users []model.User

	db := u.app.Database.GetDb()

	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Print(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &users,
	})

}

func (u *Users) GetUserByUserId(c *gin.Context) {

	var user model.User
	userId := c.Params.ByName("userId")

	db := u.app.Database.GetDb()

	if err := db.Where("id=?", userId).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Print(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
	})

}

func (u *Users) PostUser(c *gin.Context) {
	var user model.User
	var userLogin model.UserLogin
	var body model.CreateUserBody
	database := u.app.Database.GetDb()
	transaction := database.Begin()

	// body, err := json.Marshal(body)
	// if err != nil {
	// }
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.Where("email=?", body.Email).First(&userLogin).Error
	if err == nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": "user email already exist"})
		return
	}

	err = database.Model(&user).Create(&model.User{Nickname: body.Nickname, Birthday: body.Birthday, PhoneCountryCode: body.PhoneCountryCode, PhoneNumber: body.PhoneNumber, Gender: body.Gender, Region: body.Region}).Error
	if err != nil {
		transaction.Rollback()
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	h := hmac.New(sha256.New, []byte(viper.GetString("token_manager.secret")))
	h.Write([]byte(body.Password))
	password := hex.EncodeToString(h.Sum(nil))
	err = database.Model(&userLogin).Create(&model.UserLogin{Email: body.Email, Password: string(password[:])}).Error
	if err != nil {
		transaction.Rollback()
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	transaction.Commit()

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
	})
}
func (u *Users) UpdateUserByUserId(c *gin.Context) {
	var user model.User
	var body model.UpdateUserBody
	database := u.app.Database.GetDb()
	userId := c.Params.ByName("userId")

	if err := database.Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.Model(&user).Update(body).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
	})

}
func (u *Users) DeleteUserByUserId(c *gin.Context) {
	var user model.User
	// var userRaw model.UpdateUserBody;
	// userRaw, err:= c.GetRawData()
	// if err != nil {}

	// log.Print(string(userRaw))
	database := u.app.Database.GetDb()

	result := database.Find(&user)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}
