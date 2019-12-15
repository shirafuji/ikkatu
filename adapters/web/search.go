package web

import (
	"html/template"
	"net/http"

	engine "github.com/shirafuji/ikkatu/engines"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	tabelogResp := engine.SearchTabelog(&engine.TabelogRequest{Area: r.URL.Query().Get("area"), Genre: r.URL.Query().Get("genre")})
	if tabelogResp.Error != nil {
		http.Error(w, tabelogResp.Error.Message, tabelogResp.Error.Code)
		return
	}
	tabelogResult := tabelogResp.Result
	ikkyuResp := engine.SearchIkkyu(&engine.IkkyuRequest{Area: r.URL.Query().Get("area"), Genre: r.URL.Query().Get("genre")})
	if ikkyuResp.Error != nil {
		http.Error(w, ikkyuResp.Error.Message, ikkyuResp.Error.Code)
		return
	}
	ikkyuResult := ikkyuResp.Result
	Results := struct {
		Tabelog *engine.TabelogResult
		Ikkyu   *engine.IkkyuResult
	}{
		Tabelog: tabelogResult,
		Ikkyu:   ikkyuResult,
	}
	t := template.Must(template.ParseFiles("templates/home.html"))
	if err := t.Execute(w, Results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
