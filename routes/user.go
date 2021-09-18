package routes

import (
	"net/http"
	"project/config"
	"project/model/response"
	"project/model/user"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func userRegist(c echo.Context) error {
	var userRegister user.UserRegister
	c.Bind(&userRegister)

	// validasi
	if userRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "Mohon isi Nama",
			Data:    nil,
		})
	} else if userRegister.Password == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "Mohon isi Password",
			Data:    nil,
		})
	} else if userRegister.Email == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "Mohon isi Email",
			Data:    nil,
		})
	}
	var userDB user.User
	userDB.Name = userRegister.Name
	userDB.Password = userRegister.Password
	userDB.Email = userRegister.Email

	result := config.DB.Create(&userDB)
	if result != nil {
		return c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Message: "Berhasil register",
			Data:    userDB,
		})
	}
	return c.JSON(http.StatusInternalServerError, response.Response{
		Code:    http.StatusInternalServerError,
		Message: "Error ketika input data user ke config.DB",
		Data:    nil,
	})
}

func getUser(c echo.Context) error {
	users := []user.User{}
	result := config.DB.Find(&users)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.Response{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data user dari config.DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data user",
		Data:    users,
	})
}

func userDetail(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Gagal konversi userId",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    user.User{ID: userId},
	})
}

// type user.UserLogin struct {
// 	Email    string `json:"email"`
// 	Password int    `json:"password"`
// }

// type UserUpdate struct {
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// func updateUser(c echo.Context) error {
// 	// var userRegister UserRegister
// 	var user user.User
// 	u := new(UserUpdate)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	// id, _ := strconv.Atoi(c.Param("id"))
// 	user.Name = u.Name
// 	return c.JSON(http.StatusOK, user.User)
// }

// func deleteUser(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	delete(user.User, id)
// 	return c.NoContent(http.StatusNoContent)
// }

// func userLogin(c echo.Context) error {
// 	userLogin := user.UserLogin{}
// 	c.Bind(&userLogin)

// 	return c.JSON(http.StatusOK, response.Response{
// 		Code:    http.StatusOK,
// 		Message: "Berhasil",
// 		Data:    userLogin,
// 	})
// }
