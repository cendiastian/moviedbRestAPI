package users

import (
	"context"
	"fmt"
	"net/http"
	"project/ca/business/users"
	"project/ca/controllers"
	"project/ca/controllers/users/requests"
	"project/ca/controllers/users/responses"
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
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomainLogin(user))
}

func (UserController UserController) UserDetail(c echo.Context) error {
	// fmt.Println("UserDetail")

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

func (UserController UserController) UserRole(id int) string {
	role := ""
	user, err := UserController.UserUC.UserDetail(context.Background(), id)
	if err == nil {
		role = user.Name
	}
	return role
}
