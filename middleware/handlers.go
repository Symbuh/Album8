package middleware

import (
	//"database/sql"
	// "encoding/json"
	// "fmt"
	// "go-postgres/models"
	// "log"
	// "net/http"
	// "os"
	// "strconv"

	// "github.com/gorilla/mux"
	// "github.com/joho/godotenv" // package used to read the .env file

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Symbuh/foundant-technologies-challenge/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// Check for preflight options request
	if r.Method == "OPTIONS" {
		return
	}

	var image models.Image

	// decode the json request
	err := json.NewDecoder(r.Body).Decode(&image)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Print(image.Name)
	fmt.Print(image.URL)
	fmt.Print(image.Tags)

	// call insert function to insert an ID
	insertedImageID, err := insertImage(image)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// attempt to insert tags
	insertedTagID, err := insertTags(insertedImageID, image.Tags)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print(insertedTagID)

	res := response{
		ID:      insertedImageID,
		Message: "Image saved successfully",
	}
	json.NewEncoder(w).Encode(res)
}

// Return a single image by ID
func GetImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	image, err := getImage(int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send the response
	json.NewEncoder(w).Encode(image)
}

// Returns all images
func GetAllImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	images, err := getAllImages()

	if err != nil {
		log.Fatalf("Unable to get all images! %v", err)
	}

	json.NewEncoder(w).Encode(images)
}

func GetImageByTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)

	images, err := getImagesByTag(params["tag"])

	if err != nil {
		log.Fatalf("Unable to get all images! %v", err)
	}

	json.NewEncoder(w).Encode(images)
}

func GetTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	images, err := getTags()

	if err != nil {
		log.Fatalf("Unable to get all images! %v", err)
	}
	fmt.Print(images)

	json.NewEncoder(w).Encode(images)
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deletedRows, err := deleteImage(int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	deletedTagRows, err := delete_image_tags(int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows+deletedTagRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
