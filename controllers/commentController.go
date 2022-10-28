package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	"github.com/rasyidev/dts-h8-mygram/database"
	"github.com/rasyidev/dts-h8-mygram/models"
)

func PostComment(ctx *gin.Context) {
	db := database.NewDB()
	id, _ := uuid.NewV4()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"]

	requestBody := models.PostCommentRequest{}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		log.Fatal(err.Error())
	}
	comment := models.Comment{
		ID:        id.String(),
		CreatedAt: time.Now(),
		Message:   requestBody.Message,
		PhotoID:   requestBody.PhotoID,
		UpdatedAt: time.Now(),
		UserID:    userID.(string),
	}

	// validate the request body
	if err := validator.New().Struct(requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Validation error: " + err.Error(),
		})
		return
	}

	if err := db.Debug().Create(&comment).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success posting a comment",
		"data":    comment.ToPostCommentResponse(),
	})
}

func GetComments(ctx *gin.Context) {
	db := database.NewDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"]
	comments := []models.PostCommentResponse{}
	user := models.UserCommentResponse{}
	photo := models.PhotoCommentResponse{}

	if err := db.Model(models.Comment{}).Where("user_id=?", userID).Find(&comments).Error; err != nil {
		log.Println("You haven't post any photo")
		ctx.AbortWithStatusJSON(http.StatusNoContent, gin.H{
			"status":  http.StatusNoContent,
			"message": "No Content" + err.Error(),
		})
		return
	}

	if err := db.Raw("SELECT email, username FROM users WHERE id=?", userID).Scan(&user).Error; err != nil {
		log.Println("Error: ", err.Error())
	}

	for i := range comments {
		comments[i].User = user
	}

	for i := range comments {
		comments[i].Photo = photo
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success getting all of your photos",
		"data":    comments,
	})

}
