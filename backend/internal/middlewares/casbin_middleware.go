package middlewares

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

func CasbinMiddleware(enforcer *casbin.Enforcer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := c.Request().Header.Get("X-User-ID")
			path := c.Request().URL.Path
			method := c.Request().Method

			allowed, err := enforcer.Enforce(userID, path, method)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			if !allowed {
				return echo.NewHTTPError(http.StatusForbidden, "Permission denied")
			}

			return next(c)
		}
	}
}
