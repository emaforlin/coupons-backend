package helpers

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type CustomJWTClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func (c *CustomJWTClaims) Valid() error {
	switch c.Role {
	case "Customer":
		return nil
	case "FoodPlace":
		return nil
	case "adm":
		return nil
	default:
		return errors.New("invalid role")
	}
}
