package omdb

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/ca/business/omdb"

	"gorm.io/gorm"
)

type MysqlAPIRepository struct {
	httpClient http.Client
}

func NewMysqlAPIRepository(connect *gorm.DB) omdb.Repository {
	return &MysqlAPIRepository{
		httpClient: http.Client{},
	}
}

func (rep *MysqlAPIRepository) GetAPI(ctx context.Context, ImdbId string) (omdb.GetAPI, error) {
	var movie GetMovieAPI
	fmt.Println(ImdbId)
	req, err := http.NewRequest("GET", "http://www.omdbapi.com/?apikey=8b8a25e8&", nil)
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
