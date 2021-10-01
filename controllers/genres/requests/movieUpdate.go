package requests

import (
	"project/business/movies"
)

type MovieUpdate struct {
	Id    int    `json:"id"`
	Title string `json:"Title"`
	Type  string `json:"Type"`
}

func (movie *MovieUpdate) ToDomain() movies.Movie {
	return movies.Movie{
		Id:    movie.Id,
		Title: movie.Title,
		Type:  movie.Type,
	}
}
