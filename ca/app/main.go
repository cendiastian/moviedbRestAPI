package main

import (
	"log"
	_middleware "project/ca/app/middlewares"
	"project/ca/app/routes"

	_apiRepo "project/ca/drivers/thirdparty/omdb"

	_genreUC "project/ca/business/genres"
	_GenreCtrl "project/ca/controllers/genres"
	_genreRepo "project/ca/drivers/databases/genres"
	_genredb "project/ca/drivers/databases/genres"

	_movieUC "project/ca/business/movies"
	_movieCtrl "project/ca/controllers/movies"
	_movieRepo "project/ca/drivers/databases/movies"
	_moviedb "project/ca/drivers/databases/movies"

	_payUC "project/ca/business/payments"
	_payCtrl "project/ca/controllers/payments"
	_payRepo "project/ca/drivers/databases/payments"
	_paydb "project/ca/drivers/databases/payments"

	_ratingUC "project/ca/business/ratings"
	_ratingCtrl "project/ca/controllers/ratings"
	_ratingRepo "project/ca/drivers/databases/ratings"
	_ratingdb "project/ca/drivers/databases/ratings"

	_subsUC "project/ca/business/subscription"
	_subsCtrl "project/ca/controllers/subscription"
	_subsRepo "project/ca/drivers/databases/subscription"
	_subsdb "project/ca/drivers/databases/subscription"

	_transUC "project/ca/business/transactions"
	_transCtrl "project/ca/controllers/transactions"
	_transRepo "project/ca/drivers/databases/transactions"
	_transdb "project/ca/drivers/databases/transactions"

	_userUC "project/ca/business/users"
	_userCtrl "project/ca/controllers/users"
	_userRepo "project/ca/drivers/databases/users"

	_userdb "project/ca/drivers/databases/users"

	_mysqlDriver "project/ca/drivers/mysql"

	// _genreUC "project/ca/business/genres"
	// _apiUC "project/ca/business/omdb"

	// _apidb "project/ca/drivers/databases/omdb"
	// _genreCtrl "project/ca/controllers/genres"
	// _apiCtrl "project/ca/controllers/apicription"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service Run on Debug mode")
	}
}

func DB_Migrate(db *gorm.DB) {
	db.AutoMigrate(&_transdb.Transaction{})
	db.AutoMigrate(&_ratingdb.Ratings{})
	db.AutoMigrate(&_userdb.Users{})
	db.AutoMigrate(&_paydb.Payment_method{})
	db.AutoMigrate(&_subsdb.SubcriptionPlan{})
	db.AutoMigrate(&_genredb.Genres{})
	db.AutoMigrate(&_moviedb.Movies{})
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

	apiRepo := _apiRepo.NewMysqlAPIRepository(Connect)
	GenreRepo := _genreRepo.NewMysqlGenreRepository(Connect)
	GenreUC := _genreUC.NewGenreUsecase(GenreRepo, timeoutContext)
	GenreCtrl := _GenreCtrl.NewGenreController(GenreUC)

	movieRepo := _movieRepo.NewMysqlMovieRepository(Connect)
	movieUC := _movieUC.NewMovieUsecase(movieRepo, timeoutContext, GenreRepo, apiRepo)
	movieCtrl := _movieCtrl.NewMovieController(movieUC)

	subsRepo := _subsRepo.NewMysqlsubsRepository(Connect)
	subsUC := _subsUC.NewSubsUsecase(subsRepo, timeoutContext)
	subsCtrl := _subsCtrl.NewSubcriptionController(subsUC)

	transRepo := _transRepo.NewMysqlTransRepository(Connect)
	transUC := _transUC.NewTransUsecase(transRepo, timeoutContext)
	transCtrl := _transCtrl.NewTransController(transUC)

	payRepo := _payRepo.NewMysqlpayRepository(Connect)
	payUC := _payUC.NewPaymentUsecase(payRepo, timeoutContext)
	payCtrl := _payCtrl.NewPaymentController(payUC)

	rateRepo := _ratingRepo.NewMysqlRatingRepository(Connect)
	rateUC := _ratingUC.NewRateUsecase(rateRepo, timeoutContext)
	rateCtrl := _ratingCtrl.NewRatingController(rateUC)

	routesInit := routes.ControllerList{
		UserController:        *userCtrl,
		MovieController:       *movieCtrl,
		SubcriptionController: *subsCtrl,
		TransController:       *transCtrl,
		PaymentController:     *payCtrl,
		RatingController:      *rateCtrl,
		GenreController:       *GenreCtrl,
		JwtConfig:             configJWT.Init(),
	}

	routesInit.RouteRegister(e)
	log.Fatal((e.Start(viper.GetString("server.address"))))
}
