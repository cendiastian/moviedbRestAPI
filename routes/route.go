package routes

import (
	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {

	e := echo.New()
	eTest := e.Group("test/")
	eTest.GET("users", getUser)
	eTest.POST("users/register", userRegist)
	eTest.GET("users/:userId", userDetail)
	eTest.POST("users/login", userLogin)
	eTest.PUT("users/update", updateUser)
	eTest.DELETE("users/delete", deleteUser)
	// eTest1 := e.Group("movies/")
	// eTest1.GET("movies", getMovie)
	return e
}
