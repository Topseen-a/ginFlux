package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "goGin/db"
    "goGin/models"
)


func RequiredAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	userID := claims["sub"].(float64)

	var user models.User
	if err := db.DB.First(&user, uint(userID)).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	
	c.Set("currentUser", user)
	c.Next()
}