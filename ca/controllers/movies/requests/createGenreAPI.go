package requests

import "project/ca/business/movies"

type CreateGenreAPI struct {
	Name string `json:"Name"`
}

func (genre *CreateGenreAPI) ToDomainGenre() movies.Genre {
	return movies.Genre{
		Name: genre.Name,
	}
}
func (genre *CreateGenreAPI) ToDomainMovie() movies.Movie {
	return movies.Movie{
		Genre: []movies.Genre{
			{Name: genre.Name},
		},
	}
}

type GenreString struct {
	Genre string
}
