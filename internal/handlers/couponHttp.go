package handlers

import (
	"fmt"
	"net/http"

	"github.com/emaforlin/coupons-app/internal/usecases"
	"github.com/emaforlin/coupons-app/pkg/models"
	"github.com/labstack/echo/v4"
)

type couponHttpHandler struct {
	usecase usecases.CouponUsecase
}

// ClaimCoupon implements CouponHandler.
func (h *couponHttpHandler) ClaimCoupon(c echo.Context) error {
	panic("unimplemented")
}

// CreateCoupon implements CouponHandler.
func (h *couponHttpHandler) CreateCoupon(c echo.Context) error {
	reqBody := new(models.AddCoupon)

	if err := c.Bind(reqBody); err != nil {
		return errorMsg(c, "bad request")
	}
	id, err := h.usecase.Create(reqBody)

	if err != nil {
		return errorMsg(c, "error creating coupon")
	}
	return responseMsg(c, fmt.Sprintf("%d", id), http.StatusCreated)
}

// DeleteCoupon implements CouponHandler.
func (*couponHttpHandler) DeleteCoupon(c echo.Context) error {
	panic("unimplemented")
}

// RetrieveCoupons implements CouponHandler.
func (h *couponHttpHandler) RetrieveCoupons(c echo.Context) error {
	reqBody := new(models.GetCoupons)

	if err := c.Bind(reqBody); err != nil {
		return errorMsg(c, "bad request")
	}

	coupons, err := h.usecase.Get(reqBody)
	if err != nil {
		return errorMsg(c, "error retrieving coupons")
	}

	if len(coupons) < 1 {
		return errorMsg(c, "not found", http.StatusNotFound)
	} else {
		return c.JSON(http.StatusFound, coupons)
	}
}

// UpdateCoupon implements CouponHandler.
func (*couponHttpHandler) UpdateCoupon(c echo.Context) error {
	panic("unimplemented")
}

// UseCoupon implements CouponHandler.
func (*couponHttpHandler) UseCoupon(c echo.Context) error {
	panic("unimplemented")
}

func NewCouponHttpHandler(u usecases.CouponUsecase) CouponHandler {
	return &couponHttpHandler{
		usecase: u,
	}
}
