package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", web.homeHandler)
	log.Println("Server up on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
