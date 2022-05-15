package usecase

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/blacknvcone/efishery-testcase/svc-fetch/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type iamUseCase struct {
	contextTimeout time.Duration
	JWTSecret      string
}

func NewIAMUseCase(jwtSecretKey string, timeout time.Duration) domain.IIamUseCase {
	return &iamUseCase{
		contextTimeout: timeout,
		JWTSecret:      jwtSecretKey,
	}
}

func verifyToken(ts, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("mySecret"), nil
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
				"success": false,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		tokenString := strings.Replace(ts, "Bearer ", "", -1)

		//Verify
		token, err := verifyToken(tokenString, uc.JWTSecret)

		if token != nil && err == nil {
			claim, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				c.Set("claim", claim)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid Token",
			})
			c.Abort()
			return
		}

	}
}

func (uc *iamUseCase) IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {

		res, exist := c.Get("claim")
		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid Session",
				"data":    nil,
			})
			c.Abort()
			return
		}

		// Parsing claim interface and Validate Role
		v := reflect.ValueOf(res)
		if v.Kind() == reflect.Map {
			for _, key := range v.MapKeys() {
				value := v.MapIndex(key)
				//fmt.Println(key.Interface(), value.Interface())
				if key.Interface() == "role" {
					if value.Interface() == "super-admin" || value.Interface() == "admin" {
						c.Next()
						return
					}
					c.JSON(http.StatusForbidden, gin.H{
						"success": false,
						"message": "Insufficient access level !",
					})
					c.Abort()
					return
				}
			}
		}

	}
}
