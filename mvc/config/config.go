package config

import (
	"project/mvc/model/movie"
	"project/mvc/model/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:12345@tcp(127.0.0.1:3306)/movie?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB failed connect")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&user.User{}, &movie.Movie{}, &movie.Category{})
}
