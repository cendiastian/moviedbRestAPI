package users

import (
	"context"
	"errors"
	"fmt"
	"project/app/middlewares"
	"project/business/premium"
	"project/helpers/encrypt"
	"time"
)

type UserUsecase struct {
	ConfigJWT      *middlewares.ConfigJWT
	Repo           Repository
	RepoPro        premium.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, repoPro premium.Repository, timeout time.Duration, configJWT *middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		RepoPro:        repoPro,
		contextTimeout: timeout,
	}
}

func (uc *UserUsecase) GetAll(c context.Context) ([]User, error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	user, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []User{}, err
	}

	return user, nil
}

// func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {
func (uc *UserUsecase) Login(ctx context.Context, domain User) (User, error) {

	if domain.Email == "" {
		return User{}, errors.New("mohon isi email")
	}

	if domain.Password == "" {
		return User{}, errors.New("mohon isi password")
	}

	// var err error
	// domain.Password, err = encrypt.Hash(domain.Password)

	// user, err := uc.Repo.Login(ctx, domain.Email, domain.Password)
	// if err != nil {
	// 	return User{}, err
	// }

	user, err := uc.Repo.Login(ctx, domain)
	if err != nil {
		return User{}, err
	}

	err = encrypt.CheckPassword(domain.Password, user.Password)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}

	user.Token, err = uc.ConfigJWT.GenerateToken(user.Id)
	// fmt.Println(user.Token)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	today := time.Now()
	pro, err := uc.RepoPro.Detail(ctx, user.Id)
	if err != nil {
		return User{}, err
	}
	if pro.Expired.Before(today) && !pro.Expired.IsZero() {
		res := premium.Premium{
			UserId:    pro.UserId,
			Type:      false,
			Expired:   time.Time{},
			UpdatedAt: time.Now(),
		}
		user.Premium, err = uc.RepoPro.Save(ctx, res)
		if err != nil {
			return User{}, err
		}
	}

	return user, nil
}

func (uc *UserUsecase) UserDetail(c context.Context, id int) (res User, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	user, err := uc.Repo.UserDetail(ctx, id)
	if err != nil {
		return User{}, err
	}

	return user, nil

}
func (uc *UserUsecase) Delete(c context.Context, domain User) (err error) {

	if domain.Id == 0 {
		return errors.New("mohon isi ID")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	exist, err := uc.Repo.UserDetail(ctx, domain.Id)
	if err != nil {
		return err
	}
	if exist.Id == 0 {
		return err
	}

	err = uc.Repo.Delete(ctx, domain.Id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUsecase) Update(c context.Context, domain User) (err error) {

	if domain.Id == 0 {
		return errors.New("mohon isi ID")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	_, err = uc.Repo.UserDetail(ctx, domain.Id)
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

func (uc *UserUsecase) Register(c context.Context, domain User) (User, error) {

	if domain.Name == "" {
		return User{}, errors.New("mohon isi Nama")
	}
	if domain.Email == "" {
		return User{}, errors.New("mohon isi Email")
	}
	if domain.Password == "" {
		return User{}, errors.New("mohon isi Password")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()
	// domain.Password, _ = encrypt.Hash(domain.Password)
	user, err := uc.Repo.Register(ctx, domain)
	if err != nil {
		return User{}, err
	}
	res := premium.Premium{
		UserId:  user.Id,
		Type:    false,
		Expired: time.Time{},
	}
	_, err = uc.RepoPro.Save(ctx, res)
	if err != nil {
		return User{}, err
	}

	return user, nil

}
