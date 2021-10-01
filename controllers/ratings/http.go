package ratings

import (
	"net/http"
	"project/business/ratings"
	"project/controllers"
	"project/controllers/ratings/requests"
	"project/controllers/ratings/responses"

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

// func (RatingController RatingController) GetAllRate(c echo.Context) error {
// 	// fmt.Println("UserDetail")

// 	Movie, err := strconv.Atoi(c.Param("Movie"))
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	ctx := c.Request().Context()
// 	// trans, err := RatingController.RateUC.GetAllRate(ctx, Movie)
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controllers.NewSuccesResponse(c, responses.ToListDomain(trans))
// }

func (RatingController RatingController) Delete(c echo.Context) error {

	RatingDelete := requests.RatingDelete{}
	c.Bind(&RatingDelete)

	ctx := c.Request().Context()
	err := RatingController.RateUC.Delete(ctx, RatingDelete.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Menghapus Rating")
}

func (RatingController RatingController) Update(c echo.Context) error {

	RatingUpdate := requests.RatingUpdate{}
	c.Bind(&RatingUpdate)

	ctx := c.Request().Context()
	err := RatingController.RateUC.Update(ctx, RatingUpdate.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Merubah Data Rating")
}

func (RatingController RatingController) Create(c echo.Context) error {

	RatingCreate := requests.RatingCreate{}
	c.Bind(&RatingCreate)

	ctx := c.Request().Context()
	rate, err := RatingController.RateUC.Create(ctx, RatingCreate.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(rate))
}
func (RatingController RatingController) Detail(c echo.Context) error {
	// fmt.Println("UserDetail")

	RatingDelete := requests.RatingDelete{}
	c.Bind(&RatingDelete)

	ctx := c.Request().Context()
	trans, err := RatingController.RateUC.Detail(ctx, RatingDelete.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(trans))
}
