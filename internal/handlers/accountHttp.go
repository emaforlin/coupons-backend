package handlers

import (
	"net/http"

	"github.com/emaforlin/coupons-app/internal/usecases"
	"github.com/emaforlin/coupons-app/pkg/models"
	"github.com/labstack/echo/v4"
)

type accountHttpHandler struct {
	accountUsecase usecases.AccountUsecase
}

func (h *accountHttpHandler) Login(c echo.Context) error {
	reqBody := &models.Login{}
	if err := c.Bind(reqBody); err != nil {
		return errorMsg(c, "bad request")
	}

	if err := c.Validate(reqBody); err != nil {
		return errorMsg(c, err.Error())
	}

	_, err := h.accountUsecase.Authenticate(reqBody)
	if err != nil {
		return errorMsg(c, "unauthorized", http.StatusUnauthorized)
	}

	token, err := h.accountUsecase.Authorize(reqBody)
	if err != nil {
		return errorMsg(c, "unauthorized", http.StatusUnauthorized)
	}
	return responseMsg(c, token, http.StatusOK)
}

func (h *accountHttpHandler) SignupFoodPlace(c echo.Context) error {
	reqBody := new(models.AddFoodPlaceAccountData)

	if err := c.Bind(reqBody); err != nil {
		return errorMsg(c, "bad request")
	}

	if err := c.Validate(reqBody); err != nil {
		return errorMsg(c, err.Error())
	}

	if err := h.accountUsecase.AddFoodPlaceAccount(reqBody); err != nil {
		return errorMsg(c, "error creating food place account")
	}

	return responseMsg(c, "account successfully created", http.StatusCreated)
}

func (h *accountHttpHandler) SignupPerson(c echo.Context) error {
	reqBody := &models.AddPersonAccountData{}

	if err := c.Bind(reqBody); err != nil {
		return errorMsg(c, "bad request")
	}

	if err := h.accountUsecase.AddPersonAccount(reqBody); err != nil {
		return errorMsg(c, "error creating person account")
	}

	return responseMsg(c, "account successfully created", http.StatusCreated)
}

func NewAccountHttpHandler(u usecases.AccountUsecase) AccountHandler {
	return &accountHttpHandler{
		accountUsecase: u,
	}
}
