package middlewares

import (
	"project/mvc/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

func GenerateTokenJWT(id int) (string, error) {
	claims := JwtCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString([]byte(constants.SECRET_JWT_USER))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func GetClaimsUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)
	return int(userId)
}
