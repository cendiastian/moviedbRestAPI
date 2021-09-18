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
	// eTest.POST("users/login", userLogin)
	// eTest.PUT("/users/:id", updateUser)
	// eTest.DELETE("/users/:id", deleteUser)
	return e
}
