package users

import (
	"net/http"
	"project/ca/business/users"
	"project/ca/controllers"
	"project/ca/controllers/users/requests"
	"project/ca/controllers/users/responses"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUC users.Usecase
}

func NewUserController(UserUsecase users.Usecase) *UserController {
	return &UserController{
		UserUC: UserUsecase,
	}
}

func (userController UserController) Login(c echo.Context) error {
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
