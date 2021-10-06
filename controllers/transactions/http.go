package transactions

import (
	"net/http"
	"project/business/transactions"
	"project/controllers"
	"project/controllers/transactions/requests"
	"project/controllers/transactions/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransController struct {
	TransUC transactions.Usecase
}

func NewTransController(TransUsecase transactions.Usecase) *TransController {
	return &TransController{
		TransUC: TransUsecase,
	}
}

func (TransController TransController) CreateTransaction(c echo.Context) error {

	CreateTransaction := requests.Transaction{}
	c.Bind(&CreateTransaction)

	ctx := c.Request().Context()
	trans, err := TransController.TransUC.CreateTransaction(ctx, CreateTransaction.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.CreateTransaction(trans))
}
func (TransController TransController) DetailTrans(c echo.Context) error {
	// fmt.Println("UserDetail")

	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	ctx := c.Request().Context()
	trans, err := TransController.TransUC.DetailTrans(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainTransaction(trans))
}
