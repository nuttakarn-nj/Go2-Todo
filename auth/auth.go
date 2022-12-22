package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AccessToken(signature []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// create claims (Payloads)
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
			Audience: "Nuttakarn",
		}

		// sign with headers
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// get token
		signString, err := token.SignedString(signature)

		// failed
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		// success
		ctx.JSON(http.StatusOK, gin.H{
			"token": signString,
		})
	}
}
