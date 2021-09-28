package genres

import (
	"context"
	"project/ca/app/middlewares"
	"time"
)

type Usecases struct {
	ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewGenreUsecase(repo Repository, timeout time.Duration) Usecase {
	return &Usecases{
		// ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *Usecases) GetAllGenre(c context.Context) ([]Genre, error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	genre, err := uc.Repo.GetAllGenre(ctx)
	if err != nil {
		return []Genre{}, err
	}

	return genre, nil
}
