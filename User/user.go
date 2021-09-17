package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	eTest := e.Group("test/")
	eTest.GET("users", getUser)
	eTest.POST("users/login", userLogin)
	eTest.POST("users/register", userRegist)
	eTest.GET("users/:userId", userDetail)
	// eTest.PUT("/users/:id", updateUser)
	// eTest.DELETE("/users/:id", deleteUser)
	e.Start(":8000")
}

var DB *gorm.DB

func InitDB() {
	dsn := "cendiastian:12345.@tcp(127.0.0.1:3306)/moviedb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB failed connect")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&User{})
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type User struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password int    `json:"password"`
}

type UserRegister struct {
	Name  string `json:"name"`
	Email string `json:"email"`

	Password string `json:"password"`
}

var (
	users = map[int]*User{}
	seq   = 1
)

func userRegist(c echo.Context) error {
	var userRegister UserRegister
	c.Bind(&userRegister)

	// validasi
	if userRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Mohon isi Nama",
			Data:    nil,
		})
	} else if userRegister.Password == "" {
		return c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Mohon isi Password",
			Data:    nil,
		})
	} else if userRegister.Email == "" {
		return c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Mohon isi Email",
			Data:    nil,
		})
	}

	var userDB User
	userDB.Name = userRegister.Name
	userDB.Password = userRegister.Password
	userDB.Email = userRegister.Email

	result := DB.Create(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika input data user ke DB",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Berhasil register",
		Data:    userDB,
	})
}

func getUser(c echo.Context) error {
	users := []User{}
	result := DB.Find(&users)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data user dari DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data user",
		Data:    users,
	})
}

func userLogin(c echo.Context) error {
	userLogin := UserLogin{}
	c.Bind(&userLogin)

	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    userLogin,
	})
}

func userDetail(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: "Gagal konversi userId",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    User{ID: userId},
	})
}
