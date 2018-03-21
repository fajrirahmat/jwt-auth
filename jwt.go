package util

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

func ValidateWithRSAPublicKey(tokenString string, key []byte) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(key)
	})

	return validate(token, err)
}

func ValidateWithRSAPrivateKey(tokenString string, key []byte) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPrivateKeyFromPEM(key)
	})

	return validate(token, err)
}

func validate(token *jwt.Token, err error) bool {
	if err != nil {
		log.Fatal(err)
	}

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			log.Fatal("That's not even token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			log.Fatal("Timing is everything")
		} else {
			log.Fatal("Couldn't handle this token")
		}
	} else {
		log.Fatal("Couldn't handle this token")
	}
	return !token.Valid
}
