package responses

import (
	"project/ca/business/genres"
	"project/ca/business/movies"
	"project/ca/business/ratings"
	"time"
)

type GenreResponse struct {
	Id   int    `json:"id"`
	Name string `json:"Name"`
}

type MovieResponse struct {
	Id        int               `json:"id"`
	Title     string            `json:"Title"`
	Year      string            `json:"Year"`
	ImdbId    string            `json:"imdbId"`
	Type      string            `json:"Type"`
	Poster    string            `json:"Poster"`
	Genre     []genres.Genre    `json:"Genre"`
	Ratings   []ratings.Ratings `json:"Ratings"`
	Rating    float32           `json:"Rating"`
	Writer    string            `json:"Writer"`
	Actors    string            `json:"Actors"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

func FromDomainGenre(domain genres.Genre) GenreResponse {
	return GenreResponse{
		Id:   domain.Id,
		Name: domain.Name,
	}
}

func FromDomainMovie(domain movies.Movie) MovieResponse {
	return MovieResponse{
		Id:        domain.Id,
		Title:     domain.Title,
		Year:      domain.Year,
		ImdbId:    domain.ImdbId,
		Type:      domain.Type,
		Poster:    domain.Poster,
		Genre:     domain.Genre,
		Ratings:   domain.Ratings,
		Rating:    domain.Rating,
		Writer:    domain.Writer,
		Actors:    domain.Actors,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(domain []movies.Movie) (response []MovieResponse) {
	for _, movie := range domain {
		response = append(response, FromDomainMovie(movie))
	}
	return
}
