package movies

import (
	"fmt"
	"net/http"
	"project/ca/business/movies"
	"project/ca/controllers"
	"project/ca/controllers/movies/requests"
	"project/ca/controllers/movies/responses"
	"strconv"
	"strings"

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

func (MovieController MovieController) CreateMovieAPI(c echo.Context) error {
	ImdbId := (c.Param("ImdbId"))
	CreateMovieAPI := requests.CreateMovieAPI{}
	CreateGenreAPI := requests.CreateGenreAPI{}

	ctx := c.Request().Context()
	get, err := MovieController.MovieUC.GetAPI(ctx, ImdbId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	GenreName := strings.Split(get.Genre, ", ")

	for _, v := range GenreName {
		CreateGenreAPI.Name = v
		fmt.Println(CreateGenreAPI.Name)
		scan, err := MovieController.MovieUC.ScanGenre(ctx, CreateGenreAPI.ToDomainGenre())
		if err != nil {
			return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
		}
		CreateMovieAPI.Genre = append(CreateMovieAPI.Genre, scan)
		if err != nil {
			return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
		}
	}

	CreateMovieAPI = requests.CreateMovieAPI{
		Title:  get.Title,
		ImdbId: get.ImdbId,
		Year:   get.Year,
		Type:   get.Type,
		Poster: get.Poster,
		Genre:  CreateMovieAPI.Genre,
		Writer: get.Writer,
		Actors: get.Actors,
	}
	fmt.Println(CreateMovieAPI.Year)

	movie, err := MovieController.MovieUC.CreateMovieAPI(ctx, CreateMovieAPI.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainMovie(movie))
}

/* func (MovieController MovieController) CreateGenreAPI(c echo.Context) error {

// 	CreateMovieAPI := requests.CreateMovieAPI{}
// 	CreateGenreAPI := requests.CreateGenreAPI{}
// 	c.Bind(&CreateMovieAPI)

// 	ctx := c.Request().Context()
// 	movie, err := MovieController.MovieUC.CreateMovieAPI(ctx, CreateMovieAPI.ToDomain())
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}
// 	genre, res := MovieController.MovieUC.CreateMovieAPI(ctx, CreateGenreAPI.ToDomainGenre())
// 	if res != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controllers.NewSuccesResponse(c, responses.FromDomainGenre(genre))
 } */

func (MovieController MovieController) MovieDetail(c echo.Context) error {
	// fmt.Println("MovieDetail")

	Id, err := strconv.Atoi(c.Param("Id"))
	fmt.Println(Id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	movie, err := MovieController.MovieUC.MovieDetail(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainMovie(movie))
}

func (MovieController MovieController) SearchMovie(c echo.Context) error {

	title := c.QueryParam("title")

	ctx := c.Request().Context()
	movie, err := MovieController.MovieUC.SearchMovie(ctx, title)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomain(movie))
}

func (MovieController MovieController) FilterOrder(c echo.Context) error {

	Order := c.QueryParam("Order")

	fmt.Println(Order)
	ctx := c.Request().Context()
	movie, err := MovieController.MovieUC.FilterOrder(ctx, Order)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomain(movie))
}

func (MovieController MovieController) FilterGenre(c echo.Context) error {
	// var FilterGenre []movies.Movie
	Genre := c.QueryParam("Genre")
	fmt.Println(Genre)
	ctx := c.Request().Context()
	movie, err := MovieController.MovieUC.FilterGenre(ctx, Genre)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	// for i, v := range movie {
	// 	for _, b := range v.Genre {
	// 		if b.Name == Genre {
	// 			FilterGenre = append(FilterGenre, movie[i])
	// 		}
	// 	}
	// }

	return controllers.NewSuccesResponse(c, responses.ToListDomain(movie))
}

func (MovieController MovieController) DeleteAll(c echo.Context) error {

	ctx := c.Request().Context()
	error := MovieController.MovieUC.DeleteAll(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainMovie(movies.Movie{}))
}

func (MovieController MovieController) DeleteMovie(c echo.Context) error {

	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	movie, err := MovieController.MovieUC.DeleteMovie(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainMovie(movie))
}

func (MovieController MovieController) GetAllMovie(c echo.Context) error {

	ctx := c.Request().Context()
	movie, err := MovieController.MovieUC.GetAllMovie(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomain(movie))
}

func (MovieController MovieController) UpdateMovie(c echo.Context) error {

	Update := requests.MovieUpdate{}
	c.Bind(&Update)

	ctx := c.Request().Context()
	err := MovieController.MovieUC.UpdateMovie(ctx, Update.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Merubah Data User")
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
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}

func (UserController UserController) MovieDetail(c echo.Context) error {
	// fmt.Println("MovieDetail")

	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	user, err := UserController.UserUC.UserDetail(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}

func (UserController UserController) GetAll(c echo.Context) error {

	ctx := c.Request().Context()
	user, err := UserController.UserUC.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomain(user))
}

func (UserController UserController) Delete(c echo.Context) error {

	userDelete := requests.UserDelete{}
	c.Bind(&userDelete)

	ctx := c.Request().Context()
	err := UserController.UserUC.Delete(ctx, userDelete.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Menghapus User")
}

func (UserController UserController) Update(c echo.Context) error {

	userUpdate := requests.UserUpdate{}
	c.Bind(&userUpdate)

	ctx := c.Request().Context()
	err := UserController.UserUC.Update(ctx, userUpdate.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Merubah Data User")
}

func (UserController UserController) Register(c echo.Context) error {

	UserRegister := requests.UserRegister{}
	c.Bind(&UserRegister)

	ctx := c.Request().Context()
	user, err := UserController.UserUC.Register(ctx, UserRegister.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}
*/
