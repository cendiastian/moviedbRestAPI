package omdb

import "context"

type GetAPI struct {
	Title  string
	Year   string
	ImdbId string
	Type   string
	Poster string
	Genre  string
	Writer string
	Actors string
}

type Repository interface {
	GetAPI(ctx context.Context, ImdbId string) (GetAPI, error)
}
