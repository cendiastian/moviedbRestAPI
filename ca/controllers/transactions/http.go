package transactions

import (
	"net/http"
	"project/ca/business/transactions"
	"project/ca/controllers"
	"project/ca/controllers/transactions/requests"
	"project/ca/controllers/transactions/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PaymentUC transactions.Usecase
}

func NewPaymentController(paymentUsecase transactions.Usecase) *PaymentController {
	return &PaymentController{
		PaymentUC: paymentUsecase,
	}
}

func (PaymentController PaymentController) Detail(c echo.Context) error {
	// fmt.Println("UserDetail")

	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	pay, err := PaymentController.PaymentUC.Detail(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(pay))
}

func (PaymentController PaymentController) GetAll(c echo.Context) error {

	ctx := c.Request().Context()
	pay, err := PaymentController.PaymentUC.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomain(pay))
}

func (PaymentController PaymentController) Delete(c echo.Context) error {

	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	pay, err := PaymentController.PaymentUC.Delete(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(pay))
}

func (PaymentController PaymentController) Update(c echo.Context) error {

	Update := requests.Update{}
	c.Bind(&Update)

	ctx := c.Request().Context()
	err := PaymentController.PaymentUC.Update(ctx, Update.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Merubah Data User")
}

func (PaymentController PaymentController) Register(c echo.Context) error {

	Register := requests.Payment_method{}
	c.Bind(&Register)

	ctx := c.Request().Context()
	user, err := PaymentController.PaymentUC.Register(ctx, Register.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}
func (PaymentController PaymentController) CreateTransaction(c echo.Context) error {

	CreateTransaction := requests.Transaction{}
	c.Bind(&CreateTransaction)

	ctx := c.Request().Context()
	user, err := PaymentController.PaymentUC.CreateTransaction(ctx, CreateTransaction.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainTransaction(user))
}
