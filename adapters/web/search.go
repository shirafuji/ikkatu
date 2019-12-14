package web

import (
	"html/template"
	"net/http"

	engine "github.com/shirafuji/ikkatu/engines"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	resp := engine.SearchTabelog(&engine.TabelogRequest{Area: r.URL.Query().Get("area"), Genre: r.URL.Query().Get("genre")})
	if resp.Error != nil {
		http.Error(w, resp.Error.Message, resp.Error.Code)
		return
	}
	tabelogResult := resp.Result
	t := template.Must(template.ParseFiles("templates/home.html"))
	if err := t.Execute(w, tabelogResult); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
