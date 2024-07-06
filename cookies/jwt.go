package cookies

import (
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
)

type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

func CreateJWTToken(login string) (string, error) {
	claims := &Claims{
		Login: login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func SetJWTCookie(c echo.Context, token string) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	c.SetCookie(cookie)
}

func GetJWTFromCookie(c echo.Context) (*jwt.Token, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, err
}

func ClearJWTCookie(c echo.Context) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
	}
	c.SetCookie(cookie)
}
