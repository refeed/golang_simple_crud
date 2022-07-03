package services

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	SECRET_KEY = "SECRETKEY" // TODO: Read from env
	ISSUER     = "golang-jwt"
)

type authCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) string {
	claims := &authCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    ISSUER,
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		panic(err)
	}
	return t
}

// Returns the username from the token, if it is empty string then the token
// is invalid
func GetUsernameFromToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*authCustomClaims); ok && token.Valid {
		return claims.Username, nil
	}
	return "", nil
}
