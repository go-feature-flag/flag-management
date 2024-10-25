package handler

import (
	"database/sql"
	"errors"
	"fmt"
	daoErr "github.com/go-feature-flag/app-api/dao/err"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"

	"github.com/go-feature-flag/app-api/dao"
	"github.com/go-feature-flag/app-api/model"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type FlagAPIHandler struct {
	dao dao.Flags
}

// NewFlagAPIHandler creates a new instance of the FlagAPIHandler handler
// It is a controller class to handle the feature flag configuration logic
func NewFlagAPIHandler(dao dao.Flags) FlagAPIHandler {
	return FlagAPIHandler{dao: dao}
}

// GetAllFeatureFlags is returning the list of all the flags
// @Summary      Return all the flags available
// @Tags Feature Flag management API
// @Description  GET request to get all the flags available.
// @Success      200  {object} []model.FeatureFlag "Success"
// @Failure      500 {object} model.HTTPError "Internal server error"
// @Router       /v1/flags [get]
func (f FlagAPIHandler) GetAllFeatureFlags(c echo.Context) error {
	flags, err := f.dao.GetFlags(c.Request().Context())
	if err != nil {
		return c.JSON(model.NewHTTPError(http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusOK, flags)
}

// GetFeatureFlagByID is returning the flag belonging to the given ID
// @Summary      Return all the information about a flag
// @Tags Feature Flag management API
// @Description  GET all the information about a flag with a specific .
// @Param        id path string true "ID of the feature flag"
// @Success      200  {object} model.FeatureFlag "Success"
// @Failure      404 {object} model.HTTPError "Not Found"
// @Failure      500 {object} model.HTTPError "Internal server error"
// @Router       /v1/flags/{id} [get]
func (f FlagAPIHandler) GetFeatureFlagByID(c echo.Context) error {
	flag, err := f.dao.GetFlagByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return f.handleDaoError(c, err)
	}
	return c.JSON(http.StatusOK, flag)
}

// CreateNewFlag is creating a new flag
// @Summary      Create a new feature flag with the given configuration.
// @Tags Feature Flag management API
// @Description  POST will insert in the database the new feature flag with all his properties,
// @Description  and it will add all the associated rules too.
// @Param 		 data body model.FeatureFlag true "Payload which represents the flag to insert"
// @Success      201  {object} model.FeatureFlag "Created"
// @Failure      400 {object} model.HTTPError "Bad Request"
// @Failure      409 {object} model.HTTPError "Conflict - when trying to insert a flag with a name that already exists"
// @Failure      500 {object} model.HTTPError "Internal server error"
// @Router       /v1/flags [post]
func (f FlagAPIHandler) CreateNewFlag(c echo.Context) error {
	var flag model.FeatureFlag
	if err := c.Bind(&flag); err != nil {
		return c.JSON(model.NewHTTPError(http.StatusBadRequest, err))
	}

	// Check if flag with this name exists
	_, err := f.dao.GetFlagByName(c.Request().Context(), flag.Name)
	if err == nil {
		return c.JSON(model.NewHTTPError(http.StatusConflict, fmt.Errorf("flag with name %s already exists", flag.Name)))
	}
	if err != nil && err.Code() != daoErr.NotFound {
		return c.JSON(model.NewHTTPError(http.StatusInternalServerError, err))
	}

	// Add field that are not in the request
	if flag.ID == "" {
		flag.ID = uuid.NewString()
	}
	flag.CreatedDate = time.Now()
	flag.LastUpdatedDate = time.Now()
	// TODO: remove this line and extract the information from the token
	flag.LastModifiedBy = "toto"

	if code, err := validateFlag(flag); err != nil {
		return c.JSON(model.NewHTTPError(code, err))
	}
	/**
	TODO: Add a validation layer here, it should check:
	- the flag name is not empty
	- that reference from variation in rule exists in variations
	- the type of the variations is correct
	- ...
	*/

	id, err := f.dao.CreateFlag(c.Request().Context(), flag)
	if err != nil {
		return c.JSON(model.NewHTTPError(http.StatusInternalServerError, err))
	}
	flag.ID = id

	// TODO: Check what to return here because it has not all the new id created in the DAO (example rule ID)
	return c.JSON(http.StatusCreated, flag)
}

func validateFlag(flag model.FeatureFlag) (int, error) {
	// Check if the flag name is valid
	if flag.Name == "" {
		return http.StatusBadRequest, errors.New("flag name is required")
	}

	if status, err := validateRule(flag.DefaultRule, true); err != nil {
		return status, err
	}

	if flag.VariationType == "" {
		return http.StatusBadRequest, errors.New("flag type is required")
	}

	for _, rule := range flag.GetRules() {
		if status, err := validateRule(&rule, false); err != nil {
			return status, err
		}
	}

	return http.StatusOK, nil
}

func validateRule(rule *model.Rule, isDefault bool) (int, error) {
	if rule == nil ||
		(rule.ProgressiveRollout == nil &&
			rule.Percentages == nil &&
			(rule.VariationResult == nil || *rule.VariationResult == "")) {
		err := fmt.Errorf("invalid rule %s", rule.Name)
		if isDefault {
			err = errors.New("flag default rule is invalid")
		}
		return http.StatusBadRequest, err
	}

	if !isDefault {
		if rule.Query == "" {
			return http.StatusBadRequest, errors.New("rule query is required")
		}
	}
	return http.StatusOK, nil
}

// UpdateFlagByID is updating the flag with the given ID
// @Summary      Updates the flag with the given ID
// @Tags Feature Flag management API
// @Description  PUT - Updates the flag with the given ID with what is in the payload. It will replace completely the feature flag.
// @Param        id path string true "ID of the feature flag"
// @Param 		 data body model.FeatureFlag true "Payload which represents the flag to update"
// @Success      200  {object} model.FeatureFlag "Success"
// @Failure      400 {object} model.HTTPError "Bad Request"
// @Failure      404 {object} model.HTTPError "Not Found"
// @Failure      500 {object} model.HTTPError "Internal server error"
// @Router       /v1/flags/{id} [put]
func (f FlagAPIHandler) UpdateFlagByID(c echo.Context) error {
	// check if the flag exists
	_, err := f.dao.GetFlagByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return f.handleDaoError(c, err)
	}

	// update the flag
	var flag model.FeatureFlag
	if err := c.Bind(&flag); err != nil {
		return c.JSON(model.NewHTTPError(http.StatusBadRequest, err))
	}

	if code, err := validateFlag(flag); err != nil {
		return c.JSON(model.NewHTTPError(code, err))
	}

	if flag.ID == "" {
		flag.ID = c.Param("id")
	}
	flag.LastUpdatedDate = time.Now()

	err = f.dao.UpdateFlag(c.Request().Context(), flag)
	if err != nil {
		return c.JSON(model.NewHTTPError(http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusOK, flag)
}

// DeleteFlagByID is deleting the flag with the given ID
// @Summary      Delete the flag with the given ID
// @Tags Feature Flag management API
// @Description  DELETE - Delete the flag with the given ID.
// @Param        id path string true "ID of the feature flag"
// @Success      204  {object} model.FeatureFlag "No Content"
// @Failure      400 {object} model.HTTPError "Bad Request"
// @Failure      404 {object} model.HTTPError "Not Found"
// @Failure      500 {object} model.HTTPError "Internal server error"
// @Router       /v1/flags/{id} [delete]
func (f FlagAPIHandler) DeleteFlagByID(c echo.Context) error {
	idParam := c.Param("id")
	err := f.dao.DeleteFlagByID(c.Request().Context(), idParam)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(model.NewHTTPError(http.StatusNotFound, fmt.Errorf("flag with id %s not found", idParam)))
		}
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "22P02":
				return c.JSON(model.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid UUID format")))
			default:
				return c.JSON(model.NewHTTPError(http.StatusInternalServerError, err))
			}
		}
		return c.JSON(model.NewHTTPError(http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusNoContent, nil)
}

// UpdateFeatureFlagStatus is updating the flag status with the given ID
// @Summary      Update the status of the flag with the given ID
// @Tags Feature Flag management API
// @Description  PATCH - Update the status of the flag with the given ID
// @Param        id path string true "ID of the feature flag"
// @Param 		 data body model.FeatureFlagStatusUpdate true "The patch query to update the flag status"
// @Success      200  {object} model.FeatureFlag "Success"
// @Failure      400 {object} model.HTTPError "Bad Request"
// @Failure      404 {object} model.HTTPError "Not Found"
// @Failure      500 {object} model.HTTPError "Internal server error"
// @Router       /v1/flags/{id}/status [patch]
func (f FlagAPIHandler) UpdateFeatureFlagStatus(c echo.Context) error {
	idParam := c.Param("id")
	flag, err := f.dao.GetFlagByID(c.Request().Context(), idParam)
	if err != nil {
		return f.handleDaoError(c, err)
	}

	var statusUpdate model.FeatureFlagStatusUpdate
	if err := c.Bind(&statusUpdate); err != nil {
		return c.JSON(model.NewHTTPError(http.StatusBadRequest, err))
	}

	flag.Disable = &statusUpdate.Disable
	flag.LastUpdatedDate = time.Now()
	err = f.dao.UpdateFlag(c.Request().Context(), flag)
	if err != nil {
		return c.JSON(model.NewHTTPError(http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusOK, flag)
}

// handleDaoError is a helper function to handle the dao errors and return the correct HTTP status code.
func (f FlagAPIHandler) handleDaoError(c echo.Context, err daoErr.DaoError) error {
	switch err.Code() {
	case daoErr.NotFound:
		return c.JSON(model.NewHTTPError(http.StatusNotFound, fmt.Errorf("flag with id %s not found", c.Param("id"))))
	case daoErr.InvalidUUID:
		return c.JSON(model.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid UUID format")))
	default:
		return c.JSON(model.NewHTTPError(http.StatusInternalServerError, err))
	}
}
