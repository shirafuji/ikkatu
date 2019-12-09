package main

import (
	"github.com/shirafuji/ikkatu/adapters/web"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", web.HomeHandler)
	log.Println("Server up on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
