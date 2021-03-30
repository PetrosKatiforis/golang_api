package middleware

import (
	"net/http"
	"os"

	"context"

	"pioyi/golang_api/interfaces"

	"github.com/dgrijalva/jwt-go"
)

func WithAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Getting the cookie
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Couldn't find cookie!", http.StatusUnauthorized)
			return
		}

		cookieValue := cookie.Value
		claims := &interfaces.Claims{}

		// Parsing the Json Web Token
		jwtKey := []byte(os.Getenv("JWT_SECRET"))
		token, err := jwt.ParseWithClaims(cookieValue, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		// Invalid token error
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token!", http.StatusUnauthorized)
			return
		}

		// Passing the token data to the request's context
		ctx := context.WithValue(r.Context(), "token", &interfaces.ContextData{Token: token.Raw, Data: claims})

		// Calling the route handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
