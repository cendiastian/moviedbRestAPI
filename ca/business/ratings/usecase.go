package ratings

import (
	"context"
	"errors"
	"project/ca/app/middlewares"
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

// func (uc *RateUseCase) GetAllRate(c context.Context, id int) (Ratings, error) {
// 	ctx, error := context.WithTimeout(c, uc.contextTimeout)
// 	defer error()

// 	Rate, err := uc.Repo.GetAllRate(ctx, id)
// 	if err != nil {
// 		return Ratings{}, err
// 	}

// 	return Rate, nil
// }

func (uc *RateUseCase) Detail(c context.Context, movie int, user int) (res Ratings, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	rate, err := uc.Repo.Detail(ctx, movie, user)
	if err != nil {
		return Ratings{}, err
	}
	return rate, nil
}

func (uc *RateUseCase) Delete(c context.Context, domain Ratings) error {

	if domain.MovieId == 0 {
		return errors.New("mohon isi ID Movie")
	}
	if domain.UserId == 0 {
		return errors.New("mohon isi ID User")
	}

	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	_, err := uc.Repo.Detail(ctx, domain.MovieId, domain.UserId)
	if err != nil {
		return err
	}
	domain.UpdatedAt = time.Now()

	err = uc.Repo.Delete(ctx, domain.MovieId, domain.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (uc *RateUseCase) Update(c context.Context, domain Ratings) error {

	if domain.MovieId == 0 {
		return errors.New("mohon isi ID Movie")
	}
	if domain.UserId == 0 {
		return errors.New("mohon isi ID User")
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

	if domain.MovieId == 0 {
		return Ratings{}, errors.New("mohon isi ID Movie")
	}
	if domain.UserId == 0 {
		return Ratings{}, errors.New("mohon isi ID User")
	}
	if domain.Rate == 0 {
		return Ratings{}, errors.New("mohon isi Ratings")
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
