package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Protect(signature []byte) gin.HandlerFunc {


	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(authorization, "Bearer ")

		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)

			// check alg in headers's token
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return signature, nil
		})

		if err != nil {
			// stop to go next middleware
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// next step
		ctx.Next()
	}
}
