package movies

import (
	"fmt"
	"project/config"
	"project/model/movie"
	"project/model/response"

	"strings"

	// "project/routes"
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
	fmt.Println(query)
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

	for _, v := range createAllMovie.Search {
		res := config.DB.Find(&movieDB)

		movieDB.Title = v.Title
		movieDB.Year = v.Year
		movieDB.Poster = v.Poster
		movieDB.ImdbId = v.ImdbId
		movieDB.Type = v.Type
		movieDB.Actors = v.Actors
		movieDB.Director = v.Director
		movieDB.Plot = v.Plot
		movieDB.Writer = v.Writer
		// movieDB.Genre = createMovie.Genre
		if res != nil {
			movieDB.ID = (int(res.RowsAffected + 1))
			result := config.DB.Create(&movieDB)
			if result.Error != nil {
				return c.JSON(http.StatusInternalServerError, response.Response{
					Code:    http.StatusInternalServerError,
					Message: "Error ketika input data category ke config.DB",
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
	fmt.Println(query)
	req.URL.RawQuery = query.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &GetDetail)
	//membuat category jika belum ada
	var categoryDB movie.Category
	Genre := strings.Split(GetDetail.Genre, ",")
	for _, v := range Genre {
		result := config.DB.Find(&categoryDB)
		categoryDB.Name = string(v)
		if result != nil {
			res := config.DB.Where("Name = ?", v).First(&categoryDB)
			if res != nil {
				continue
			}
			categoryDB.ID = int(result.RowsAffected + 1)
			config.DB.Create(&categoryDB)
		} else {
			config.DB.Create(&categoryDB)
		}
	}
	// update
	result := config.DB.Model(&movie.Movie{}).Where("Title = ?", GetDetail.Title).Updates(movie.Movie{Director: GetDetail.Director, Writer: GetDetail.Writer, Actors: GetDetail.Actors, Plot: GetDetail.Plot})
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

// category := strings.Split(createMovie.Genre, ",")
// var categoryName []*gorm.DB
// type Result struct {
// 	Name string
//   }
// var ress Result
// for _, v := range category {
// 	result := config.DB.Find(&categoryDB)
// 	categoryDB.Name = v
// 	config.DB.Table("category").Select("name").Scan(&ress)
// 	if result != nil {
// 		ress := config.DB.Where("name = ?", v).First(&categoryDB)
// 		if ress.Name == v{
// 			continue
// 		}
// 		categoryDB.ID = int(result.RowsAffected)
// 		config.DB.Create(&categoryDB)
// 	} else {
// 		// fmt.Println(categoryDB)
// 		config.DB.Create(&categoryDB)
// 	}
// }
