package handlers

import "github.com/labstack/echo/v4"

type AccountHandler interface {
	Login(c echo.Context) error
	SignupPerson(c echo.Context) error
	SignupFoodPlace(c echo.Context) error
	VerifyFoodPlace(c echo.Context) error
}
