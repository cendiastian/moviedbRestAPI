package movies

import (
	"context"
	"errors"
	"fmt"
	"project/ca/app/middlewares"
	"project/ca/business/genres"
	"project/ca/business/omdb"

	// "project/ca/business/ratings"
	"strings"
	"time"
)

type Usecases struct {
	ConfigJWT middlewares.ConfigJWT
	Repo      Repository
	// RepoRate       ratings.Repository
	RepoGenre      genres.Repository
	RepoAPI        omdb.Repository
	contextTimeout time.Duration
}

func NewMovieUsecase(repo Repository, timeout time.Duration, repogenre genres.Repository, repoapi omdb.Repository) Usecase {
	return &Usecases{
		// ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
		RepoGenre:      repogenre,
		RepoAPI:        repoapi,
	}
}

func (uc *Usecases) CreateMovie(c context.Context, ImdbId string) (Movie, error) {
	var movie Movie
	var genre genres.Genre

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	movie.UpdatedAt = time.Now()

	API, err := uc.RepoAPI.GetAPI(ctx, ImdbId)
	if err != nil {
		return Movie{}, err
	}
	fmt.Println(API)
	GenreName := strings.Split(API.Genre, ", ")
	fmt.Println(GenreName)
	for _, v := range GenreName {
		genre.Name = v
		fmt.Println(genre.Name)
		scan, err := uc.RepoGenre.FirstOrCreate(ctx, genre.Name)
		if err != nil {
			return Movie{}, err
		}
		movie.Genre = append(movie.Genre, scan)
		if err != nil {
			return Movie{}, err
		}
	}
	fmt.Println(movie.Genre)
	movie = Movie{
		Title:  API.Title,
		ImdbId: API.ImdbId,
		Year:   API.Year,
		Type:   API.Type,
		Poster: API.Poster,
		Genre:  movie.Genre,
		Writer: API.Writer,
		Actors: API.Actors,
	}
	fmt.Println(movie.Genre)
	movie, err = uc.Repo.CreateMovie(ctx, movie, movie.Genre)
	if err != nil {
		return Movie{}, err
	}

	return movie, nil

}

func (uc *Usecases) MovieDetail(c context.Context, id int) (res Movie, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	// rate, err := uc.RepoRate.GetAllRate(ctx, id)
	// if err != nil {
	// 	return []Movie{}, err
	// }
	fmt.Println(id)
	movie, err := uc.Repo.MovieDetail(ctx, id)
	// fmt.Println(movie)
	if err != nil {
		return Movie{}, err
	}

	return movie, nil
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

	_, err := uc.Repo.MovieDetail(ctx, id)
	if err != nil {
		return Movie{}, err
	}
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
	_, err = uc.Repo.MovieDetail(ctx, domain.Id)
	if err != nil {
		return err
	}
	domain.UpdatedAt = time.Now()

	err = uc.Repo.UpdateMovie(ctx, domain)
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
