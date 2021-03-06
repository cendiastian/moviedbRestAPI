package subscription

import (
	"project/business/subscription"
	"project/controllers"
	"project/controllers/subscription/requests"
	"project/controllers/subscription/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SubcriptionController struct {
	SubcriptionUC subscription.Usecase
}

func NewSubcriptionController(subsUsecase subscription.Usecase) *SubcriptionController {
	return &SubcriptionController{
		SubcriptionUC: subsUsecase,
	}
}

func (SubcriptionController SubcriptionController) Detail(c echo.Context) error {
	// fmt.Println("UserDetail")

	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	ctx := c.Request().Context()
	subs, err := SubcriptionController.SubcriptionUC.Detail(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(subs))
}

func (SubcriptionController SubcriptionController) GetAll(c echo.Context) error {

	ctx := c.Request().Context()
	subs, err := SubcriptionController.SubcriptionUC.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.ToListDomain(subs))
}

func (SubcriptionController SubcriptionController) Delete(c echo.Context) error {

	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}
	ctx := c.Request().Context()
	err = SubcriptionController.SubcriptionUC.Delete(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Menghapus Data Subscription")
}

func (SubcriptionController SubcriptionController) Update(c echo.Context) error {

	Update := requests.Update{}
	c.Bind(&Update)
	ctx := c.Request().Context()

	err := SubcriptionController.SubcriptionUC.Update(ctx, Update.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Merubah Data Subscription")
}

func (SubcriptionController SubcriptionController) Createsubcription(c echo.Context) error {

	create := requests.SubcriptionPlan{}
	c.Bind(&create)
	ctx := c.Request().Context()
	user, err := SubcriptionController.SubcriptionUC.CreatePlan(ctx, create.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}
