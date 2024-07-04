package handlers

import (
	"github.com/emaforlin/coupons-app/pkg/helpers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CheckRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*helpers.CustomJWTClaims)
			if claims.Role != role {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	}
}
