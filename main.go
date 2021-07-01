package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Symbuh/foundant-technologies-challenge/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
