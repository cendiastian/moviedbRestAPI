package movies

import (
	"context"
	"errors"
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
func (uc *Usecases) FilterGenre(c context.Context, genre string) ([]Movie, error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	fmt.Println(genre)
	movie, err := uc.Repo.FilterGenre(ctx, genre)
	if err != nil {
		return []Movie{}, err
	}

	return movie, nil
}
func (uc *Usecases) FilterOrder(c context.Context, order string) ([]Movie, error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	fmt.Println(order)
	movie, err := uc.Repo.FilterOrder(ctx, order)
	if err != nil {
		return []Movie{}, err
	}

	return movie, nil
}

func (uc *Usecases) GetAllMovie(c context.Context) ([]Movie, error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	movie, err := uc.Repo.GetAllMovie(ctx)
	if err != nil {
		return []Movie{}, err
	}

	return movie, nil
}

func (uc *Usecases) DeleteMovie(c context.Context, id int) (Movie, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	del, err := uc.Repo.DeleteMovie(ctx, id)
	if err != nil {
		return Movie{}, err
	}

	return del, nil
}

func (uc *Usecases) UpdateMovie(c context.Context, domain Movie) (err error) {

	if domain.Id == 0 {
		return errors.New("mohon isi ID")
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	domain.UpdatedAt = time.Now()

	err = uc.Repo.UpdateMovie(ctx, domain.Id, domain.Title, domain.Type)
	if err != nil {
		return err
	}

	return nil

}

func (uc *Usecases) DeleteAll(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	err := uc.Repo.DeleteAll(ctx)
	if err != nil {
		return err
	}

	return nil
}
