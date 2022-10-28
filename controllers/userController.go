package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	"github.com/rasyidev/dts-h8-mygram/database"
	"github.com/rasyidev/dts-h8-mygram/helper"
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
	user := models.User{}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		log.Fatal("Error binding to user data: ", err.Error())
	}

	// encryptedPass, _ := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 1)
	// fmt.Println("encryptedPass: ", string(encryptedPass))

	if err := db.Model(&user).Where("email=?", requestBody.Email).Limit(1).Scan(&user).Error; err != nil {
		log.Println("error binding user data")
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
			log.Println("Email or password not match: ", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Email or password not match " + err.Error(),
			})
			return
		}

		fmt.Println("email and password match: ", user.Email, user.Password)
		token := helper.GenerateToken(user.ID, user.Email)
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

// UpdateUser update email and username of the logged-in user
func UpdateUser(ctx *gin.Context) {
	db := database.NewDB()
	requestBody := models.User{UpdatedAt: time.Now()}
	user := models.User{}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	fmt.Println("userData: ", userData)
	contentType := helper.GetContentType(ctx)
	fmt.Println("contentType: ", contentType)

	ctx.ShouldBindJSON(&requestBody)

	if err := db.Model(&models.User{}).Where("id=?", userData["id"]).Updates(requestBody).Scan(&user).Error; err != nil {
		log.Println("Error updating user: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error updating user " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success updating data",
		"data":    user.ToUserUpdatedResponse(),
	})
}

// UpdateUser update email and username of the logged-in user
func DeleteUser(ctx *gin.Context) {
	db := database.NewDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	fmt.Println("userData: ", userData)
	contentType := helper.GetContentType(ctx)
	fmt.Println("contentType: ", contentType)

	if err := db.Where("id=?", userData["id"]).Delete(&models.User{}).Error; err != nil {
		log.Println("Error deleting user: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error deleting user " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Your account has been successfully deleted",
	})
}
