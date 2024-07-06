package my_middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"practice/utils"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": "missing or invalid token"})
			}
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
		}

		tokenStr := cookie.Value
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
		}

		log.Printf("Authenticated user: %s", claims.Login)
		c.Set("user", claims.Login)
		return next(c)
	}
}
