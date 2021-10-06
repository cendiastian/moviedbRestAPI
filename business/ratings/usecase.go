package ratings

import (
	"context"
	"fmt"
	"project/app/middlewares"
	resp "project/business"
	"time"
)

type RateUseCase struct {
	ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewRateUsecase(repo Repository, timeout time.Duration) Usecase {
	return &RateUseCase{
		// ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *RateUseCase) Detail(c context.Context, domain Ratings) (res Ratings, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	rate, err := uc.Repo.Detail(ctx, domain.MovieId, domain.UserId)
	if err != nil {
		return Ratings{}, resp.ErrNotFound
	}
	return rate, nil
}

func (uc *RateUseCase) Delete(c context.Context, domain Ratings) error {

	if domain.MovieId == 0 || domain.UserId == 0 {
		fmt.Println("user err")
		return resp.ErrFillData
	}

	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	_, err := uc.Repo.Detail(ctx, domain.MovieId, domain.UserId)
	if err != nil {
		return resp.ErrNotFound
	}
	domain.UpdatedAt = time.Now()

	err = uc.Repo.Delete(ctx, domain.MovieId, domain.UserId)
	if err != nil {
		return resp.ErrInternalServer
	}

	return nil
}

func (uc *RateUseCase) Update(c context.Context, domain Ratings) error {

	if domain.MovieId == 0 || domain.UserId == 0 {
		return resp.ErrNotFound
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	_, err := uc.Repo.Detail(ctx, domain.MovieId, domain.UserId)
	if err != nil {
		return err
	}
	domain.UpdatedAt = time.Now()

	err = uc.Repo.Update(ctx, domain)
	if err != nil {
		return err
	}

	return nil

}

func (uc *RateUseCase) Create(c context.Context, domain Ratings) (Ratings, error) {

	if domain.MovieId == 0 || domain.UserId == 0 || domain.Rate == 0 {
		return Ratings{}, resp.ErrNotFound
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	Rate, err := uc.Repo.Create(ctx, domain)
	if err != nil {
		return Ratings{}, err
	}

	return Rate, nil

}
