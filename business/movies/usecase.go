package movies

import (
	"context"
	"fmt"
	"project/app/middlewares"
	resp "project/business"
	"project/business/genres"
	"project/business/omdb"
	"project/business/premium"

	// "project/business/ratings"
	"strings"
	"time"
)

type Usecases struct {
	ConfigJWT middlewares.ConfigJWT
	Repo      Repository
	// RepoRate       ratings.Repository
	RepoGenre      genres.Repository
	RepoAPI        omdb.Repository
	RepoPro        premium.Repository
	contextTimeout time.Duration
}

func NewMovieUsecase(repo Repository, timeout time.Duration, repogenre genres.Repository, repoapi omdb.Repository, repoPro premium.Repository) Usecase {
	return &Usecases{
		// ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
		RepoGenre:      repogenre,
		RepoAPI:        repoapi,
		RepoPro:        repoPro,
	}
}

func (uc *Usecases) CreateMovie(c context.Context, ImdbId string) (Movie, error) {
	var movie Movie
	var genre genres.Genre

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	movie.UpdatedAt = time.Now()
	if ImdbId == "" {
		return Movie{}, resp.ErrFillData
	}

	API, err := uc.RepoAPI.GetAPI(ctx, ImdbId)
	if err != nil {
		return Movie{}, resp.ErrAPIFound
	}
	fmt.Println(API)
	GenreName := strings.Split(API.Genre, ", ")
	fmt.Println(GenreName)
	for _, v := range GenreName {
		genre.Name = v
		fmt.Println(genre.Name)
		scan, err := uc.RepoGenre.FirstOrCreate(ctx, genre.Name)
		if err != nil {
			return Movie{}, resp.ErrInternalServer
		}
		movie.Genre = append(movie.Genre, scan)
		if err != nil {
			return Movie{}, resp.ErrInternalServer
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
		return Movie{}, resp.ErrInternalServer
	}

	return movie, nil

}

func (uc *Usecases) MovieDetail(c context.Context, id int, user int) (res Movie, err error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	if id == 0 || user == 0 {
		return Movie{}, resp.ErrFillData
	}
	pro, err := uc.RepoPro.Detail(ctx, user)
	if err != nil {
		return Movie{}, resp.ErrNotFound
	}
	if !pro.Type {
		return Movie{}, resp.ErrNotProFound
	}
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
	if title == "" {
		return []Movie{}, resp.ErrFillData
	}
	movie, err := uc.Repo.SearchMovie(ctx, title)
	if err != nil {
		return []Movie{}, resp.ErrInternalServer
	}
	if len(movie) == 0 {
		return []Movie{}, resp.ErrNotFound
	}
	return movie, nil
}
func (uc *Usecases) FilterGenre(c context.Context, genre string) ([]Movie, error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	fmt.Println(genre)
	if genre == "" {
		return []Movie{}, resp.ErrFillData
	}
	movie, err := uc.Repo.FilterGenre(ctx, genre)
	if err != nil {
		return []Movie{}, resp.ErrInternalServer
	}
	if len(movie) == 0 {
		return []Movie{}, resp.ErrNotFound
	}
	return movie, nil
}
func (uc *Usecases) FilterOrder(c context.Context, order string) ([]Movie, error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	fmt.Println(order)
	if order == "" {
		return []Movie{}, resp.ErrFillData
	}
	movie, err := uc.Repo.FilterOrder(ctx, order)
	if err != nil {
		return []Movie{}, resp.ErrInternalServer
	}
	if len(movie) == 0 {
		return []Movie{}, resp.ErrNotFound
	}
	return movie, nil
}

func (uc *Usecases) GetAllMovie(c context.Context) ([]Movie, error) {
	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()

	movie, err := uc.Repo.GetAllMovie(ctx)
	if err != nil {
		return []Movie{}, resp.ErrInternalServer
	}
	if len(movie) == 0 {
		return []Movie{}, resp.ErrNotFound
	}
	return movie, nil
}

func (uc *Usecases) DeleteMovie(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	if id == 0 {
		return resp.ErrFillData
	}
	_, err := uc.Repo.MovieDetail(ctx, id)
	if err != nil {
		return resp.ErrNotFound
	}
	err = uc.Repo.DeleteMovie(ctx, id)
	if err != nil {
		return resp.ErrInternalServer
	}

	return nil
}

func (uc *Usecases) UpdateMovie(c context.Context, domain Movie) (err error) {

	if domain.Id == 0 {
		return resp.ErrFillData
	}

	ctx, error := context.WithTimeout(c, uc.contextTimeout)
	defer error()
	_, err = uc.Repo.MovieDetail(ctx, domain.Id)
	if err != nil {
		return resp.ErrNotFound
	}
	domain.UpdatedAt = time.Now()

	err = uc.Repo.UpdateMovie(ctx, domain)
	if err != nil {
		return resp.ErrInternalServer
	}

	return nil

}

func (uc *Usecases) DeleteAll(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	movie, err := uc.Repo.GetAllMovie(ctx)
	if err != nil {
		return resp.ErrInternalServer
	}
	if len(movie) == 0 {
		return resp.ErrNotFound
	}

	err = uc.Repo.DeleteAll(ctx)
	if err != nil {
		return err
	}

	return nil
}
