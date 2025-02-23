package web

import (
	"html/template"
	"net/http"

	engine "github.com/shirafuji/ikkatu/engines"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	tabelogChan := make(chan *engine.TabelogResponse)
	ikkyuChan := make(chan *engine.IkkyuResponse)
	yahooChan := make(chan *engine.YahooResponse)
	go func() {
		tabelogResp := engine.SearchTabelog(&engine.TabelogRequest{Area: r.URL.Query().Get("area"), Genre: r.URL.Query().Get("genre")})
		tabelogChan <- tabelogResp
	}()
	go func() {
		ikkyuResp := engine.SearchIkkyu(&engine.IkkyuRequest{Area: r.URL.Query().Get("area"), Genre: r.URL.Query().Get("genre")})
		ikkyuChan <- ikkyuResp
	}()
	go func() {
		yahooResp := engine.SearchYahoo(&engine.YahooRequest{Area: r.URL.Query().Get("area"), Genre: r.URL.Query().Get("genre")})
		yahooChan <- yahooResp
	}()
	tabelogResp := <-tabelogChan
	ikkyuResp := <-ikkyuChan
	yahooResp := <-yahooChan

	if tabelogResp.Error != nil {
		http.Error(w, tabelogResp.Error.Message, tabelogResp.Error.Code)
		return
	}
	tabelogResult := tabelogResp.Result
	if ikkyuResp.Error != nil {
		http.Error(w, ikkyuResp.Error.Message, ikkyuResp.Error.Code)
		return
	}
	ikkyuResult := ikkyuResp.Result
	if yahooResp.Error != nil {
		http.Error(w, yahooResp.Error.Message, yahooResp.Error.Code)
		return
	}
	yahooResult := yahooResp.Result
	Results := struct {
		Tabelog *engine.TabelogResult
		Ikkyu   *engine.IkkyuResult
		Yahoo   *engine.YahooResult
		Area    string
		Genre   string
	}{
		Tabelog: tabelogResult,
		Ikkyu:   ikkyuResult,
		Yahoo:   yahooResult,
		Area:    r.URL.Query().Get("area"),
		Genre:   r.URL.Query().Get("genre"),
	}
	t := template.Must(template.ParseFiles("templates/home.html"))
	if err := t.Execute(w, Results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
