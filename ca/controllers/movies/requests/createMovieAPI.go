package requests

import (
	"project/ca/business/movies"
)

type CreateMovieAPI struct {
	Title  string         `json:"Title"`
	Year   string         `json:"Year"`
	ImdbId string         `json:"imdbId"`
	Type   string         `json:"Type"`
	Poster string         `json:"Poster"`
	Genre  []movies.Genre `json:"Genre"`
	Writer string         `json:"Writer"`
	Actors string         `json:"Actors"`
}

// type Genre struct {
// 	Genre  []movies.Genre `json:"Genre"`
// }

func (movie *CreateMovieAPI) ToDomain() movies.Movie {
	return movies.Movie{
		Title:  movie.Title,
		Year:   movie.Year,
		ImdbId: movie.ImdbId,
		Type:   movie.Type,
		Poster: movie.Poster,
		Genre:  movie.Genre,
		Writer: movie.Writer,
		Actors: movie.Actors,
	}
}
