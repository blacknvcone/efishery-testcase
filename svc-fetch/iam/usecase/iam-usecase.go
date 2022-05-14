package usecase

import (
	"fmt"
	"net/http"
	"time"

	"github.com/blacknvcone/efishery-testcase/svc-fetch/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type iamUseCase struct {
	contextTimeout time.Duration
}

func NewIAMUseCase(timeout time.Duration) domain.IAMUseCase {
	return &iamUseCase{
		contextTimeout: timeout,
	}
}

func verifyToken(ts string) (*jwt.Token, error) {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("lontongbalap"), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func (uc *iamUseCase) AuthorizationHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Retrieving Token
		ts := c.Request.Header.Get("Authorization")
		if ts == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		//Verify
		token, err := verifyToken(ts)
		if token != nil && err == nil {
			_, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				c.Abort()
				return

			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Invalid Token",
			})
			c.Abort()
			return

		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Invalid Token",
		})
		c.Abort()
		return
	}
}
