package controller

import (
	"github.com/gin-gonic/gin"
	"go-hello/helper"
	"go-hello/model"
	"net/http"
)

func Register(context *gin.Context) {
	var userRegister model.UserRegister

	if err := context.ShouldBind(&userRegister); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user := model.User{
		Username: userRegister.Username,
		Password: userRegister.Password,
		Fullname: userRegister.Fullname,
		Email:    userRegister.Email,
	}

	createdUser, err := user.Create()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": createdUser})
}

func Login(context *gin.Context) {
	var userLogin model.UserLogin

	if err := context.ShouldBind(&userLogin); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var user model.User
	user, err := model.FindUserByUsername(userLogin.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := user.ValidatePassword(userLogin.Password); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	jwt, err := helper.GenerateJWT(user)
	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
