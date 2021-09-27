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

// func (uc *RateUseCase) GetAll(c context.Context) ([]Rating, error) {
// 	ctx, error := context.WithTimeout(c, uc.contextTimeout)
// 	defer error()

// 	Rate, err := uc.Repo.GetAll(ctx)
// 	if err != nil {
// 		return []Rating{}, err
// 	}

// 	return Rate, nil
// }

// func (uc *RateUseCase) Detail(c context.Context, id int) (res Rating, err error) {
// 	ctx, error := context.WithTimeout(c, uc.contextTimeout)
// 	defer error()
// 	Sub, err := uc.Repo.Detail(ctx, id)
// 	if err != nil {
// 		return Rating{}, err
// 	}
// 	return Sub, nil
// }

func (uc *RateUseCase) Delete(c context.Context, domain Rating) error {

	if domain.Movie_Id == 0 {
		return errors.New("mohon isi ID Movie")
	}
	if domain.User_Id == 0 {
		return errors.New("mohon isi ID User")
	}

	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	domain.UpdatedAt = time.Now()

	err := uc.Repo.Update(ctx, domain.Movie_Id, domain.User_Id, domain.Rate)
	if err != nil {
		return err
	}

	return nil
}

func (uc *RateUseCase) Update(c context.Context, domain Rating) error {

	if domain.Movie_Id == 0 {
		return errors.New("mohon isi ID Movie")
	}
	if domain.User_Id == 0 {
		return errors.New("mohon isi ID User")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	err := uc.Repo.Update(ctx, domain.Movie_Id, domain.User_Id, domain.Rate)
	if err != nil {
		return err
	}

	return nil

}

func (uc *RateUseCase) Create(c context.Context, domain Rating) (Rating, error) {

	if domain.Movie_Id == 0 {
		return Rating{}, errors.New("mohon isi ID Movie")
	}
	if domain.User_Id == 0 {
		return Rating{}, errors.New("mohon isi ID User")
	}
	if domain.Rate == 0 {
		return Rating{}, errors.New("mohon isi Rating")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	Rate, err := uc.Repo.Create(ctx, domain.Movie_Id, domain.User_Id, domain.Rate)
	if err != nil {
		return Rating{}, err
	}

	return Rate, nil

}
