package helpers

import (
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/khasyah-fr/phinco-attendance/entities"
)

func DataToResponse(statusCode int, data interface{}) []byte {
	var response entities.ResponseEntity

	response.Status = statusCode
	response.Data = data

	json, _ := json.Marshal(response)

	return json
}

func GenerateJWTToken(username string) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) of the token
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

	// Sign the token with a secret key
	secretKey := []byte("secret")
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
