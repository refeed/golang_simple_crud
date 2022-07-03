package controllers

import (
	"golangSimpleCrud/contracts"
	"golangSimpleCrud/models"
	"golangSimpleCrud/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (c *LoginController) Login(ctx *gin.Context) {
	var loginReq contracts.LoginReq
	err := ctx.BindJSON(&loginReq)
	if err != nil {
		return
	}

	isUserAuthenticated := models.IsUserPasswordMatch(loginReq.Username, loginReq.Password)
	if !isUserAuthenticated {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token := services.GenerateToken(loginReq.Username)

	// TODO: Add also a header that makes the browser will save the token in a cookie
	ctx.JSON(200, contracts.LoginRes{Token: token})
}
