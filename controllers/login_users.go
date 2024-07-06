package controllers

import (
	"net/http"
	"practice/my_middleware"

	"github.com/labstack/echo/v4"
	"practice/cookies"
	"practice/models"
	"practice/utils"
)

// LoginUser godoc
// @Summary Login a user
// @Description Authenticate user and set JWT cookie
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserWeb true "User info"
// @Success 200 {string} string "success"
// @Failure 400 {string} string "invalid request"
// @Failure 401 {string} string "unauthorized"
// @Router /login [post]
func LoginUser(c echo.Context) error {
	user := new(models.UserWeb)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}

	db := my_middleware.GetDB(c.Request().Context())
	var foundUser models.UserDTO
	if err := db.Where("login = ?", user.Login).First(&foundUser).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "invalid login or password"})
	}

	if !utils.CheckPasswordHash(user.Password, foundUser.Password) {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "invalid login or password"})
	}

	token, err := cookies.CreateJWTToken(foundUser.Login)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "failed to create token"})
	}
	cookies.SetJWTCookie(c, token)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
