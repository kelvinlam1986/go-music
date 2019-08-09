package router

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-music/controllers"
	"net/http"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		logrus.Info(fmt.Sprintf("Request path %s", path))
		if path == "/authenticate" && c.Request.Method == http.MethodPost {
			c.Next()
			return
		}

		tokenString, err := request.HeaderExtractor{"access-token"}.ExtractToken(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "error" : "Please pass the access-token in the Header request"} )
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(controllers.SECRET_KEY), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access Denied; Please check the access token"})
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Next()
			return
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access Denied; Please check the access token"})
			return
		}
	}
}
