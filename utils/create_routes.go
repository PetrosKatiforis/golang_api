package utils

import (
	"net/http"
	"pioyi/golang_api/middleware"
	"pioyi/golang_api/routes"

	"github.com/gorilla/mux"
)

func CreateRoutes() http.Handler {
	router := mux.NewRouter()

	// User Routes
	router.HandleFunc("/users/new", routes.CreateUser).Methods("POST")
	router.HandleFunc("/users/login", routes.Login).Methods("POST")
	router.HandleFunc("/users/logout", routes.Logout).Methods("POST")

	// Posts Routes
	router.HandleFunc("/posts/all", routes.GetAllPosts).Methods("GET")
	router.HandleFunc("/users/{id}/posts", routes.GetPostsByUser).Methods("GET")

	// Authentication Required
	router.Handle("/posts/new", middleware.WithAuthentication(http.HandlerFunc(routes.CreatePost))).Methods("POST")
	router.Handle("/posts/{id}/delete", middleware.WithAuthentication(http.HandlerFunc(routes.DeletePost))).Methods("DELETE")

	return router
}
