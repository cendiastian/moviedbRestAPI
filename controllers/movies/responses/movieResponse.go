package responses

import (
	"project/business/movies"
	_genreRespone "project/controllers/genres/responses"
	_ratingRespone "project/controllers/ratings/responses"
	"time"
)

type MovieResponse struct {
	Id        int                             `json:"id"`
	Title     string                          `json:"title"`
	Year      string                          `json:"year"`
	ImdbId    string                          `json:"imdbId"`
	Type      string                          `json:"type"`
	Poster    string                          `json:"poster"`
	Genre     []_genreRespone.GenreResponse   `json:"genre"`
	Ratings   []_ratingRespone.RatingResponse `json:"ratings"`
	Rating    float32                         `json:"rating"`
	Writer    string                          `json:"writer"`
	Actors    string                          `json:"actors"`
	CreatedAt time.Time                       `json:"createdAt"`
	UpdatedAt time.Time                       `json:"updatedAt"`
}

func FromDomainMovie(domain movies.Movie) MovieResponse {
	return MovieResponse{
		Id:        domain.Id,
		Title:     domain.Title,
		Year:      domain.Year,
		ImdbId:    domain.ImdbId,
		Type:      domain.Type,
		Poster:    domain.Poster,
		Genre:     _genreRespone.ToListDomain(domain.Genre),
		Ratings:   _ratingRespone.ToListDomain(domain.Ratings),
		Rating:    domain.Rating,
		Writer:    domain.Writer,
		Actors:    domain.Actors,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
