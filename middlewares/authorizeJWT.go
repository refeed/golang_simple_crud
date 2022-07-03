package middlewares

import (
	"golangSimpleCrud/models"
	"golangSimpleCrud/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA_LEN = len("Bearer ")
		authHeader := ctx.GetHeader("Authorization")

		if len(authHeader) < BEARER_SCHEMA_LEN {
			// Might be better to reply with no body
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			ctx.Abort()
			return
		}

		tokenString := authHeader[BEARER_SCHEMA_LEN:]
		username, err := services.GetUsernameFromToken(tokenString)

		if username == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("username", username)
	}
}

// Must be chained after middlewares.AuthRequired()
func AdminRoleRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.MustGet("username").(string)
		user, error := models.GetUserById(username)

		if error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			ctx.Abort()
			return
		}
		if user.Role != "admin" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Admin role required"})
			ctx.Abort()
			return
		}
	}
}
