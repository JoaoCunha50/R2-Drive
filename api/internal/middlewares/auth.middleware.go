package middlewares

import (
	"api/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func HasAuthentication(c *gin.Context) {
	tokenStr := utils.ExtractToken(c)
	if tokenStr == "" {
		c.Next()
	}

	token, err := utils.TokenDecode(tokenStr)
	if err != nil || !token.Valid {
		c.Next()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.Next()
	}

	currentUser := map[string]any {
		"id": uint(claims["sub"].(float64)),
		"email": claims["email"].(string),
		"username": claims["username"].(string),
		"name": claims["name"].(string),
		"role": claims["role"].(string),
	}

	c.Set("currentUser", currentUser)
	c.Next()
}
