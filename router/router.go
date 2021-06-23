package router

import (
	"github.com/Symbuh/foundant-technologies-challenge/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/{id}", middleware.GetImage).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user", middleware.GetAllImages).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newuser", middleware.CreateImage).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteuser/{id}", middleware.DeleteImage).Methods("DELETE", "OPTIONS")

	return router
}
