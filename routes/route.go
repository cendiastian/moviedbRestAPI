package routes

import (
	"project/controllers/movies"
	"project/controllers/users"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {

	e := echo.New()
	eTest := e.Group("test/")
	eTest.GET("users", users.GetUser)
	eTest.POST("users/register", users.UserRegist)
	eTest.GET("users/:userId", users.UserDetail)
	eTest.POST("users/login", users.UserLogin)
	eTest.PUT("users/update", users.UpdateUser)
	eTest.DELETE("users/delete", users.DeleteUser)

	eTest1 := e.Group("movies/")
	eTest1.POST("create/:searchTerm", movies.CreateMovie)
	eTest1.PUT("detailAPI/:title", movies.GetDetailFromAPI)
	return e
}
