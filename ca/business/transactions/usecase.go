package transactions

import (
	"context"
	"errors"
	"project/ca/app/middlewares"
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
		return []Payment_method{}, err
	}

	return pay, nil
}

func (uc *PaymentUsecase) Detail(c context.Context, id int) (res Payment_method, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	pay, err := uc.Repo.Detail(ctx, id)
	if err != nil {
		return Payment_method{}, err
	}

	return pay, nil

}
func (uc *PaymentUsecase) Delete(c context.Context, id int) (Payment_method, error) {

	if id == 0 {
		return Payment_method{}, errors.New("mohon isi ID")
	}

	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	del, err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return Payment_method{}, err
	}

	return del, nil
}

func (uc *PaymentUsecase) Update(c context.Context, domain Payment_method) (err error) {

	if domain.Id == 0 {
		return errors.New("mohon isi ID")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	err = uc.Repo.Update(ctx, domain.Id, domain.Name, domain.Status)
	if err != nil {
		return err
	}

	return nil

}

func (uc *PaymentUsecase) Register(c context.Context, domain Payment_method) (Payment_method, error) {

	if domain.Name == "" {
		return Payment_method{}, errors.New("mohon isi Nama")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	pay, err := uc.Repo.Register(ctx, domain.Name, domain.Status)
	if err != nil {
		return Payment_method{}, err
	}

	return pay, nil

}

func (uc *PaymentUsecase) CreateTransaction(c context.Context, domain Transaction) (Transaction, error) {

	if domain.Payment_method_id == 0 {
		return Transaction{}, errors.New("mohon isi Nama")
	}
	if domain.User_Id == 0 {
		return Transaction{}, errors.New("mohon isi Nama")
	}
	if domain.Plan_Id == 0 {
		return Transaction{}, errors.New("mohon isi Nama")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	pay, err := uc.Repo.CreateTransaction(ctx, domain.Payment_method_id, domain.User_Id, domain.Plan_Id)
	if err != nil {
		return Transaction{}, err
	}

	return pay, nil

}
