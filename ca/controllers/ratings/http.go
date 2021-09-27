package ratings

import (
	"net/http"
	"project/ca/business/ratings"
	"project/ca/controllers"
	"project/ca/controllers/ratings/requests"
	"project/ca/controllers/ratings/responses"

	"github.com/labstack/echo/v4"
)

type RatingController struct {
	RateUC ratings.Usecase
}

func NewRatingController(RateUseCase ratings.Usecase) *RatingController {
	return &RatingController{
		RateUC: RateUseCase,
	}
}

func (RatingController RatingController) Delete(c echo.Context) error {

	RatingDelete := requests.RatingDelete{}
	c.Bind(&RatingDelete)

	ctx := c.Request().Context()
	err := RatingController.RateUC.Delete(ctx, RatingDelete.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Menghapus User")
}

func (RatingController RatingController) Update(c echo.Context) error {

	RatingUpdate := requests.RatingUpdate{}
	c.Bind(&RatingUpdate)

	ctx := c.Request().Context()
	err := RatingController.RateUC.Update(ctx, RatingUpdate.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Merubah Data User")
}

func (RatingController RatingController) Create(c echo.Context) error {

	RatingCreate := requests.RatingCreate{}
	c.Bind(&RatingCreate)

	ctx := c.Request().Context()
	user, err := RatingController.RateUC.Create(ctx, RatingCreate.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}
