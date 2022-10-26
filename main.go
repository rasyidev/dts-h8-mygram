package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rasyidev/dts-h8-mygram/controllers"
)

func main() {
	r := gin.Default()
	r.POST("api/v1/users/register", controllers.CreateUser)

	r.Run(":9090")
}
