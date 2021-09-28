package search

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://www.omdbapi.com/?apikey=8b8a25e8&s=", nil)
	resp, _ := client.Do(req)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	e := echo.New()
	e.POST("/search", searchMovie)
	e.Start(":8000")
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Search struct {
	Result      []Movie `json:"search"`
	TotalResult string  `json:"totalResults"`
	Response    string  `json:"response"`
	MovieTitle  string  `json:"title"`
}

type Movie struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

type SearchRequest struct {
	MovieTitle string
	Page       int
}

type SearchResponse struct {
	Response     []Movie `json:"Search"`
	ErrorMessage string  `json:"Error,omitempty"`
}

func DecodeGRPCSearchRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SearchRequest)
	return SearchRequest{
		MovieTitle: req.MovieTitle,
		Page:       int(req.Page),
	}, nil
}

func EncodeGRPCSearchResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(SearchResponse)
	result := make([]*pb.Movie, 0)

	for _, movie := range resp.Response {
		result = append(result, &pb.Movie{
			Title:  movie.Title,
			Year:   movie.Year,
			Type:   movie.Type,
			ImdbID: movie.IMDBID,
			Poster: movie.Poster,
		})
	}

	return &pb.SearchResponse{
		MovieList: result,
		Err:       resp.ErrorMessage,
	}, nil
}
