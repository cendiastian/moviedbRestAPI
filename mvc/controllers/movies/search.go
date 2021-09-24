package movies

// import (
// 	// "fmt"
// 	// "project/mvc/config"

// 	"project/mvc/config"
// 	"project/mvc/model/movie"

// 	// "strings"

// 	// "project/routes"
// 	// "strconv"
// 	// "log"

// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func DecodeGRPCSearchRequest(c echo.Context, r interface{}) (interface{}, error) {
// 	req := r.(*movie.SearchRequest)
// 	return movie.SearchRequest{
// 		MovieTitle: req.MovieTitle,
// 	}, nil
// }
// func EncodeGRPCSearchResponse(c echo.Context, r interface{}) error {
// 	// var movieDB movie.Movie
// 	title := c.Param("title")
// 	resp := r.(*movie.SearchResponse)

// 	search := config.DB.Where("Title LIKE ?", "%title%").Find(&resp.Response)
// 	result := make([]*movie.Movie, 0)
// 	for _, v := range resp.Response {
// 		result = append(result, &movie.Movie{
// 			Title:  v.Title,
// 			Year:   v.Year,
// 			Type:   v.Type,
// 			ImdbId: v.ImdbId,
// 			Poster: v.Poster,
// 		})
// 	}

// 	return c.JSON(http.StatusOK, movie.SearchResponse{
// 		Code:     http.StatusOK,
// 		Message:  "Berhasil mencari Movie",
// 		Response: []movie.Movie{},
// 	})

// }
