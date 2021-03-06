package users

import (
	"fmt"
	"net/http"
	"project/business/users"
	"project/controllers"
	"project/controllers/users/requests"
	"project/controllers/users/responses"
	"strconv"

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
	fmt.Println("Login")

	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()
	// user, err := userController.UserUC.Login(ctx, userLogin.Email, userLogin.Password)
	user, err := userController.UserUC.Login(ctx, userLogin.ToDomain())
	fmt.Println(user.Token)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomainLogin(user))
}

func (UserController UserController) UserDetail(c echo.Context) error {
	// fmt.Println("UserDetail")

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

	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}
	ctx := c.Request().Context()
	_, err = UserController.UserUC.Delete(ctx, Id)
	if err != nil {
		return controllers.NewErrorResponse(c, controllers.ErrorCode(err), err)
	}

	return controllers.UpdateSuccesResponse(c, "Berhasil Menghapus User")
}

func (UserController UserController) Update(c echo.Context) error {

	userUpdate := requests.UserUpdate{}
	c.Bind(&userUpdate)

	ctx := c.Request().Context()
	_, err := UserController.UserUC.Update(ctx, userUpdate.ToDomain())
	if err != nil {
		// if err ==
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
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}
