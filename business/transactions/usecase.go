package transactions

import (
	"context"
	"errors"
	"fmt"
	"project/app/middlewares"
	"project/business/premium"
	"project/business/subscription"
	"time"
)

type TransUsecase struct {
	ConfigJWT middlewares.ConfigJWT
	Repo      Repository
	RepoPro   premium.Repository
	RepoSubs  subscription.Repository

	contextTimeout time.Duration
}

func NewTransUsecase(repo Repository, timeout time.Duration, repoPro premium.Repository, repoSubs subscription.Repository) Usecase {
	return &TransUsecase{
		Repo:           repo,
		contextTimeout: timeout,
		RepoPro:        repoPro,
		RepoSubs:       repoSubs,
	}
}

func (uc *TransUsecase) CreateTransaction(c context.Context, domain Transaction) (Transaction, error) {

	if domain.Payment_method_id == 0 {
		return Transaction{}, errors.New("mohon isi PaymentMethod")
	}
	if domain.User_Id == 0 {
		return Transaction{}, errors.New("mohon isi UserId")
	}
	if domain.Plan_Id == 0 {
		return Transaction{}, errors.New("mohon isi PlanId")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	fmt.Println(domain.Plan_Id)
	Sub, err := uc.RepoSubs.Detail(ctx, domain.Plan_Id)
	fmt.Println(Sub)
	if err != nil {
		return Transaction{}, err
	}

	domain.UpdatedAt = time.Now()
	pro := premium.Premium{
		UserId:    domain.User_Id,
		Type:      true,
		Expired:   Sub.Exp,
		UpdatedAt: domain.UpdatedAt,
	}
	_, err = uc.RepoPro.Save(ctx, pro)
	if err != nil {
		return Transaction{}, err
	}
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
