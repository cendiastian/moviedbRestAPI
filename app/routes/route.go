package routes

import (
	"project/app/middlewares"
	"project/controllers/genres"
	"project/controllers/movies"
	"project/controllers/payments"
	"project/controllers/ratings"
	"project/controllers/subscription"
	"project/controllers/transactions"
	"project/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware      middlewares.MongoConfig
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

	cl.LoggerMiddleware.Start(e)
	// USER
	e.POST("users/login", cl.UserController.Login)
	// e.GET("users", cl.UserController.Login)
	e.GET("users/:Id", cl.UserController.UserDetail)
	e.GET("users", cl.UserController.GetAll)
	e.DELETE("users/:id", cl.UserController.Delete)
	e.PUT("users/:id", cl.UserController.Update, middleware.JWTWithConfig(cl.JwtConfig))
	e.POST("users/register", cl.UserController.Register)

	// MOVIE_API
	e.POST("movies/CreateMovie/:ImdbId", cl.MovieController.CreateMovie)

	// MOVIE
	e.GET("movies/detail", cl.MovieController.MovieDetail, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("movies", cl.MovieController.GetAllMovie)
	e.PUT("movies/update", cl.MovieController.UpdateMovie)
	e.DELETE("movies/delete/:Id", cl.MovieController.DeleteMovie)
	e.DELETE("movies/deleteAll", cl.MovieController.DeleteAll)
	e.GET("movies/search", cl.MovieController.SearchMovie)
	e.GET("movies/genre", cl.MovieController.FilterGenre)
	e.GET("movies/order", cl.MovieController.FilterOrder)

	// GENRE
	e.GET("genre/all", cl.GenreController.GetAllGenre)

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
	e.POST("rate/Create", cl.RatingController.Create, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("rate/detail", cl.RatingController.Detail)
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
