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
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Symbuh/foundant-technologies-challenge/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgres db
func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

// CreateUser create a user in the postgres db
func CreateImage(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// Check for preflight options request
	if r.Method == "OPTIONS" {
		return
	}
	// w.Header().Set("Access-Control-Allow-Methods", "POST")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Acess-Control-Allow-Headers, Authorization, X-Requested-With")

	// create an empty user of type models.User
	var image models.Image

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&image)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Print(image.Name)
	fmt.Print(image.URL)
	fmt.Print(image.Tags)
	// call insert user function and pass the user
	insertedImageID, err := insertImage(image)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	insertedTagID, err := insertTags(insertedImageID, image.Tags)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print(insertedTagID)
	// format a response object
	res := response{
		ID:      insertedImageID,
		Message: "Image saved successfully",
	}
	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(http.StatusOK)
	// }
	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetUser will return a single user by its id
func GetImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call the getUser function with user id to retrieve a single user
	image, err := getImage(int64(id))
	fmt.Println("Here is our error: ")
	fmt.Println(err)
	if err != nil {
		fmt.Println("We have entered the magical error block in question")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Println("We are about to send the naughty response back")
	// send the response
	json.NewEncoder(w).Encode(image)
}

// GetAllUser will return all the users
func GetAllImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the users in the db
	images, err := getAllImages()

	if err != nil {
		log.Fatalf("Unable to get all images! %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(images)
}

// DeleteUser delete user's detail in the postgres db
func DeleteImage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call the deleteUser, convert the int to int64
	deletedRows, err := deleteImage(int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

//------------------------- handler functions ----------------
// insert one user in the DB
func insertImage(image models.Image) (int64, error) {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO images (url, name, description) VALUES ($1, $2, $3) RETURNING image_id;`

	/*
		Here I believe we'll have to accept an array of tags and insert them into
		the database seperately

		I think it may be necessary to create an images with tags model
	*/

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, image.URL, image.Name, image.Description).Scan(&id)

	if err != nil {
		// log.Fatalf("Unable to execute the query. %v", err)
		return id, err
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id, nil
}

func insertTags(id int64, tags []string) (int64, error) {
	db := createConnection()

	defer db.Close()
	fmt.Print(tags)
	sqlStatement := (`INSERT INTO image_tags (image_id, tags) VALUES ($1, $2) returning image_id`)

	var image_id int64

	err := db.QueryRow(sqlStatement, id, pq.Array(tags)).Scan(&image_id)

	if err != nil {
		fmt.Printf("Unable to execute the query, failed to insert tags!")
		return image_id, err
	}

	fmt.Printf("Inserted a single record %v", image_id)

	return image_id, nil
}

// get one user from the DB by its userid
func getImage(id int64) (models.Image, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var image models.Image
	//var tag models.Tag
	// create the select sql query
	sqlStatement := `SELECT images.image_id, images.url, images.name, images.description, image_tags.tags FROM images LEFT OUTER JOIN image_tags ON images.image_id=image_tags.image_id WHERE images.image_id=$1;`
	// We may have to search by image name here insetead but for now this will work.
	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)
	fmt.Print("row:")
	fmt.Println(row)
	// unmarshal the row object to user
	err := row.Scan(&image.ID, &image.Name, &image.Description, &image.URL, pq.Array(&image.Tags))
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return image, err
	case nil:
		return image, nil
	default:
		return image, err
	}

	// return empty user on error
	// return image, err
}

// get one user from the DB by its userid
func getAllImages() ([]models.Image, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var images []models.Image

	// create the select sql query
	sqlStatement := `SELECT images.image_id, images.url, images.name, images.description, image_tags.tags FROM images LEFT OUTER JOIN image_tags ON images.image_id=image_tags.image_id;`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var image models.Image

		// unmarshal the row object to user
		err = rows.Scan(&image.ID, &image.URL, &image.Name, &image.Description, pq.Array(&image.Tags))

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		images = append(images, image)

	}

	// return empty user on error
	return images, err
}

// delete user in the DB
func deleteImage(id int64) (int64, error) {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM images WHERE image_id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		return -1, err
		// return res, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return rowsAffected, err
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected, nil
}

func delete_image_tags(id int64) (int64, error) {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM image_tags WHERE image_id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		// log.Fatalf("Unable to execute the query. %v", err)
		return -1, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return -1, err
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected, nil
}
