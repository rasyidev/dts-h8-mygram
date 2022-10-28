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
)

func PostPhoto(ctx *gin.Context) {
	db := database.NewDB()
	id, _ := uuid.NewV4()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"]

	requestBody := models.PhotoPostRequest{}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		log.Fatal(err.Error())
	}
	photo := models.Photo{ID: id.String(), Title: requestBody.Title, Caption: requestBody.Caption, PhotoUrl: requestBody.PhotoUrl, CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: userID.(string)}

	// validate the request body
	if err := validator.New().Struct(requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Validation error: " + err.Error(),
		})
		return
	}

	if err := db.Debug().Create(&photo).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success creating new user",
		"data":    photo.ToPhotoPostedResponse(),
	})
}

func GetPhotos(ctx *gin.Context) {
	db := database.NewDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"]
	photos := []models.GetAllPhotosReponse{}
	user := models.UserSecured{}

	if err := db.Model(models.Photo{}).Where("user_id=?", userID).Find(&photos).Error; err != nil {
		log.Println("You haven't post any photo")
		ctx.AbortWithStatusJSON(http.StatusNoContent, gin.H{
			"status":  http.StatusNoContent,
			"message": "No Content" + err.Error(),
		})
		return
	}

	if err := db.Raw("SELECT email, username FROM users WHERE id=?", userID).Scan(&user).Error; err != nil {
		log.Println("Eror: ", err.Error())
	}

	for i := range photos {
		photos[i].User = user
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success getting all of your photos",
		"data":    photos,
	})

}

func UpdatePhoto(ctx *gin.Context) {
	db := database.NewDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"]
	requestBody := models.PhotoPostRequest{}
	photo := models.Photo{}
	photoID := ctx.Param("photoID")
	fmt.Println("photoID: ", photoID)

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if err := db.Where("user_id=?", userID).Find(&photo).Limit(1).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "You haven't post any photo yet :" + err.Error(),
		})
		return
	}

	photo.UpdatedAt = time.Now()

	if err := db.Model(&models.Photo{}).Where("id=?", photoID).Updates(requestBody).Scan(&photo).Error; err != nil {
		log.Println("Error updating user: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error updating user " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success updating your photos",
		"data":    photo.ToUpdatedPhotosResponse(),
	})
}

func DeletePhoto(ctx *gin.Context) {
	db := database.NewDB()
	photoID := ctx.Param("photoID")
	contentType := helper.GetContentType(ctx)
	fmt.Println("contentType: ", contentType)

	if err := db.Where("id=?", photoID).Delete(&models.Photo{}).Error; err != nil {
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
