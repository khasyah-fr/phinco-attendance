package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func JwtAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization header"))
			return
		}

		// Extract the token from the "Bearer" scheme
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the signing method and secret key used to sign the token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid token signing method")
			}

			// Replace "your-secret-key" with your actual secret key used to sign the token
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token: " + err.Error()))
			return
		}

		// Check if the token is valid
		if token.Valid {
			// Token is valid, proceed to the next middleware or handler
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token"))
			return
		}
	})
}
