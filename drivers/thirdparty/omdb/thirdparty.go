package omdb

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/business/omdb"
)

type Omdb struct {
	Client http.Client
	Url    string
	Key    string
	Symbol string
}

func NewOmdbAPI(api Omdb) *Omdb {
	return &Omdb{
		Client: http.Client{},
		Url:    api.Url,
		Key:    api.Key,
		Symbol: api.Symbol,
	}
}

func (api *Omdb) GetAPI(ctx context.Context, ImdbId string) (omdb.GetAPI, error) {
	var movie GetMovieAPI
	fmt.Println(ImdbId)
	req, err := http.NewRequest("GET", api.Url+api.Key+api.Symbol, nil)
	if err != nil {
		return omdb.GetAPI{}, err
	}

	query := req.URL.Query()
	query.Add("i", ImdbId)
	req.URL.RawQuery = query.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return omdb.GetAPI{}, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &movie)

	return movie.ToDomainAPI(), nil
}
