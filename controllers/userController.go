package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	"github.com/rasyidev/dts-h8-mygram/database"
	"github.com/rasyidev/dts-h8-mygram/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx *gin.Context) {
	db := database.NewDB()
	id, _ := uuid.NewV4()

	user := models.User{ID: id.String(), CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Fatal(err.Error())
	}

	// validate the request body
	if err := validator.New().Struct(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Validation error: " + err.Error(),
		})
		return
	}

	// hashing the password
	result, err := bcrypt.GenerateFromPassword([]byte(user.Password), 1)
	if err != nil {
		log.Fatal("Error hasing password:", err.Error())
	}
	user.Password = string(result)

	if err := db.Create(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success creating new user",
		"data":    user.ToUserCreatedResponse(),
	})
}

func Login(ctx *gin.Context) {
	db := database.NewDB()
	requestBody := models.User{}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		log.Fatal("Error binding to user data: ", err.Error())
	}

	encryptedPass, _ := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 1)

	if err := db.Where("email=? AND password=?", requestBody.Email, string(encryptedPass)).Limit(1).Error; err != nil {
		log.Fatal("Email not found: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Email not found: " + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"token": "abcdefghijklmnopqrstuvwxyz",
		})
	}
}
