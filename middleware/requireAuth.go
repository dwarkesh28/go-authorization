package middleware

import (
	"fmt"
	"go-jwt/dto"
	"go-jwt/initializers"
	"go-jwt/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	log.Print("tokenString:", tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		log.Print("exp:", claims["exp"])
		// Check expiration time
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			log.Print("1")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Find the user with token sub
		var user models.User
		// initializers.DB.Table("users").Select("id", "email", "password", ).Where("ID = ?", claims["sub"]).Scan(&user)
		initializers.DB.First(&user, "id = ?", claims["sub"])
		if user.ID == 0 {
			log.Print("2")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Attach to request
		users := dto.ToDto(user)
		c.Set("user", users)

		// continue
		c.Next()

	} else {
		log.Print("3")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
