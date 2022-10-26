package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
