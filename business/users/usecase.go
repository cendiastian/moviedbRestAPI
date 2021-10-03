package users

import (
	"context"
	resp "project/business"
	"project/business/premium"
	"time"
)

type UserUsecase struct {
	// ConfigJWT      *middlewares.ConfigJWT
	Repo           Repository
	RepoPro        premium.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, repoPro premium.Repository, timeout time.Duration /*configJWT *middlewares.ConfigJWT*/) Usecase {
	return &UserUsecase{
		// ConfigJWT:      configJWT,
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
		return []User{}, resp.ErrInternalServer
	}
	if len(user) == 0 {
		return []User{}, resp.ErrNotFound
	}

	return user, nil
}

// func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {
func (uc *UserUsecase) Login(ctx context.Context, domain User) (User, error) {

	if domain.Email == "" || domain.Password == "" {
		return User{}, resp.ErrFillData
	}

	user, err := uc.Repo.Login(ctx, domain)
	if err != nil {
		return User{}, resp.ErrUsernamePasswordNotFound
	}

	// err = encrypt.CheckPassword(domain.Password, user.Password)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return User{}, resp.ErrUsernamePasswordNotFound
	// }

	// user.Token, err = uc.ConfigJWT.GenerateToken(user.Id)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return User{}, resp.ErrInternalServer
	// }
	today := time.Now()
	pro, err := uc.RepoPro.Detail(ctx, user.Id)
	if err != nil {
		return User{}, resp.ErrNotFound
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
			return User{}, resp.ErrInternalServer
		}
	}

	return user, nil
}

func (uc *UserUsecase) UserDetail(c context.Context, id int) (res User, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	if id == 0 {
		return User{}, resp.ErrFillData
	}

	user, err := uc.Repo.UserDetail(ctx, id)
	if err != nil {
		return User{}, resp.ErrNotFound
	}

	return user, nil

}
func (uc *UserUsecase) Delete(c context.Context, id int) (res User, err error) {

	if id == 0 {
		return User{}, resp.ErrFillData
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	_, err = uc.Repo.UserDetail(ctx, id)
	if err != nil {
		return User{}, resp.ErrNotFound
	}

	del, err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return User{}, resp.ErrNotFound
	}

	return del, nil
}

func (uc *UserUsecase) Update(c context.Context, domain User) (res User, err error) {

	if domain.Id == 0 {
		return User{}, resp.ErrFillData
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	_, err = uc.Repo.UserDetail(ctx, domain.Id)
	if err != nil {
		return User{}, resp.ErrNotFound
	}
	domain.UpdatedAt = time.Now()

	up, err := uc.Repo.Update(ctx, domain)
	if err != nil {
		return User{}, resp.ErrNotFound
	}

	return up, nil

}

func (uc *UserUsecase) Register(c context.Context, domain User) (User, error) {

	if domain.Name == "" || domain.Email == "" || domain.Password == "" {
		return User{}, resp.ErrFillData
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()
	// domain.Password, _ = encrypt.Hash(domain.Password)
	user, err := uc.Repo.Register(ctx, domain)
	if err != nil {
		return User{}, resp.ErrInternalServer
	}
	res := premium.Premium{
		UserId:  user.Id,
		Type:    false,
		Expired: time.Time{},
	}
	_, err = uc.RepoPro.Save(ctx, res)
	if err != nil {
		return User{}, resp.ErrInternalServer
	}

	return user, nil

}
