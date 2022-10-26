package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rasyidev/dts-h8-mygram/database"
	"github.com/rasyidev/dts-h8-mygram/models"
)

func CreateUser(ctx gin.Context) {
	db := database.NewDB()
	user := models.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Create(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
}
