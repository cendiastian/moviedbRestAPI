package middlewares

import (
	"net/http"
	controller "project/ca/controllers"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtMyClaims struct {
	UserId int
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtMyClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return controller.NewErrorResponse(c, http.StatusForbidden, e)
		}),
	}
}

func (jwtConf *ConfigJWT) GenerateToken(UserId int) (string, error) {
	claims := JwtMyClaims{
		UserId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	//buat token dengan claims
	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, err
}

// GetUser from jwt ...
func GetUser(c echo.Context) *JwtMyClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtMyClaims)
	return claims
}
