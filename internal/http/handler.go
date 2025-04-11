package http

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func (*Handler) currentUserID(c echo.Context) (uint64, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return 0, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrInvalidToken
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return 0, ErrInvalidToken
	}

	return uint64(userID), nil
}
