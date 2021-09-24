package movies

import (
	"fmt"
	"project/mvc/config"
	"project/mvc/model/movie"
	"project/mvc/model/response"

	"strings"

	// "strconv"
	// "log"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateMovie(c echo.Context) error {
	searchTerm := (c.Param("searchTerm"))
	var createAllMovie movie.GetAllMovie

	req, _ := http.NewRequest("GET", "http://www.omdbapi.com/?apikey=8b8a25e8&", nil)

	query := req.URL.Query()
	query.Add("s", searchTerm)
	req.URL.RawQuery = query.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &createAllMovie)

	var movieDB movie.Movie
	var Allmovie []movie.Movie
	for _, v := range createAllMovie.Search {
		res := config.DB.Find(&Allmovie)
		movieDB = movie.Movie{
			Title:    v.Title,
			Year:     v.Year,
			Poster:   v.Poster,
			ImdbId:   v.ImdbId,
			Type:     v.Type,
			Actors:   v.Actors,
			Director: v.Director,
			Plot:     v.Plot,
			Writer:   v.Writer,
			// Genre: []movie.Category,
		}
		// movieDB.Genre = createMovie.Genre
		// count := config.DB.Last(&movieDB)
		if res != nil {
			sameTitle := config.DB.Raw("SELECT * FROM `movies` WHERE Title = ? LIMIT 1 ", movieDB.Title).Scan(&movieDB)
			if sameTitle.RowsAffected == 0 {
				movieDB.ID = int(res.RowsAffected + 1)
				result := config.DB.Create(&movieDB)
				if result.Error != nil {
					return c.JSON(http.StatusInternalServerError, response.Response{
						Code:    http.StatusInternalServerError,
						Message: "Error ketika input data category ke config.DB",
						Data:    nil,
					})
				}
			} else {
				return c.JSON(http.StatusBadRequest, response.Response{
					Code:    http.StatusBadRequest,
					Message: "Movie sudah ada",
					Data:    nil,
				})
			}
		} else {
			result := config.DB.Create(&movieDB)
			if result.Error != nil {
				return c.JSON(http.StatusInternalServerError, response.Response{
					Code:    http.StatusInternalServerError,
					Message: "Error ketika input data category ke config.DB",
					Data:    nil,
				})
			}
		}
	}

	return c.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Message: "Berhasil mendapat Movie",
		Data:    createAllMovie,
	})
}

func GetDetailFromAPI(c echo.Context) error {
	title := (c.Param("title"))

	var GetDetail movie.GetDetail

	req, _ := http.NewRequest("GET", "http://www.omdbapi.com/?apikey=8b8a25e8&", nil)

	query := req.URL.Query()
	query.Add("t", title)
	req.URL.RawQuery = query.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &GetDetail)

	var movieDB movie.Movie
	var categoryDB movie.Category
	// var category []movie.Categor
	GenreName := strings.Split(GetDetail.Genre, ", ")
	fmt.Println(GenreName)
	for _, v := range GenreName {
		fmt.Println(v)
		// res := config.DB.FirstOrCreate(&categoryDB, movie.Category{Name: v})
		// res := config.DB.Where("Name = ?", v).First(&categoryDB)
		res := config.DB.Raw("SELECT * FROM `categories`WHERE Name = ? LIMIT 1 ", v).Scan(&categoryDB)
		fmt.Println(res)
		// if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		if res.RowsAffected == 0 {
			config.DB.Raw("SELECT id FROM `categories` ORDER BY `categories`.`id` DESC").Scan(&categoryDB)
			categoryDB = movie.Category{
				Name: v,
				ID:   +1,
			}
			// categoryDB.Name = v
			// categoryDB.ID += 1
			fmt.Println(categoryDB)
			config.DB.Create(categoryDB)
			movieDB.Genre = append(movieDB.Genre, categoryDB)
		} else {
			fmt.Println(categoryDB)
			movieDB.Genre = append(movieDB.Genre, categoryDB)
		}
		// if res != nil {
		// 	fmt.Println(categoryDB)
		// 	category = append(category, categoryDB)
		// } else {
		// 	config.DB.Create(Genre)
		// 	fmt.Println(Genre)
		// 	category = append(category, Genre)
		// }
		fmt.Println(movieDB.Genre)
	}

	// update
	result := config.DB.Model(&movie.Movie{}).Where("Title = ?", GetDetail.Title).Updates(movie.Movie{Director: GetDetail.Director, Writer: GetDetail.Writer, Actors: GetDetail.Actors, Plot: GetDetail.Plot, Genre: movieDB.Genre})
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika input data category ke config.DB",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Message: "Berhasil mendapat detail",
		Data:    GetDetail,
	})
}
