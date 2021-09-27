package movies

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/ca/business/movies"

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

func (rep *MysqlMovieRepository) GetAPI(ctx context.Context, ImdbId string) (movies.GetAPI, error) {
	var movie GetMovieAPI

	req, err := http.NewRequest("GET", "http://www.omdbapi.com/?apikey=8b8a25e8&", nil)
	if err != nil {
		return movies.GetAPI{}, err
	}

	query := req.URL.Query()
	query.Add("i", ImdbId)
	req.URL.RawQuery = query.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return movies.GetAPI{}, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &movie)

	return movie.ToDomainAPI(), nil
}

func (rep *MysqlMovieRepository) CreateMovieAPI(ctx context.Context, Title string, Year string, ImdbId string, Type string, Poster string, Genre []movies.Genre, Writer string, Actors string) (movies.Movie, error) {
	movie := Movies{
		Title:  Title,
		Year:   Year,
		ImdbId: ImdbId,
		Type:   Type,
		Poster: Poster,
		Genre:  Genre,
		Writer: Writer,
		Actors: Actors,
	}

	result := rep.Connect.Create(&movie)
	if result.Error != nil {
		return movies.Movie{}, result.Error
	}

	return movie.ToDomainMovie(), nil
}

func (rep *MysqlMovieRepository) MovieDetail(ctx context.Context, id int) (movies.Movie, error) {
	var movie Movies

	result := rep.Connect.First(&movie, "id= ?", id)
	if result.Error != nil {
		return movies.Movie{}, result.Error
	}
	res := rep.Connect.Preload("Genre").Find(&movie)
	if res.Error != nil {
		return movies.Movie{}, res.Error
	}
	return movie.ToDomainMovie(), nil
}

func (rep *MysqlMovieRepository) ScanGenre(ctx context.Context, name string) (movies.Genre, error) {
	var genre Genres
	genre.Name = name
	result := rep.Connect.FirstOrCreate(&genre, "Name= ?", name)

	if result.Error != nil {
		return movies.Genre{}, result.Error
	}

	return genre.ToDomainGenre(), nil
}

func (rep *MysqlMovieRepository) SearchMovie(ctx context.Context, title string) ([]movies.Movie, error) {
	var movie []Movies

	fmt.Println(title)

	result := rep.Connect.Where("title LIKE ?", title+"%").Find(&movie)
	if result.Error != nil {
		return []movies.Movie{}, result.Error
	}

	result = rep.Connect.Where("title LIKE ?", title+"%").Find(&movie)
	if result.Error != nil {
		return []movies.Movie{}, result.Error
	}

	result = rep.Connect.Where("title LIKE ?", "%"+title+"%").Find(&movie)
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
		Joins("JOIN movie_genres on movie_genres.movies_id = movies.id JOIN genres on movie_genres.genre_id = genres.id AND genres.name = ? ",
			genre).Find(&movie)
	if result.Error != nil {
		return []movies.Movie{}, result.Error
	}
	// result = rep.Connect.Find(&movie)
	// if result.Error != nil {
	// 	return []movies.Movie{}, result.Error
	// }

	// for i, v := range movie {
	// 	for _, n := range v.Genre {
	// 		if n.Name == Genre {
	// 		FilterGenre = append(FilterGenre, []movies.Movie{})
	// 		}
	// 	}
	// }

	// result = rep.Connect.Not("Genre = ?", ]).Find(&movie)
	// if result.Error != nil {
	// 	return []movies.Movie{}, result.Error
	// }
	// // result := rep.Connect.Joins("Genre").First(&movie, "genre.name = ?", genre)
	// if result.Error != nil {
	// 	return []movies.Movie{}, result.Error
	// }
	// q := rep.Connect.Where("Genre = ?", Genre).Find(&movie)
	// // q := rep.Connect.Where("Genre__name", genre).Find(&movie)
	// if q.Error != nil {
	// 	return []movies.Movie{}, q.Error
	// }
	// fmt.Println(genre)
	// fmt.Println(result)
	// if result.Error != nil {
	// 	return []movies.Movie{}, result.Error
	// }

	return ToListDomain(movie), nil
}
func (rep *MysqlMovieRepository) FilterOrder(ctx context.Context, order string) ([]movies.Movie, error) {
	fmt.Println("Tess" + order)
	var movie []Movies
	// res := rep.Connect.Find(&movie)
	// if res.Error != nil {
	// 	return []movies.Movie{}, res.Error
	// }
	if order == "Terlama" {
		result := rep.Connect.Preload("Genre").Find(&movie)
		if result.Error != nil {
			return []movies.Movie{}, result.Error
		}
	} else {
		result := rep.Connect.Preload("Genre").Order("id desc").Find(&movie)
		if result.Error != nil {
			return []movies.Movie{}, result.Error
		}
	}
	return ToListDomain(movie), nil
}

func (rep *MysqlMovieRepository) GetAllMovie(ctx context.Context) ([]movies.Movie, error) {
	var movie []Movies
	result := rep.Connect.Preload("Genre").Find(&movie)
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

func (rep *MysqlMovieRepository) UpdateMovie(ctx context.Context, id int, Title string, Type string) error {
	result := rep.Connect.Where("id = ?", id).Updates(&Movies{Title: Title, Type: Type})

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
