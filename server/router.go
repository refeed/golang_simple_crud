package server

import (
	"golangSimpleCrud/controllers"
	"golangSimpleCrud/middlewares"

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
		login := new(controllers.LoginController)
		user := new(controllers.UserController)

		v1.POST("/login", login.Login)
		v1.GET("/me", middlewares.AuthRequired(), user.GetMe)

		userGroup := v1.Group("user")
		userGroup.Use(middlewares.AuthRequired())
		{
			userGroup.POST("/", middlewares.AdminRoleRequired(), user.Create)
			userGroup.GET("/", middlewares.AdminRoleRequired(), user.GetAll)
			userGroup.GET("/:id", middlewares.AdminRoleRequired(), user.GetOne)
			userGroup.PATCH("/:id", middlewares.AdminRoleRequired(), user.Update)
			userGroup.DELETE("/:id", middlewares.AdminRoleRequired(), user.Delete)
		}
	}

	return router
}
