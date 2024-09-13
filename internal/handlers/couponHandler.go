package handlers

import "github.com/labstack/echo/v4"

type CouponHandler interface {
	CreateCoupon(c echo.Context) error
	RetrieveCoupons(c echo.Context) error
	UpdateCoupon(c echo.Context) error
	DeleteCoupon(c echo.Context) error
	ClaimCoupon(c echo.Context) error
	UseCoupon(c echo.Context) error
}
