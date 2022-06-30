package server

import (
	"golangSimpleCrud/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// TODO: Add AuthMiddleware here
	// router.Use(middleware...)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.POST("/", user.Create)
			userGroup.GET("/", user.GetAll)
			userGroup.GET("/:id", user.GetOne)
			userGroup.PATCH("/:id", user.Update)
			userGroup.DELETE("/:id", user.Delete)
		}
	}

	return router
}
