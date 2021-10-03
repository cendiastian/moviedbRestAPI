package payments

import (
	"context"
	"project/app/middlewares"
	resp "project/business"
	"time"
)

type PaymentUsecase struct {
	ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewPaymentUsecase(repo Repository, timeout time.Duration) Usecase {
	return &PaymentUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *PaymentUsecase) GetAll(c context.Context) ([]Payment_method, error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	pay, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Payment_method{}, resp.ErrInternalServer
	}
	if len(pay) == 0 {
		return []Payment_method{}, resp.ErrNotFound
	}
	return pay, nil
}

func (uc *PaymentUsecase) Detail(c context.Context, id int) (res Payment_method, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	if id == 0 {
		return Payment_method{}, resp.ErrFillData
	}

	pay, err := uc.Repo.Detail(ctx, id)
	if err != nil {
		return Payment_method{}, resp.ErrNotFound
	}

	return pay, nil

}
func (uc *PaymentUsecase) Delete(c context.Context, id int) (Payment_method, error) {

	if id == 0 {
		return Payment_method{}, resp.ErrFillData
	}

	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	_, err := uc.Repo.Detail(ctx, id)
	if err != nil {
		return Payment_method{}, resp.ErrNotFound
	}
	del, err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return Payment_method{}, resp.ErrInternalServer
	}

	return del, nil
}

func (uc *PaymentUsecase) Update(c context.Context, domain Payment_method) (Payment_method, error) {

	if domain.Id == 0 {
		return Payment_method{}, resp.ErrFillData
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	_, err := uc.Repo.Detail(ctx, domain.Id)
	if err != nil {
		return Payment_method{}, resp.ErrNotFound
	}
	domain.UpdatedAt = time.Now()

	upt, err := uc.Repo.Update(ctx, domain)
	if err != nil {
		return Payment_method{}, resp.ErrInternalServer
	}

	return upt, nil

}

func (uc *PaymentUsecase) Register(c context.Context, domain Payment_method) (Payment_method, error) {

	if domain.Name == "" {
		return Payment_method{}, resp.ErrFillData
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	pay, err := uc.Repo.Register(ctx, domain)
	if err != nil {
		return Payment_method{}, resp.ErrInternalServer
	}

	return pay, nil

}
