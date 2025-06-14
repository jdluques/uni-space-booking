package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ParseUUIDParam(c echo.Context, param string) (uuid.UUID, error) {
	idStr := c.Param(param)
	if idStr == "" {
		return uuid.Nil, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s is required", param))
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid %s format", param))
	}
	return id, nil
}
