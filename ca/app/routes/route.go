package routes

import (
	"project/ca/controllers/movies"
	"project/ca/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig       middleware.JWTConfig
	UserController  users.UserController
	MovieController movies.MovieController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("users/login", cl.UserController.Login)
	e.GET("users", cl.UserController.Login, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("users/:Id", cl.UserController.UserDetail)
	e.GET("users", cl.UserController.GetAll)
	e.DELETE("users/:id", cl.UserController.Delete)
	e.PUT("users/:id", cl.UserController.Update)
	e.POST("users/register", cl.UserController.Register)

	e.POST("movies/CreateMovie/:ImdbId", cl.MovieController.CreateMovieAPI)
	e.GET("movies/MovieDetail/:Id", cl.MovieController.MovieDetail)
	e.GET("movies/SearchMovie/:Title", cl.MovieController.SearchMovie)
}
