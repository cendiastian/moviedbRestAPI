package payments

import (
	"project/business/payments"
	"project/controllers"
	"project/controllers/payments/requests"
	"project/controllers/payments/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PaymentUC payments.Usecase
}

func NewPaymentController(paymentUsecase payments.Usecase) *PaymentController {
	return &PaymentController{
		PaymentUC: paymentUsecase,
	}
}

func (PaymentController PaymentController) Detail(c echo.Context) error {
	// fmt.Println("UserDetail")

	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	ctx := c.Request().Context()
	pay, err := PaymentController.PaymentUC.Detail(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(pay))
}

func (PaymentController PaymentController) GetAll(c echo.Context) error {

	ctx := c.Request().Context()
	pay, err := PaymentController.PaymentUC.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomain(pay))
}

func (PaymentController PaymentController) Delete(c echo.Context) error {

	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	ctx := c.Request().Context()
	_, err = PaymentController.PaymentUC.Delete(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Menghapus Data Payment Method")
}

func (PaymentController PaymentController) Update(c echo.Context) error {

	Update := requests.Update{}
	c.Bind(&Update)

	ctx := c.Request().Context()
	_, err := PaymentController.PaymentUC.Update(ctx, Update.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Merubah Data Payment Method")
}

func (PaymentController PaymentController) Register(c echo.Context) error {

	Register := requests.Payment_method{}
	c.Bind(&Register)

	ctx := c.Request().Context()
	user, err := PaymentController.PaymentUC.Register(ctx, Register.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}
