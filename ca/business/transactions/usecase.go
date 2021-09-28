package transactions

import (
	"context"
	"errors"
	"project/ca/app/middlewares"

	// "project/ca/business/subscription"
	"time"
)

type TransUsecase struct {
	ConfigJWT middlewares.ConfigJWT
	Repo      Repository

	contextTimeout time.Duration
}

func NewTransUsecase(repo Repository, timeout time.Duration) Usecase {
	return &TransUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *TransUsecase) CreateTransaction(c context.Context, domain Transaction) (Transaction, error) {

	if domain.Payment_method_id == 0 {
		return Transaction{}, errors.New("mohon isi Nama")
	}
	if domain.User_Id == 0 {
		return Transaction{}, errors.New("mohon isi Nama")
	}
	if domain.Plan_Id == 0 {
		return Transaction{}, errors.New("mohon isi Nama")
	}

	// Sub, err := uc.RepoSubs.Detail(ctx, id)

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	pay, err := uc.Repo.CreateTransaction(ctx, domain)
	if err != nil {
		return Transaction{}, err
	}

	return pay, nil

}
func (uc *TransUsecase) DetailTrans(c context.Context, id int) (res Transaction, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	trans, err := uc.Repo.DetailTrans(ctx, id)
	if err != nil {
		return Transaction{}, err
	}

	return trans, nil

}
