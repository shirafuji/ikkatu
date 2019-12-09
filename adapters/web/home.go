package web

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/home.html"))
	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}
