package routes

import (
	"encoding/json"
	"net/http"
	"pioyi/golang_api/database"
	"pioyi/golang_api/entity"
	"pioyi/golang_api/helpers"
	"pioyi/golang_api/interfaces"

	"github.com/gorilla/mux"
)

func GetAllPosts(w http.ResponseWriter, req *http.Request) {
	var posts []entity.Post
	// Fetching the API's 15 latests posts
	database.SQLDatabase.Order("id desc").Limit(15).Find(&posts)

	json.NewEncoder(w).Encode(posts)
}

func GetPostsByUser(w http.ResponseWriter, req *http.Request) {
	var posts []entity.Post
	userId := mux.Vars(req)["id"]

	// Fetching the user's 15 latest posts and returing them
	database.SQLDatabase.Where("user_id = ?", userId).Find(&posts).Order("id desc").Limit(15)
	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, req *http.Request) {
	// Using the user's context
	username := req.Context().Value("token").(*interfaces.ContextData).Data.Username

	// Reading the request's json body
	var newPost entity.Post
	_ = json.NewDecoder(req.Body).Decode(&newPost)

	// Validation
	errors := helpers.Validate(newPost)

	if errors != nil {
		http.Error(w, "Invalid user input", http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}

	// Fetching the user, adding the post to his account
	var user entity.User
	database.SQLDatabase.Where("username = ?", username).Find(&user)
	database.SQLDatabase.Model(&user).Association("Posts").Append(&newPost)

	// Returning the new post
	json.NewEncoder(w).Encode(newPost)
}

func DeletePost(w http.ResponseWriter, req *http.Request) {
	// Getting the post's id and the user data from the request's context
	// (Originating from the authentication middleware)
	postId := mux.Vars(req)["id"]
	userId := req.Context().Value("token").(*interfaces.ContextData).Data.UserId

	// Fetching target psot
	var post entity.Post
	err := database.SQLDatabase.Where("id = ? AND user_id = ?", postId, userId).Find(&post).Error

	// Checking if the post was found
	if err != nil {
		http.Error(w, "Something went wrong!", http.StatusBadRequest)
		return
	}

	// Deleting the record from the database
	database.SQLDatabase.Unscoped().Delete(&post)

	// Returning message
	json.NewEncoder(w).Encode(map[string]string{
		"message": "The post was deleted successfully!",
	})
}
