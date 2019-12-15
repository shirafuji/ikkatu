package engine

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/shirafuji/ikkatu/domain"
)

type IkkyuResponse struct {
	Result *IkkyuResult
	Error  *domain.Error
}

type IkkyuRequest struct {
	Area  string
	Genre string
}

type IkkyuResult struct {
	Result []struct {
		Name   string `json:"name"`
		Url    string `json:"url"`
		Rating string `json:"rating"`
		Budget string `json:"budget"`
		Info   string `json:"info"`
	} `json:"result"`
}

func SearchIkkyu(req *IkkyuRequest) *IkkyuResponse {
	values := url.Values{}
	values.Add("area", req.Area)
	values.Add("genre", req.Genre)
	resp, err := http.Get(domain.APIURL + "/ikkyu" + "?" + values.Encode())
	if err != nil {
		return &IkkyuResponse{
			Error: &domain.Error{
				Code:    500,
				Message: err.Error(),
			},
		}
	}
	data := &IkkyuResult{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return &IkkyuResponse{
			Error: &domain.Error{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	return &IkkyuResponse{
		Result: data,
	}
}
