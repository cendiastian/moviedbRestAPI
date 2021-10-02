package genres

import (
	"net/http"
	"project/business/genres"
	"project/controllers"
	"project/controllers/genres/responses"

	"github.com/labstack/echo/v4"
)

type GenreController struct {
	GenreUC genres.Usecase
}

func NewGenreController(GenreUsecase genres.Usecase) *GenreController {
	return &GenreController{
		GenreUC: GenreUsecase,
	}
}

func (GenreController GenreController) GetAllGenre(c echo.Context) error {

	ctx := c.Request().Context()
	Genre, err := GenreController.GenreUC.GetAllGenre(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomain(Genre))
}
