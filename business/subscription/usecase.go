package subscription

import (
	"context"
	"fmt"
	"project/app/middlewares"
	resp "project/business"
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
		return []SubcriptionPlan{}, resp.ErrInternalServer
	}
	if len(Subs) == 0 {
		return []SubcriptionPlan{}, resp.ErrNotFound
	}

	return Subs, nil
}

func (uc *subsUseCase) Detail(c context.Context, id int) (res SubcriptionPlan, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	if id == 0 {
		return SubcriptionPlan{}, resp.ErrFillData
	}

	Sub, err := uc.Repo.Detail(ctx, id)
	if err != nil {
		return SubcriptionPlan{}, resp.ErrNotFound
	}

	return Sub, nil

}
func (uc *subsUseCase) Delete(c context.Context, id int) error {

	if id == 0 {
		return resp.ErrFillData
	}

	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	_, err := uc.Repo.Detail(ctx, id)
	if err != nil {
		return resp.ErrNotFound
	}
	_, err = uc.Repo.Delete(ctx, id)
	if err != nil {
		return resp.ErrInternalServer
	}

	return nil
}

func (uc *subsUseCase) Update(c context.Context, domain SubcriptionPlan) (err error) {

	if domain.Id == 0 {
		return resp.ErrFillData
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	_, err = uc.Repo.Detail(ctx, domain.Id)
	if err != nil {
		return resp.ErrNotFound
	}
	domain.UpdatedAt = time.Now()

	err = uc.Repo.Update(ctx, domain)
	if err != nil {
		return resp.ErrInternalServer
	}

	return nil

}

func (uc *subsUseCase) CreatePlan(c context.Context, domain SubcriptionPlan) (SubcriptionPlan, error) {

	if domain.Name == "" || domain.Expired == "" || domain.Exp.IsZero() || domain.Price == 0 {
		return SubcriptionPlan{}, resp.ErrFillData
	}
	fmt.Println(domain.Exp)
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	subs, err := uc.Repo.CreatePlan(ctx, domain)
	if err != nil {
		return SubcriptionPlan{}, resp.ErrInternalServer
	}

	return subs, nil

}
