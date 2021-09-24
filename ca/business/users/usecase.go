package users

import (
	"context"
	"errors"
	"project/ca/app/middlewares"
	"project/ca/helpers/encrypt"
	"time"
)

type UserUsecase struct {
	ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		// ConfigJWT:      configJWT,
		Repo:           repo,
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

	var err error
	domain.Password, err = encrypt.Hash(domain.Password)

	// user, err := uc.Repo.Login(ctx, domain.Email, domain.Password)
	if err != nil {
		return User{}, err
	}

	user, err := uc.Repo.Login(ctx, domain.Email, domain.Password)
	if err != nil {
		return User{}, err
	}

	user.Token, err = uc.ConfigJWT.GenerateToken(user.Id)
	if err != nil {
		return User{}, err
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
func (uc *UserUsecase) Delete(c context.Context, domain User) (User, error) {

	if domain.Id == 0 {
		return User{}, errors.New("mohon isi ID")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	exist, err := uc.Repo.UserDetail(ctx, domain.Id)
	if err != nil {
		return User{}, err
	}
	if exist == (User{}) {
		return User{}, err
	}

	user, err := uc.Repo.Delete(ctx, domain.Id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (uc *UserUsecase) Update(c context.Context, domain User) (User, error) {

	if domain.Id == 0 {
		return User{}, errors.New("mohon isi ID")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	user, err := uc.Repo.Update(ctx, domain.Id, domain.Email, domain.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil

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

	user, err := uc.Repo.Register(ctx, domain.Name, domain.Email, domain.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil

}
