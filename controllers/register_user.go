package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"practice/cookies"
	"practice/cruds"
	"practice/models"
	"practice/my_middleware"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with login and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserWeb true "User info"
// @Success 200 {string} string "success"
// @Failure 400 {string} string "invalid request"
// @Router /register [post]
func RegisterUser(c echo.Context) error {
	user := new(models.UserWeb)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}

	newUser := models.UserDTO{
		Login:     user.Login,
		Password:  user.Password,
		IsCourier: user.IsCourier,
	}
	db := my_middleware.GetDB(c.Request().Context())
	cruds.CrudRegisterUser(newUser, db)

	token, err := cookies.CreateJWTToken(newUser.Login)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "failed to create token"})
	}
	cookies.SetJWTCookie(c, token)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
