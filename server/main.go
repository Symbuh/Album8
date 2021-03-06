package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Symbuh/foundant-technologies-challenge/server/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	http.Handle("/", r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))
	log.Fatal(http.ListenAndServe(":8080", r))
}
