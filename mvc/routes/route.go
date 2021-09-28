package routes

import (
	"project/mvc/constants"
	"project/mvc/controllers/movies"
	"project/mvc/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoute() *echo.Echo {

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	jwt := middleware.JWT([]byte(constants.SECRET_JWT_USER))

	eTest := e.Group("test/")
	eTest.GET("users", users.GetUser, jwt)
	eTest.POST("users/register", users.UserRegist)
	eTest.GET("users/:userId", users.UserDetail, jwt)
	eTest.POST("users/login", users.UserLogin)
	eTest.PUT("users/update", users.UpdateUser)
	eTest.DELETE("users/delete", users.DeleteUser)

	eTest1 := e.Group("movies/")
	eTest1.POST("create/:searchTerm", movies.CreateMovie)
	eTest1.PUT("detailAPI/:title", movies.GetDetailFromAPI)

	return e
}
