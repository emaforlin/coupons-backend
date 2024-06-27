package handlers

import (
	"net/http"

	"github.com/emaforlin/coupons-app/pkg/models"
	"github.com/emaforlin/coupons-app/pkg/usecases"
	"github.com/labstack/echo/v4"
)

type accountHttpHandler struct {
	accountUsecase usecases.AccountUsecase
}

// SignupFoodPlace implements AccountHandler.
func (a *accountHttpHandler) SignupFoodPlace(c echo.Context) error {
	reqBody := &models.AddFoodPlaceData{}

	if err := c.Bind(reqBody); err != nil {
		return response(c, http.StatusBadRequest, "error binding body")
	}

	if err := c.Validate(reqBody); err != nil {
		return response(c, http.StatusBadRequest, err.Error())
	}

	if err := a.accountUsecase.AddFoodPlaceAccount(reqBody); err != nil {
		return response(c, http.StatusBadRequest, err.Error())
	}

	return response(c, http.StatusCreated, "account successfully created")
}

// RegisterPersonAccount implements AccountHandler.
func (a *accountHttpHandler) SignupPerson(c echo.Context) error {
	reqBody := &models.AddPersonAccountData{}

	if err := c.Bind(reqBody); err != nil {
		return response(c, http.StatusBadRequest, "error binding body")
	}

	if err := c.Validate(reqBody); err != nil {
		return response(c, http.StatusBadRequest, err.Error())
	}

	if err := a.accountUsecase.AddPersonAccount(reqBody); err != nil {
		return response(c, http.StatusBadRequest, err.Error())
	}

	return response(c, http.StatusCreated, "account successfully created")
}

func NewAccountHttpHandler(u usecases.AccountUsecase) AccountHandler {
	return &accountHttpHandler{
		accountUsecase: u,
	}
}
