package movies

import (
	"context"
	"fmt"
	_genre_b "project/ca/business/genres"
	"project/ca/business/movies"

	// "project/ca/drivers/databases/ratings"

	"gorm.io/gorm"
)

type MysqlMovieRepository struct {
	Connect *gorm.DB
}

func NewMysqlMovieRepository(connect *gorm.DB) movies.Repository {
	return &MysqlMovieRepository{
		Connect: connect,
	}
}

func (rep *MysqlMovieRepository) CreateMovie(ctx context.Context, domain movies.Movie, array []_genre_b.Genre) (movies.Movie, error) {
	movie := FromDomain(domain)
	// genre := _genre_d.FromListDomainGenre(array)
	// var tes Movies
	// tes := Movies{
	// 	Id:      movie.Id,
	// 	Title:   movie.Title,
	// 	Year:    movie.Year,
	// 	ImdbId:  movie.ImdbId,
	// 	Type:    movie.Type,
	// 	Poster:  movie.Poster,
	// 	Genre:   genre,
	// 	Ratings: movie.Ratings,
	// 	Rating:  movie.Rating,
	// 	Writer:  movie.Writer,
	// 	Actors:  movie.Actors,
	// }
	result := rep.Connect.Create(&movie)
	if result.Error != nil {
		return movies.Movie{}, result.Error
	}

	return movie.ToDomainMovie(), nil
}

func (rep *MysqlMovieRepository) MovieDetail(ctx context.Context, id int) (movies.Movie, error) {
	var movie Movies
	// var Rating []ratings.Ratings
	// res := rep.Connect.Where("MovieId = ?", id).Find(Rating)
	// if res.Error != nil {
	// 	return movies.Movie{}, res.Error
	// }
	var total float32

	result := rep.Connect.Preload("Genre").Preload("Ratings.User").First(&movie, "id= ?", id)
	if result.Error != nil {
		fmt.Println("DB")
		return movies.Movie{}, result.Error
	}
	for _, v := range movie.Ratings {
		total += v.Rate
	}
	movie.Rating = total / float32(len(movie.Ratings))
	// movie.Ratings = (Ratings.Rate / len(Rating))
	return movie.ToDomainMovie(), nil
}

func (rep *MysqlMovieRepository) SearchMovie(ctx context.Context, title string) ([]movies.Movie, error) {
	var movie []Movies
	// fmt.Println(title)

	result := rep.Connect.Preload("Genre").Preload("Ratings.User").Where("title LIKE ?", title+"%").Find(&movie)
	if result.Error != nil {
		return []movies.Movie{}, result.Error
	}

	result = rep.Connect.Preload("Genre").Preload("Ratings.User").Where("title LIKE ?", title+"%").Find(&movie)
	if result.Error != nil {
		return []movies.Movie{}, result.Error
	}

	result = rep.Connect.Preload("Genre").Preload("Ratings.User").Where("title LIKE ?", "%"+title+"%").Find(&movie)
	if result.Error != nil {
		return []movies.Movie{}, result.Error
	}

	return ToListDomain(movie), nil
	// 	var result map[string]interface{}
	// db.Model(&User{}).First(&result, "id = ?", 1
}
func (rep *MysqlMovieRepository) FilterGenre(ctx context.Context, genre string) ([]movies.Movie, error) {
	var movie []Movies
	// var Genre Genres
	// result := rep.Connect.Where("name = ?", genre).Find(&Genre)
	// if result.Error != nil {
	// 	return []movies.Movie{}, result.Error
	// }
	// result := rep.Connect.Preload("Genre").Where("id IN (SELECT movies_id FROM movie_genres WHERE name IN ?)", []string{genre}).Find(&movie)
	result := rep.Connect.Preload("Genre").
		Joins("JOIN movie_genres on movie_genres.movies_id = movies.id JOIN genres on movie_genres.genres_id = genres.id AND genres.name = ? ",
			genre).Preload("Ratings.User").Find(&movie)
	if result.Error != nil {
		return []movies.Movie{}, result.Error
	}
	return ToListDomain(movie), nil
}
func (rep *MysqlMovieRepository) FilterOrder(ctx context.Context, order string) ([]movies.Movie, error) {
	fmt.Println("Tess" + order)
	var movie []Movies
	// res := rep.Connect.Find(&movie)
	// if res.Error != nil {
	// 	return []movies.Movie{}, res.Error
	// }
	if order == "oldest" {
		result := rep.Connect.Preload("Genre").Preload("Ratings.User").Find(&movie)
		if result.Error != nil {
			return []movies.Movie{}, result.Error
		}
	} else {
		result := rep.Connect.Preload("Genre").Preload("Ratings.User").Order("id desc").Find(&movie)
		if result.Error != nil {
			return []movies.Movie{}, result.Error
		}
	}
	return ToListDomain(movie), nil
}

func (rep *MysqlMovieRepository) GetAllMovie(ctx context.Context) ([]movies.Movie, error) {
	var movie []Movies
	result := rep.Connect.Preload("Genre").Preload("Ratings.User").Find(&movie)
	if result.Error != nil {
		return []movies.Movie{}, result.Error
	}
	return ToListDomain(movie), nil
}

func (rep *MysqlMovieRepository) DeleteAll(ctx context.Context) error {
	var movie []Movies

	result := rep.Connect.Find(&movie)
	if result.Error != nil {
		return result.Error
	}

	result = rep.Connect.Unscoped().Delete(&movie)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rep *MysqlMovieRepository) DeleteMovie(ctx context.Context, id int) (movies.Movie, error) {
	var movie Movies
	result := rep.Connect.Where("id = ?", id).Delete(&movie)

	if result.Error != nil {
		return movies.Movie{}, result.Error
	}

	return movie.ToDomainMovie(), nil
}

func (rep *MysqlMovieRepository) UpdateMovie(ctx context.Context, domain movies.Movie) error {
	movie := FromDomain(domain)
	result := rep.Connect.Where("id = ?", movie.Id).Updates(&Movies{Title: movie.Title, Type: movie.Type})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// func (rep *MysqlRepository) CreateGenreAPI(ctx context.Context, name string) (movies.Genre, error) {
// 	var genre Genres
// 	result := rep.Connect.Create(&genre)

// 	if result.Error != nil {
// 		return movies.Genre{}, result.Error
// 	}

// 	return genre.ToDomainGenre(), nil
// }

/*
func (rep *MysqlMovieRepository) Login(ctx context.Context, email string, password string) (users.User, error) {
	var user Users
	result := rep.Connect.First(&user, "email = ? AND password = ?",
		email, password)
	if result.Error != nil {
		return users.User{}, result.Error
	}
	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) Delete(ctx context.Context, id int) error {
	var user Users
	result := rep.Connect.Delete(&user, "id= ?", id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rep *MysqlUserRepository) Update(ctx context.Context, id int, email string, password string) error {
	result := rep.Connect.Where("id = ?", id).Updates(&Users{Email: email, Password: password})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
*/
