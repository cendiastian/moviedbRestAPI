package main

import (
	"log"
	_middleware "project/app/middlewares"
	"project/app/routes"

	_apiRepo "project/drivers/thirdparty/omdb"

	_genreUC "project/business/genres"
	_GenreCtrl "project/controllers/genres"
	_genreRepo "project/drivers/databases/genres"
	_genredb "project/drivers/databases/genres"

	_movieUC "project/business/movies"
	_movieCtrl "project/controllers/movies"
	_movieRepo "project/drivers/databases/movies"
	_moviedb "project/drivers/databases/movies"

	_payUC "project/business/payments"
	_payCtrl "project/controllers/payments"
	_payRepo "project/drivers/databases/payments"
	_paydb "project/drivers/databases/payments"

	_ratingUC "project/business/ratings"
	_ratingCtrl "project/controllers/ratings"
	_ratingRepo "project/drivers/databases/ratings"
	_ratingdb "project/drivers/databases/ratings"

	_subsUC "project/business/subscription"
	_subsCtrl "project/controllers/subscription"
	_subsRepo "project/drivers/databases/subscription"
	_subsdb "project/drivers/databases/subscription"

	_transUC "project/business/transactions"
	_transCtrl "project/controllers/transactions"
	_transRepo "project/drivers/databases/transactions"
	_transdb "project/drivers/databases/transactions"

	_userUC "project/business/users"
	_userCtrl "project/controllers/users"
	_userRepo "project/drivers/databases/users"
	_userdb "project/drivers/databases/users"

	_proRepo "project/drivers/databases/premium"
	_prodb "project/drivers/databases/premium"

	_mongodbDriver "project/drivers/mongodb"
	_mysqlDriver "project/drivers/mysql"

	// _omdb "project/thirdparty/omdb"

	// _genreUC "project/business/genres"
	// _apiUC "project/business/omdb"

	// _apidb "project/drivers/databases/omdb"
	// _genreCtrl "project/controllers/genres"
	// _apiCtrl "project/controllers/apicription"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app\config.json`)
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
	db.AutoMigrate(&_prodb.Premium{})
}

func main() {

	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	mongoConfig := _mongodbDriver.ConfigDb{
		DbHost: viper.GetString(`mongodb.host`),
		DbPort: viper.GetString(`mongodb.port`),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	OmdbApi := _apiRepo.Omdb{
		Url:    viper.GetString(`omdb.url`),
		Key:    viper.GetString(`omdb.key`),
		Symbol: viper.GetString(`omdb.symbol`),
	}

	OmdbApiRepo := _apiRepo.NewOmdbAPI(OmdbApi)

	Connect := configDB.InitialDB()

	LogCol := _middleware.InitialCollection(struct {
		DbName     string
		Collection string
	}{
		DbName:     viper.GetString(`mongodb.dbname`),
		Collection: viper.GetString(`mongodb.collection`),
	})

	InitMongo := mongoConfig.InitialDb()
	LoggerMiddleware := _middleware.InitialConfig(InitMongo, LogCol)

	DB_Migrate(Connect)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	proRepo := _proRepo.NewMysqlPremiumRepository(Connect)

	userRepo := _userRepo.NewMysqlUserRepository(Connect)
	userUC := _userUC.NewUserUsecase(userRepo, proRepo, timeoutContext, &configJWT)
	userCtrl := _userCtrl.NewUserController(userUC)

	// apiRepo := _apiRepo.NewMysqlAPIRepository(Connect)
	GenreRepo := _genreRepo.NewMysqlGenreRepository(Connect)
	GenreUC := _genreUC.NewGenreUsecase(GenreRepo, timeoutContext)
	GenreCtrl := _GenreCtrl.NewGenreController(GenreUC)

	movieRepo := _movieRepo.NewMysqlMovieRepository(Connect)
	movieUC := _movieUC.NewMovieUsecase(movieRepo, timeoutContext, GenreRepo, OmdbApiRepo, proRepo)
	movieCtrl := _movieCtrl.NewMovieController(movieUC)

	subsRepo := _subsRepo.NewMysqlsubsRepository(Connect)
	subsUC := _subsUC.NewSubsUsecase(subsRepo, timeoutContext)
	subsCtrl := _subsCtrl.NewSubcriptionController(subsUC)

	transRepo := _transRepo.NewMysqlTransRepository(Connect)
	transUC := _transUC.NewTransUsecase(transRepo, timeoutContext, proRepo, subsRepo)
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
		LoggerMiddleware:      *LoggerMiddleware,
		JwtConfig:             configJWT.Init(),
	}

	routesInit.RouteRegister(e)
	log.Fatal((e.Start(viper.GetString("server.address"))))
}
