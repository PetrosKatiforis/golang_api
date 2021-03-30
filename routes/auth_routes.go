package routes

import (
	"encoding/json"
	"net/http"
	"pioyi/golang_api/database"
	"pioyi/golang_api/entity"
	"pioyi/golang_api/helpers"
	"pioyi/golang_api/interfaces"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {
	// Reading the body, creating the new user instance
	var newUser entity.User
	_ = json.NewDecoder(req.Body).Decode(&newUser)
	errors := helpers.Validate(newUser)
	w.Header().Set("Content-Type", "application/json")

	if errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}

	doesExist := !database.SQLDatabase.Find(&entity.User{}, "username = ?", newUser.Username).RecordNotFound()
	if doesExist {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"username": "Username already exists!",
		})
		return
	}

	// Hashing user's password
	newUser.Password = helpers.HashAndSalt(newUser.Password)

	// Generating token and response
	database.SQLDatabase.Create(&newUser)
	token, _ := helpers.GenerateJwtCookie(&w, &newUser)
	response := interfaces.UserResponse{User: newUser, Token: token}

	// Returning the json data
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, req *http.Request) {
	// Parsign the request's body
	var userData entity.User
	_ = json.NewDecoder(req.Body).Decode(&userData)
	w.Header().Set("Content-Type", "application/json")

	// Searching for the user inside the database
	var targetUser entity.User
	doesExist := database.SQLDatabase.Where("username = ?", userData.Username).Find(&targetUser).Error

	// Performing some validation
	errors := make(map[string]string)
	if doesExist != nil {
		errors["username"] = "Username does not exist!"
	} else {
		// Checking whether the two passwords match
		matchError := bcrypt.CompareHashAndPassword([]byte(targetUser.Password), []byte(userData.Password))
		if matchError != nil {
			errors["password"] = "Passwords do not match!"
		}
	}

	// Returning the errors if there are any
	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}

	// Sending token cookie
	token, _ := helpers.GenerateJwtCookie(&w, &targetUser)

	// Returning useful data
	json.NewEncoder(w).Encode(map[string]string{
		"token":    token,
		"username": userData.Username,
	})
}

func Logout(w http.ResponseWriter, req *http.Request) {
	helpers.ClearJwtCookie(&w)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "The cookie was deleted successfully",
	})
}
