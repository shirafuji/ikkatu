package engine

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/shirafuji/ikkatu/domain"
)

type YahooResponse struct {
	Result *YahooResult
	Error  *domain.Error
}

type YahooRequest struct {
	Area  string
	Genre string
}

type YahooResult struct {
	Result []struct {
		Name   string `json:"name"`
		Url    string `json:"url"`
		Rating string `json:"rating"`
		Budget string `json:"budget"`
		Info   string `json:"info"`
	} `json:"result"`
}

func SearchYahoo(req *YahooRequest) *YahooResponse {
	values := url.Values{}
	values.Add("area", req.Area)
	values.Add("genre", req.Genre)
	resp, err := http.Get(domain.APIURL + "/yahoo" + "?" + values.Encode())
	if err != nil {
		return &YahooResponse{
			Error: &domain.Error{
				Code:    500,
				Message: err.Error(),
			},
		}
	}
	data := &YahooResult{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return &YahooResponse{
			Error: &domain.Error{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	return &YahooResponse{
		Result: data,
	}
}
