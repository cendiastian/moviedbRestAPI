package transactions

import (
	"context"
	"fmt"
	"project/app/middlewares"
	resp "project/business"
	"project/business/premium"

	"time"
)

type TransUsecase struct {
	ConfigJWT middlewares.ConfigJWT
	Repo      Repository
	RepoPro   premium.Repository
	// RepoSubs  subscription.Repository

	contextTimeout time.Duration
}

func NewTransUsecase(repo Repository, timeout time.Duration, repoPro premium.Repository /* repoSubs subscription.Repository*/) Usecase {
	return &TransUsecase{
		Repo:           repo,
		contextTimeout: timeout,
		RepoPro:        repoPro,
		// RepoSubs:       repoSubs,
	}
}

func (uc *TransUsecase) CreateTransaction(c context.Context, domain Transaction) (Transaction, error) {

	if domain.Payment_method_id == 0 || domain.User_Id == 0 || domain.Plan_Id == 0 {
		return Transaction{}, resp.ErrFillData
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	fmt.Println(domain.Plan_Id)
	pay, err := uc.Repo.CreateTransaction(ctx, domain)
	if err != nil {
		return Transaction{}, resp.ErrInternalServer
	}
	// Sub, err := uc.RepoSubs.Detail(ctx, domain.Plan_Id)
	// fmt.Println(Sub)
	// if err != nil {
	// 	return Transaction{}, resp.ErrNotFound
	// }

	domain.UpdatedAt = time.Now()
	pro := premium.Premium{
		UserId:    domain.User_Id,
		Type:      true,
		Expired:   pay.Subscription_Plan.Exp,
		UpdatedAt: domain.UpdatedAt,
	}
	_, err = uc.RepoPro.Save(ctx, pro)
	if err != nil {
		return Transaction{}, resp.ErrInternalServer
	}

	return pay, nil

}
func (uc *TransUsecase) DetailTrans(c context.Context, id int) (res Transaction, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	if id == 0 {
		return Transaction{}, resp.ErrFillData
	}

	trans, err := uc.Repo.DetailTrans(ctx, id)
	if err != nil {
		return Transaction{}, resp.ErrNotFound
	}

	return trans, nil

}
