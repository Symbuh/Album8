package middleware

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Symbuh/foundant-technologies-challenge/server/models"
	"github.com/lib/pq"
)

func createConnection() *sql.DB {

	// Open the connection

	// use godotenv to load the .env file
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Read the POSTGRES_URL from the .env and connect to the db.
	// db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	db, err := sql.Open("postgres", "postgres://lputehky:YEyVfD43lSR0IQGCHDr8pnY3UywZs5mz@kashin.db.elephantsql.com/lputehky")

	if err != nil {
		fmt.Print("Error when opening the database connection")
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	// return the connection
	return db
}

func insertImage(image models.Image) (int64, error) {

	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO images (url, name, description) VALUES ($1, $2, $3) RETURNING image_id;`

	var id int64

	err := db.QueryRow(sqlStatement, image.URL, image.Name, image.Description).Scan(&id)

	if err != nil {
		// log.Fatalf("Unable to execute the query. %v", err)
		return id, err
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id, nil
}

func insertTags(id int64, Tags []string) (int64, error) {

	db := createConnection()

	defer db.Close()

	fmt.Print("Tags inside of the query")
	fmt.Print(Tags)
	sqlStatement := (`INSERT INTO image_tags (image_id, tags) VALUES ($1, $2) returning image_id`)

	var image_id int64

	err := db.QueryRow(sqlStatement, id, pq.Array(Tags)).Scan(&image_id)

	if err != nil {
		fmt.Printf("Unable to execute the query, failed to insert tags!")
		return image_id, err
	}

	fmt.Printf("Inserted a single record %v", image_id)

	return image_id, nil
}

func getImage(id int64) (models.Image, error) {

	db := createConnection()

	defer db.Close()

	var image models.Image

	sqlStatement := `SELECT images.image_id, images.url, images.name, images.description, image_tags.tags FROM images LEFT OUTER JOIN image_tags ON images.image_id=image_tags.image_id WHERE images.image_id=$1;`

	row := db.QueryRow(sqlStatement, id)
	fmt.Print("row:")
	fmt.Println(row)

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
}

func getAllImages() ([]models.Image, error) {

	db := createConnection()

	defer db.Close()

	var images []models.Image

	sqlStatement := `SELECT images.image_id, images.url, images.name, images.description, image_tags.tags FROM images LEFT OUTER JOIN image_tags ON images.image_id=image_tags.image_id;`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return images, err
	}

	defer rows.Close()

	for rows.Next() {
		var image models.Image

		err = rows.Scan(&image.ID, &image.URL, &image.Name, &image.Description, pq.Array(&image.Tags))

		if err != nil {
			return images, err
		}

		// append the user in the users slice
		images = append(images, image)

	}

	// return empty user on error
	return images, err
}

func getImagesByTag(tag string) ([]models.Image, error) {

	db := createConnection()

	defer db.Close()

	var images []models.Image

	sqlStatement := `SELECT images.image_id, images.url, images.name, images.description, image_tags.tags
									FROM images
									LEFT OUTER JOIN image_tags
									ON images.image_id=image_tags.image_id
									WHERE images.image_id
									IN (SELECT image_id from image_tags as ids
									WHERE tags @> ARRAY[$1]
									ORDER BY image_id);`

	rows, err := db.Query(sqlStatement, tag)

	if err != nil {
		log.Fatalf("Unable to execute query! %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var image models.Image

		err = rows.Scan(&image.ID, &image.URL, &image.Name, &image.Description, pq.Array(&image.Tags))

		if err != nil {
			return images, err
			// log.Fatalf("Unable to scan the row. %v", err)
		}
		images = append(images, image)
	}

	return images, err
}

func getTags() ([]string, error) {

	db := createConnection()

	defer db.Close()

	var tags []string

	sqlStatment := `SELECT tag AS tgs
									FROM image_tags t, unnest(t.tags) AS tag
									GROUP BY tag;`

	rows, err := db.Query(sqlStatment)

	if err != nil {
		return tags, err
	}

	defer rows.Close()

	for rows.Next() {
		var tag string

		err = rows.Scan(&tag)

		if err != nil {
			return tags, err
		}

		tags = append(tags, tag)
	}

	return tags, err
}

func deleteImage(id int64) (int64, error) {

	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM images WHERE image_id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return rowsAffected, err
	}

	return rowsAffected, nil
}

func delete_image_tags(id int64) (int64, error) {

	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM image_tags WHERE image_id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		// log.Fatalf("Unable to execute the query. %v", err)
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return -1, err
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected, nil
}
