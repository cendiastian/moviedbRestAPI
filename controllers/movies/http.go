package movies

import (
	"fmt"
	"project/business/movies"
	"project/controllers"
	"project/controllers/movies/requests"
	"project/controllers/movies/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MovieController struct {
	MovieUC movies.Usecase
}

func NewMovieController(MovieUsecase movies.Usecase) *MovieController {
	return &MovieController{
		MovieUC: MovieUsecase,
	}
}

func (MovieController MovieController) CreateMovie(c echo.Context) error {
	ImdbId := (c.Param("ImdbId"))
	ctx := c.Request().Context()
	fmt.Println(ImdbId)
	movie, err := MovieController.MovieUC.CreateMovie(ctx, ImdbId)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainMovie(movie))
}

func (MovieController MovieController) MovieDetail(c echo.Context) error {
	// fmt.Println("MovieDetail")

	Id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}
	user, err := strconv.Atoi(c.QueryParam("user"))
	fmt.Println(Id)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	ctx := c.Request().Context()
	movie, err := MovieController.MovieUC.MovieDetail(ctx, Id, user)
	// fmt.Println(movie)
	if err != nil {
		fmt.Println("error")
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainMovie(movie))
}

func (MovieController MovieController) SearchMovie(c echo.Context) error {

	title := c.QueryParam("title")

	ctx := c.Request().Context()
	movie, err := MovieController.MovieUC.SearchMovie(ctx, title)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomainSearch(movie))
}

func (MovieController MovieController) FilterOrder(c echo.Context) error {

	Order := c.QueryParam("Order")

	fmt.Println(Order)
	ctx := c.Request().Context()
	movie, err := MovieController.MovieUC.FilterOrder(ctx, Order)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomainSearch(movie))
}

func (MovieController MovieController) FilterGenre(c echo.Context) error {
	// var FilterGenre []movies.Movie
	Genre := c.QueryParam("Genre")
	fmt.Println(Genre)
	ctx := c.Request().Context()
	movie, err := MovieController.MovieUC.FilterGenre(ctx, Genre)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomainSearch(movie))
}

func (MovieController MovieController) DeleteAll(c echo.Context) error {

	ctx := c.Request().Context()
	err := MovieController.MovieUC.DeleteAll(ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Menghapus semua Data Movie")
}

func (MovieController MovieController) DeleteMovie(c echo.Context) error {

	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	ctx := c.Request().Context()
	err = MovieController.MovieUC.DeleteMovie(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Menghapus Data Movie")
}

func (MovieController MovieController) GetAllMovie(c echo.Context) error {

	ctx := c.Request().Context()
	movie, err := MovieController.MovieUC.GetAllMovie(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomainSearch(movie))
}

func (MovieController MovieController) UpdateMovie(c echo.Context) error {

	Update := requests.MovieUpdate{}
	c.Bind(&Update)

	ctx := c.Request().Context()
	err := MovieController.MovieUC.UpdateMovie(ctx, Update.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Merubah Data Movie")
}

/*Movie
	func (userController UserController) Login(c echo.Context) error {
	fmt.Println("Login")

	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()
	// user, err := userController.UserUC.Login(ctx, userLogin.Email, userLogin.Password)
	user, err := userController.UserUC.Login(ctx, userLogin.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}

func (UserController UserController) MovieDetail(c echo.Context) error {
	// fmt.Println("MovieDetail")

	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	ctx := c.Request().Context()
	user, err := UserController.UserUC.UserDetail(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}

func (UserController UserController) GetAll(c echo.Context) error {

	ctx := c.Request().Context()
	user, err := UserController.UserUC.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomain(user))
}

func (UserController UserController) Delete(c echo.Context) error {

	userDelete := requests.UserDelete{}
	c.Bind(&userDelete)

	ctx := c.Request().Context()
	err := UserController.UserUC.Delete(ctx, userDelete.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Menghapus User")
}

func (UserController UserController) Update(c echo.Context) error {

	userUpdate := requests.UserUpdate{}
	c.Bind(&userUpdate)

	ctx := c.Request().Context()
	err := UserController.UserUC.Update(ctx, userUpdate.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Merubah Data User")
}

func (UserController UserController) Register(c echo.Context) error {

	UserRegister := requests.UserRegister{}
	c.Bind(&UserRegister)

	ctx := c.Request().Context()
	user, err := UserController.UserUC.Register(ctx, UserRegister.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}
*/
