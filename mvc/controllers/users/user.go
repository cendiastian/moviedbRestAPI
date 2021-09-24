package users

import (
	"net/http"
	"project/mvc/config"
	"project/mvc/middlewares"
	"project/mvc/model/response"
	"project/mvc/model/user"
	"strconv"

	// "project/mvc/routes"

	"github.com/labstack/echo/v4"
)

func UserRegist(c echo.Context) error {
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
	if result.Error == nil {
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

func GetUser(c echo.Context) error {
	users := []user.User{}
	result := config.DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika input mendapatkan data user dari DB",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data user",
		Data:    users,
	})
}

func UserDetail(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("userId"))
	var userDB user.User

	result := config.DB.Model(&userDB).Where("Id = ?", userId).Find(&userDB)
	if result.Error == nil {
		if userDB.ID == 0 {
			return c.JSON(http.StatusInternalServerError, response.Response{
				Code:    http.StatusInternalServerError,
				Message: "User ID tidak ada",
				Data:    nil,
			})
		}

		return c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Message: "Berhasil",
			Data:    userDB,
		})
	}
	return c.JSON(http.StatusInternalServerError, response.Response{
		Code:    http.StatusInternalServerError,
		Message: "Gagal konversi userId",
		Data:    nil,
	})
}

func UserLogin(c echo.Context) error {
	userLogin := user.UserLogin{}
	c.Bind(&userLogin)
	var userDB user.User
	if userLogin.Email == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "Mohon isi Email",
			Data:    nil,
		})
	} else if userLogin.Password == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "Mohon isi Password",
			Data:    nil,
		})
	}

	result := config.DB.Where("email = ? AND password = ?", userLogin.Email, userLogin.Password).First(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Login gagal",
			Data:    nil,
		})
	} else if userDB.ID == 0 {
		return c.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Login gagal",
			Data:    nil,
		})
	}

	token, err := middlewares.GenerateTokenJWT(userDB.ID)

	if err != nil {
		return c.JSON(http.StatusForbidden, response.Response{
			Code:    http.StatusForbidden,
			Message: "Error ketika membuat Token",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Message: "Login Berhasil! Token: " + token,
		Data:    userLogin,
	})
}

func UpdateUser(c echo.Context) error {
	var userUpdate user.UserUpdate
	c.Bind(&userUpdate)
	// var userDB user.User

	result := config.DB.Model(&user.User{}).Where("id = ?", userUpdate.ID).Updates(&user.User{Name: userUpdate.Name, Password: userUpdate.Password, Email: userUpdate.Email})
	if result.Error == nil {
		return c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Message: "Berhasil",
			Data:    userUpdate,
		})

	}
	return c.JSON(http.StatusInternalServerError, response.Response{
		Code:    http.StatusInternalServerError,
		Message: "Update gagal",
		Data:    nil,
	})
}

func DeleteUser(c echo.Context) error {
	var userDelete user.UserDelete
	c.Bind(&userDelete)
	// var userDB user.User

	result := config.DB.Delete(&user.User{}, userDelete.ID)
	if result.Error == nil {
		// if userDB.ID == 0 {
		// 	return c.JSON(http.StatusInternalServerError, response.Response{
		// 		Code:    http.StatusInternalServerError,
		// 		Message: "Gagal menghapus user",
		// 		Data:    nil,
		// 	})
		// }
		return c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Message: "Berhasil",
			Data:    userDelete,
		})
	}
	return c.JSON(http.StatusInternalServerError, response.Response{
		Code:    http.StatusInternalServerError,
		Message: "Gagal menghapus user",
		Data:    nil,
	})
}
