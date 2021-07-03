package router

import (
	"github.com/Symbuh/foundant-technologies-challenge/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/image/{id}", middleware.GetImage).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/image", middleware.GetAllImages).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newimage", middleware.CreateImage).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deleteimage/{id}", middleware.DeleteImage).Methods("DELETE", "OPTIONS")
	// router.HandleFunc("/api/getImageBy/{tag}", middleware.GetImageByTag).Methods("GET", "OPTIONS")

	return router
}
