package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/config"
	"github.com/sawalreverr/recything/internal/helper"
)

func SuperAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return jwtMiddleware(c, next, "superadmin")
	}
}

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return jwtMiddleware(c, next, "admin")
	}
}

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return jwtMiddleware(c, next, "user")
	}
}

func jwtMiddleware(c echo.Context, next echo.HandlerFunc, requiredRole string) error {
	secretKey := config.GetConfig().Server.JWTSecret

	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return helper.ErrorHandler(c, http.StatusUnauthorized, "token is not provided")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid token format. use bearer token")
	}
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenStr, &helper.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return helper.ErrorHandler(c, http.StatusUnauthorized, "invalid token signature")
	}

	if claims, ok := token.Claims.(*helper.JwtCustomClaims); ok && next != nil {
		if claims.Role != requiredRole {
			return helper.ErrorHandler(c, http.StatusUnauthorized, "unauthorized")
		}
		c.Set("user", claims)
		return next(c)
	}

	return helper.ErrorHandler(c, http.StatusUnauthorized, "unauthorized")
}
