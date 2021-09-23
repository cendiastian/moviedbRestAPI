package main

import (
	"log"
	_middleware "project/ca/app/middlewares"
	"project/ca/app/routes"
	_userUC "project/ca/business/users"
	_userCtrl "project/ca/controllers/users"
	_userRepo "project/ca/drivers/databases/users"
	_userdb "project/ca/drivers/databases/users"
	_mysqlDriver "project/ca/drivers/mysql"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service Run on Debug mode")
	}
}

func DB_Migrate(db *gorm.DB) {
	db.AutoMigrate(&_userdb.Users{})
}

func main() {

	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	Connect := configDB.InitialDB()
	DB_Migrate(Connect)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepo := _userRepo.NewMysqlUserRepository(Connect)
	userUC := _userUC.NewUserUsecase(userRepo, timeoutContext)
	userCtrl := _userCtrl.NewUserController(userUC)

	routesInit := routes.ControllerList{
		UserController: *userCtrl,
		JwtConfig:      configJWT.Init(),
	}

	routesInit.RouteRegister(e)
	log.Fatal((e.Start(viper.GetString("server.address"))))
}
