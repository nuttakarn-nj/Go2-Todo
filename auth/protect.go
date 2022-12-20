package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func Protect(tokenString string) error {
	key := []byte("==mySignature==")

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		// check alg in headers
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})

	return err
}
