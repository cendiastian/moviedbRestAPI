package movies

import (
	"context"
	"fmt"
	"project/ca/app/middlewares"
	"time"
)

type Usecases struct {
	ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewMovieUsecase(repo Repository, timeout time.Duration) Usecase {
	return &Usecases{
		// ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *Usecases) GetAPI(c context.Context, ImdbId string) (res GetAPI, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	movie, err := uc.Repo.GetAPI(ctx, ImdbId)
	if err != nil {
		return GetAPI{}, err
	}

	return movie, nil
}

func (uc *Usecases) CreateMovieAPI(c context.Context, domain Movie) (Movie, error) {

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	movie, err := uc.Repo.CreateMovieAPI(ctx, domain.Title, domain.Year, domain.ImdbId, domain.Type, domain.Poster, domain.Genre, domain.Writer, domain.Actors)
	if err != nil {
		return Movie{}, err
	}

	return movie, nil

}

func (uc *Usecases) MovieDetail(c context.Context, id int) (res Movie, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	movie, err := uc.Repo.MovieDetail(ctx, id)
	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}

func (uc *Usecases) ScanGenre(c context.Context, domain Genre) (Genre, error) {

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	genre, err := uc.Repo.ScanGenre(ctx, domain.Name)
	if err != nil {
		return Genre{}, err
	}

	return genre, nil

}

func (uc *Usecases) SearchMovie(c context.Context, title string) ([]Movie, error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	fmt.Println(title)
	movie, err := uc.Repo.SearchMovie(ctx, title)
	if err != nil {
		return []Movie{}, err
	}

	return movie, nil
}

// func (uc *Usecases) CreateGenreAPI(c context.Context, domain Genre) (Genre, error) {

// 	ctx, error := context.WithTimeout(c, uc.contextTimeout)
// 	defer error()

// 	domain.UpdatedAt = time.Now()

// 	genre, err := uc.Repo.CreateGenreAPI(ctx, domain.Name)
// 	if err != nil {
// 		return Genre{}, err
// 	}

// 	return genre, nil

// }

// func (uc *MovieUsecase) GetAllMovie(c context.Context) ([]Movie, error) {
// 	ctx, error := context.WithTimeout(c, uc.contextTimeout)
// 	defer error()

// 	movie, err := uc.Repo.GetAllMovie(ctx)
// 	if err != nil {
// 		return []Movie{}, err
// 	}

// 	return movie, nil
// }

/*
// // func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {
// func (uc *UserUsecase) Login(ctx context.Context, domain User) (User, error) {

// 	if domain.Email == "" {
// 		return User{}, errors.New("mohon isi email")
// 	}

// 	if domain.Password == "" {
// 		return User{}, errors.New("mohon isi password")
// 	}

// 	var err error
// 	// domain.Password, err = encrypt.Hash(domain.Password)

// 	// user, err := uc.Repo.Login(ctx, domain.Email, domain.Password)
// 	if err != nil {
// 		return User{}, err
// 	}

// 	user, err := uc.Repo.Login(ctx, domain.Email, domain.Password)
// 	if err != nil {
// 		return User{}, err
// 	}

// 	// user.Token, err = uc.ConfigJWT.GenerateToken(user.Id)
// 	if err != nil {
// 		return User{}, err
// 	}

// 	return user, nil
// }

// func (uc *UserUsecase) Delete(c context.Context, domain User) (err error) {

// 	if domain.Id == 0 {
// 		return errors.New("mohon isi ID")
// 	}

// 	ctx, error := context.WithTimeout(c, uc.contextTimeout)
// 	defer error()

// 	exist, err := uc.Repo.UserDetail(ctx, domain.Id)
// 	if err != nil {
// 		return err
// 	}
// 	if exist == (User{}) {
// 		return err
// 	}

// 	err = uc.Repo.Delete(ctx, domain.Id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (uc *UserUsecase) Update(c context.Context, domain User) (err error) {

// 	if domain.Id == 0 {
// 		return errors.New("mohon isi ID")
// 	}

// 	ctx, error := context.WithTimeout(c, uc.contextTimeout)
// 	defer error()

// 	domain.UpdatedAt = time.Now()

// 	err = uc.Repo.Update(ctx, domain.Id, domain.Email, domain.Password)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
*/
