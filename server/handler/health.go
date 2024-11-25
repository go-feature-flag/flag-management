package handler

import (
	"fmt"
	"github.com/go-feature-flag/flag-management/server/dao"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	dao dao.FlagStorage
}

// NewHealthHandler creates a new instance of the HealthHandler handlers
func NewHealthHandler(dao dao.FlagStorage) HealthHandler {
	return HealthHandler{dao: dao}
}

type successResponse struct {
	Message string `json:"message" example:"API is up and running"`
	Code    int    `json:"code" example:"200"`
}

// Health is the health endpoint of the API
// @Summary      Health endpoint of the API
// @Tags Feature Monitoring
// @Description  Check if the API is up and running and that the database is available.
// @Success      200  {object} successResponse "Created"
// @Failure      500 {object} model.ErrorResponse "Internal server error"
// @Router       /health [get]
func (f HealthHandler) Health(c echo.Context) error {
	err := f.dao.Ping()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	err1 := c.JSON(http.StatusOK, successResponse{
		Message: "API is up and running",
		Code:    http.StatusOK,
	})
	fmt.Println(err1)
	return err1
}
