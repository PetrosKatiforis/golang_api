package helpers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pioyi/golang_api/entity"
	"pioyi/golang_api/interfaces"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) string {
	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.MinCost,
	)

	// Checking for errors
	if err != nil {
		log.Fatal("Failed to hash password: ", err.Error())
	}

	// Returning the hashed result as a string
	return string(hashedPassword)
}

func GenerateJwtCookie(w *http.ResponseWriter, user *entity.User) (string, error) {
	// Declaring the expiration time (60 Minutes)
	expirationTime := time.Now().Add(time.Minute * 60)

	// Creating the token's claims
	claims := interfaces.Claims{}
	claims.UserId = user.ID
	claims.Username = user.Username
	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signing the token with the secret key
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		fmt.Printf("Couldn't format token! Error: %s", err.Error())
		return "", err
	}

	// Creating the cookie
	http.SetCookie(*w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Expires:  expirationTime,
	})

	// Finally, returning the token
	return tokenString, nil
}

func ClearJwtCookie(w *http.ResponseWriter) {
	http.SetCookie(*w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
