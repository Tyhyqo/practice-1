package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Protected godoc
// @Summary Access a protected resource
// @Description Access a protected resource with a valid JWT
// @Tags protected
// @Produce  json
// @Success 200 {string} string "protected resource"
// @Failure 401 {string} string "unauthorized"
// @Router /protected [get]
func Protected(c echo.Context) error {
	user := c.Get("user")
	if user == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorized",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "You have accessed a protected resource",
		"user":    user,
	})
}
