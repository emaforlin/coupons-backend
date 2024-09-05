package handlers

import (
	"fmt"
	"net/http"

	"github.com/emaforlin/coupons-app/internal/usecases"
	"github.com/emaforlin/coupons-app/pkg/models"
	"github.com/labstack/echo/v4"
)

type accountHttpHandler struct {
	accountUsecase usecases.AccountUsecase
}

func (a *accountHttpHandler) Login(c echo.Context) error {
	reqBody := &models.Login{}
	if err := c.Bind(reqBody); err != nil {
		return response(c, http.StatusBadRequest, "error binding body")
	}

	if err := c.Validate(reqBody); err != nil {
		return response(c, http.StatusBadRequest, "missing required fields")
	}
	_, err := a.accountUsecase.Authenticate(reqBody)
	if err != nil {
		return response(c, http.StatusUnauthorized, "error unauthorized")
	}
	token, err := a.accountUsecase.Authorize(reqBody)
	if err != nil {
		return response(c, http.StatusUnauthorized, "error unauthorized")
	}
	return response(c, http.StatusOK, fmt.Sprintf("successfully logged in, token %s", token))
}

func (a *accountHttpHandler) SignupFoodPlace(c echo.Context) error {
	reqBody := new(models.AddFoodPlaceAccountData)

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
