package subscription

import (
	"context"
	"errors"
	"fmt"
	"project/app/middlewares"
	"time"
)

type subsUseCase struct {
	ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewSubsUsecase(repo Repository, timeout time.Duration) Usecase {
	return &subsUseCase{
		// ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *subsUseCase) GetAll(c context.Context) ([]SubcriptionPlan, error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	Subs, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []SubcriptionPlan{}, err
	}

	return Subs, nil
}

func (uc *subsUseCase) Detail(c context.Context, id int) (res SubcriptionPlan, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	Sub, err := uc.Repo.Detail(ctx, id)
	if err != nil {
		return SubcriptionPlan{}, err
	}

	return Sub, nil

}
func (uc *subsUseCase) Delete(c context.Context, id int) (SubcriptionPlan, error) {

	if id == 0 {
		return SubcriptionPlan{}, errors.New("mohon isi ID")
	}

	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	_, err := uc.Repo.Detail(ctx, id)
	if err != nil {
		return SubcriptionPlan{}, err
	}
	del, err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return SubcriptionPlan{}, err
	}

	return del, nil
}

func (uc *subsUseCase) Update(c context.Context, domain SubcriptionPlan) (err error) {

	if domain.Id == 0 {
		return errors.New("mohon isi ID")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	_, err = uc.Repo.Detail(ctx, domain.Id)
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

func (uc *subsUseCase) CreatePlan(c context.Context, domain SubcriptionPlan) (SubcriptionPlan, error) {

	if domain.Name == "" {
		return SubcriptionPlan{}, errors.New("mohon isi Nama")
	}
	if domain.Expired == "" {
		return SubcriptionPlan{}, errors.New("mohon isi Expired")
	}
	fmt.Println(domain.Exp)
	if domain.Exp.IsZero() {
		return SubcriptionPlan{}, errors.New("mohon isi Exp")
	}
	if domain.Price == 0 {
		return SubcriptionPlan{}, errors.New("mohon isi Price")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	subs, err := uc.Repo.CreatePlan(ctx, domain)
	if err != nil {
		return SubcriptionPlan{}, err
	}

	return subs, nil

}
