package responses

import (
	"project/ca/business/genres"
	"time"
)

type GenreResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"Name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain genres.Genre) GenreResponse {
	return GenreResponse{
		Id:   domain.Id,
		Name: domain.Name,
	}
}

func ToListDomain(domain []genres.Genre) (response []GenreResponse) {
	for _, genre := range domain {
		response = append(response, FromDomain(genre))
	}
	return
}

// func ToListDomain(domain []genres.Genre) (response []GenreResponse) {
// 	for _, genre := range domain {
// 		response = append(response, FromDomain(genre))
// 	}
// 	return
// }
