package helper

import (
	"log"
	"mekarTechnicalTest/utils/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//GenerateToken is a function for generate new token
func GenerateToken(username string, duration int64) string {
	mySigningKey := []byte(constants.TOKEN_KEY)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Unix() + duration,
	})
	tokenKey, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println(err)
		LogApp("Generate token key error")
	}
	return tokenKey
}

//ValidateToken is a function for validate generated token
func ValidateToken(tokenString string) bool {
	mySigningKey := []byte(constants.TOKEN_KEY)

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		log.Printf(" | %v", err)
		LogApp("Validate token error")
		return false
	}
	return token.Valid
}
