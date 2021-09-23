package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
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
