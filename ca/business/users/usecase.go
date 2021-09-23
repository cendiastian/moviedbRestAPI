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

// func (a *UserUsecase) fillAuthorDetails(c context.Context, data []User) ([]User, error) {
// 	g, ctx := errgroup.WithContext(c)
//
// 	// Get the author's id
// 	mapAuthors := map[int64]domain.Author{}
//
// 	for _, article := range data {
// 		mapAuthors[article.Author.ID] = domain.Author{}
// 	}
// 	// Using goroutine to fetch the author's detail
// 	chanAuthor := make(chan domain.Author)
// 	for authorID := range mapAuthors {
// 		authorID := authorID
// 		g.Go(func() error {
// 			res, err := a.authorRepo.GetByID(ctx, authorID)
// 			if err != nil {
// 				return err
// 			}
// 			chanAuthor <- res
// 			return nil
// 		})
// 	}
//
// 	go func() {
// 		err := g.Wait()
// 		if err != nil {
// 			logrus.Error(err)
// 			return
// 		}
// 		close(chanAuthor)
// 	}()
//
// 	for author := range chanAuthor {
// 		if author != (domain.Author{}) {
// 			mapAuthors[author.ID] = author
// 		}
// 	}
//
// 	if err := g.Wait(); err != nil {
// 		return nil, err
// 	}
//
// 	// merge the author's data
// 	for index, item := range data {
// 		if a, ok := mapAuthors[item.Author.ID]; ok {
// 			data[index].Author = a
// 		}
// 	}
// 	return data, nil
// }
//
// func (a *UserUsecase) Fetch(c context.Context, cursor string, num int64) (res []User, nextCursor string, err error) {
// 	if num == 0 {
// 		num = 10
// 	}
//
// 	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
// 	defer cancel()
//
// 	res, nextCursor, err = a.Repo.Fetch(ctx, cursor, num)
// 	if err != nil {
// 		return nil, "", err
// 	}
//
// 	res, err = a.fillAuthorDetails(ctx, res)
// 	if err != nil {
// 		nextCursor = ""
// 	}
// 	return
// }

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
