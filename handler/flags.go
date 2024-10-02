package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-feature-flag/app-api/dao"
	"github.com/go-feature-flag/app-api/model"
	"github.com/labstack/echo/v4"
)

type Flags struct {
	dao dao.Flags
}

func NewFlags(dao dao.Flags) Flags {
	return Flags{dao: dao}
}

func (f Flags) GetAllFeatureFlags(c echo.Context) error {
	flags, err := f.dao.GetFlags(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, flags)
}

func (f Flags) GetFeatureFlagsByID(c echo.Context) error {
	flag, err := f.dao.GetFlagByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, flag)
}

func (f Flags) CreateNewFlag(c echo.Context) error {
	var flag model.FeatureFlag
	if err := c.Bind(&flag); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Check if flag with this name exists
	res, err := f.dao.GetFlagByName(c.Request().Context(), flag.Name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if res.ID != "" {
		return c.JSON(http.StatusConflict, map[string]string{"error": "flag with this name already exists"})
	}

	// Create the flag
	flag.CreatedDate = time.Now()
	// TODO: remove this line
	flag.LastModifiedBy = "toto"

	id, err := f.dao.CreateFlag(c.Request().Context(), flag)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	flag.ID = id
	return c.JSON(http.StatusOK, flag)
}

func (f Flags) UpdateFlagByID(c echo.Context) error {
	// check if the flag exists
	_, err := f.dao.GetFlagByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	// update the flag
	var flag model.FeatureFlag
	if err := c.Bind(&flag); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if flag.ID == "" {
		flag.ID = c.Param("id")
	}

	flag.LastUpdatedDate = time.Now()

	err = f.dao.UpdateFlag(c.Request().Context(), flag)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, flag)
}

func (f Flags) DeleteFlagByID(c echo.Context) error {
	err := f.dao.DeleteFlagByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}

func (f Flags) UpdateFeatureFlagStatus(c echo.Context) error {
	flag, err := f.dao.GetFlagByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	// FeatureFlagStatusUpdate represents the input for updating the status of a feature flag.
	type FeatureFlagStatusUpdate struct {
		Disable bool `json:"disable"`
	}

	var statusUpdate FeatureFlagStatusUpdate
	if err := c.Bind(&statusUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	flag.Disable = &statusUpdate.Disable
	flag.LastUpdatedDate = time.Now()
	err = f.dao.UpdateFlag(c.Request().Context(), flag)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, flag)
}
