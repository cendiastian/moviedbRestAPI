package responses

import (
	"project/ca/business/genres"
)

type GenreResponse struct {
	Id   int    `json:"id"`
	Name string `json:"Name"`
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
