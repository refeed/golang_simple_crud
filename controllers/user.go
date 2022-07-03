package controllers

import (
	"errors"
	"golangSimpleCrud/contracts"
	"golangSimpleCrud/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct{}

func (u *UserController) GetAll(ctx *gin.Context) {
	users := models.GetAllUsers()
	ctx.JSON(200, users)
}

func (u *UserController) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := models.GetUserById(id)

	if errors.Is(err, mongo.ErrNoDocuments) {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	getUserRes := contracts.GetUserRes{
		Name:     user.Name,
		Role:     user.Role,
		Username: user.Username,
	}
	ctx.JSON(200, getUserRes)
}

func (u *UserController) Create(ctx *gin.Context) {
	var createUserForm contracts.CreateUserReq
	if err := ctx.MustBindWith(&createUserForm, binding.JSON); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := models.CreateUser(createUserForm)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			ctx.JSON(400, gin.H{"error": "Username already exists"})
			return
		}

		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// TODO: Consider returning the user data instead
	ctx.JSON(200, gin.H{
		"message": "Created user",
	})
}

func (u *UserController) Update(ctx *gin.Context) {
	var updateUserForm contracts.UpdateUserReq
	if err := ctx.ShouldBindJSON(&updateUserForm); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := models.UpdateUser(ctx.Param("id"), updateUserForm)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Successfully updated user",
	})
}

func (u *UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	isDeleted, err := models.DeleteUserById(id)

	if !isDeleted {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User deleted",
	})
}

func (u *UserController) GetMe(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}
