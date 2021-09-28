package routes

import (
	"project/ca/controllers/genres"
	"project/ca/controllers/movies"
	"project/ca/controllers/payments"
	"project/ca/controllers/ratings"
	"project/ca/controllers/subscription"
	"project/ca/controllers/transactions"
	"project/ca/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig             middleware.JWTConfig
	UserController        users.UserController
	MovieController       movies.MovieController
	SubcriptionController subscription.SubcriptionController
	PaymentController     payments.PaymentController
	RatingController      ratings.RatingController
	GenreController       genres.GenreController
	TransController       transactions.TransController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	// USER
	e.POST("users/login", cl.UserController.Login)
	e.GET("users", cl.UserController.Login, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("users/:Id", cl.UserController.UserDetail)
	e.GET("users", cl.UserController.GetAll)
	e.DELETE("users/:id", cl.UserController.Delete)
	e.PUT("users/:id", cl.UserController.Update, middleware.JWTWithConfig(cl.JwtConfig))
	e.POST("users/register", cl.UserController.Register)

	// MOVIE_API
	e.POST("movies/CreateMovie/:ImdbId", cl.MovieController.CreateMovie)

	// MOVIE
	e.GET("movies/MovieDetail/:Id", cl.MovieController.MovieDetail)
	e.GET("movies", cl.MovieController.GetAllMovie)
	e.PUT("movies/update", cl.MovieController.UpdateMovie)
	e.DELETE("movies/delete/:Id", cl.MovieController.DeleteMovie)
	e.DELETE("movies/deleteAll", cl.MovieController.DeleteAll)
	e.GET("movies/search", cl.MovieController.SearchMovie)
	e.GET("movies/genre", cl.MovieController.FilterGenre)
	e.GET("movies/order", cl.MovieController.FilterOrder)

	// SUBSCRIPTION
	e.POST("subs/Create", cl.SubcriptionController.Createsubcription)
	e.GET("subs/Detail/:Id", cl.SubcriptionController.Detail)
	e.GET("subs/all", cl.SubcriptionController.GetAll)
	e.PUT("subs/update", cl.SubcriptionController.Update)
	e.DELETE("subs/delete/:Id", cl.SubcriptionController.Delete)

	// PAYMENT
	e.POST("pay/Create", cl.PaymentController.Register)
	e.GET("pay/Detail/:Id", cl.PaymentController.Detail)
	e.GET("pay/all", cl.PaymentController.GetAll)
	e.PUT("pay/update", cl.PaymentController.Update)
	e.DELETE("pay/delete/:Id", cl.PaymentController.Delete)

	// TRANSACTIONS
	e.POST("transaction/Create", cl.TransController.CreateTransaction, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("transaction/Detail/:Id", cl.TransController.DetailTrans)

	// RATING
	e.POST("rate/Create", cl.RatingController.Create)
	e.GET("rate/Detail/:Movie/:User", cl.RatingController.Detail, middleware.JWTWithConfig(cl.JwtConfig))
	e.PUT("rate/update", cl.RatingController.Update)
	e.DELETE("rate/delete", cl.RatingController.Delete)
}

// func RoleValidation(role string, userControler users.UserController) echo.MiddlewareFunc {
// 	return func(hf echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			claims := middlewareApp.GetUser(c)

// 			// userRole := userControler.UserRole(claims.ID)

// 			if userRole == role {
// 				return hf(c)
// 			} else {
// 				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("forbidden roles"))
// 			}
// 		}
// 	}
// }
