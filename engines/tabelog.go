package engine

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/shirafuji/ikkatu/domain"
)

type TabelogResponse struct {
	Result interface{}
	Error  *domain.Error
}

type TabelogRequest struct {
	Area  string
	Genre string
}

type TabelogResult struct {
	Result []struct {
		Name   string `json:"name"`
		Url    string `json:"url"`
		Rating string `json:"rating"`
		Budget string `json:"budget"`
	} `json:"result"`
}

func SearchTabelog(req *TabelogRequest) *TabelogResponse {
	values := url.Values{}
	values.Add("area", req.Area)
	values.Add("genre", req.Genre)
	resp, err := http.Get(domain.APIURL + "/tabelog" + "?" + values.Encode())
	if err != nil {
		return &TabelogResponse{
			Error: &domain.Error{
				Code:    500,
				Message: err.Error(),
			},
		}
	}
	data := &TabelogResult{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return &TabelogResponse{
			Error: &domain.Error{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	return &TabelogResponse{
		Result: data,
	}
}
