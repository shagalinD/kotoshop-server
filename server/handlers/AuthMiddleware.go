package handlers

import (
	"fmt"
	"kotoshop/models"
	"kotoshop/postgres"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func createAccessToken(userId uint, role string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,                    // Subject (user identifier)
		"role": role,
		"iss": "Kotoshop",                  // Issuer
		"exp": time.Now().Add(time.Hour*24*30).Unix(), // Expiration time
		"iat": time.Now().Unix(),                 // Issued at
})

	tokenString, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

// Print information about the created token
	return tokenString, nil
}

func checkToken(tokenString string) (uint, string, error) {
	if tokenString == "" {
			return 0, "", fmt.Errorf("токен отсутствует")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("неверный метод подписи: %v", t.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
			return 0, "", fmt.Errorf("ошибка валидации токена: %v", err)
	}

	if !token.Valid {
			return 0, "", fmt.Errorf("невалидный токен")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
			return 0, "", fmt.Errorf("ошибка разбора claims")
	}

	// Проверка exp
	if exp, ok := claims["exp"].(float64); !ok || float64(time.Now().Unix()) > exp {
			return 0, "", fmt.Errorf("токен истёк")
	}

	// Проверка sub
	sub, ok := claims["sub"].(float64)
	if !ok || sub == 0 {
			return 0, "", fmt.Errorf("токен не содержит id пользователя")
	}
	role, ok := claims["role"].(string)
	if !ok || sub == 0 {
			return 0, "", fmt.Errorf("токен не содержит id пользователя")
	}

	return uint(sub), role, nil
}

func AuthMiddleware(c *gin.Context) {
	tokenString := strings.TrimSpace(c.GetHeader("Authorization"))

	if !strings.HasPrefix(tokenString, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "неверный формат токена"})
			return
	}

	token := strings.TrimSpace(strings.TrimPrefix(tokenString, "Bearer "))
	userID, role, err := checkToken(token)

	if err != nil {
		log.Printf("ошибка валидации токена: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
	}


	if err := postgres.DB.First(&models.User{}, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "не существующий id пользователя"})
			return
	}

	c.Set("userID", userID)
	c.Set("role", role)
	c.Next()
}

