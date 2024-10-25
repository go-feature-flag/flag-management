package handler

import (
	"net/http"

	"github.com/go-feature-flag/app-api/dao"
	"github.com/go-feature-flag/app-api/model"
	"github.com/labstack/echo/v4"
)

type Health struct {
	dao dao.Flags
}

// NewHealth creates a new instance of the Health handlers
func NewHealth(dao dao.Flags) Health {
	return Health{dao: dao}
}

type successResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// Health is the health endpoint of the API
// @Summary      Health endpoint of the API
// @Tags Feature Monitoring
// @Description  Check if the API is up and running and that the database is available.
// @Success      200  {object} successResponse "Created"
// @Failure      500 {object} model.HTTPError "Internal server error"
// @Router       /health [get]
func (f Health) Health(c echo.Context) error {
	err := f.dao.Ping()
	if err != nil {
		return c.JSON(model.NewHTTPError(http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusOK, successResponse{
		Message: "API is up and running",
		Code:    http.StatusOK,
	})
}
