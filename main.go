package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rasyidev/dts-h8-mygram/controllers"
	"github.com/rasyidev/dts-h8-mygram/middlewares"
)

func main() {
	r := gin.Default()

	userRouter := r.Group("/api/v1/users")
	{
		userRouter.POST("/register", controllers.CreateUser)
		userRouter.GET("/login", controllers.Login)

		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/update", controllers.UpdateUser)
		userRouter.DELETE("/delete", controllers.DeleteUser)
	}

	photoRouter := r.Group("/api/v1/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.PostPhoto)
		photoRouter.GET("/", controllers.GetPhotos)
		photoRouter.PUT("/:photoID", controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoID", controllers.DeletePhoto)
	}

	commentRouter := r.Group("/api/v1/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.PostComment)
		commentRouter.GET("/", controllers.GetComments)
		// commentRouter.PUT("/:photoID", controllers.UpdatePhoto)
		// commentRouter.DELETE("/:photoID", controllers.DeletePhoto)
	}

	r.Run(":9090")
}
