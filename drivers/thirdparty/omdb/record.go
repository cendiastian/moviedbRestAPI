package omdb

import (
	"project/business/omdb"
)

type GetMovieAPI struct {
	Title  string
	Year   string
	ImdbId string
	Type   string
	Poster string
	Genre  string
	Writer string
	Actors string
}

func (movie *GetMovieAPI) ToDomainAPI() omdb.GetAPI {
	return omdb.GetAPI{
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
