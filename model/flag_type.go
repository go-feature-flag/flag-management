package model

import (
	"errors"
	"fmt"
)

type FlagType string

const (
	FlagTypeBoolean FlagType = "boolean"
	FlagTypeString  FlagType = "string"
	FlagTypeInteger FlagType = "integer"
	FlagTypeDouble  FlagType = "double"
	FlagTypeJSON    FlagType = "json"
)

// FlagTypeFromValue converts a string to a FlagType
func FlagTypeFromValue(typeAsStr string) (FlagType, error) {
	switch typeAsStr {
	case "boolean":
		return FlagTypeBoolean, nil
	case "string":
		return FlagTypeString, nil
	case "integer":
		return FlagTypeInteger, nil
	case "double":
		return FlagTypeDouble, nil
	case "json":
		return FlagTypeJSON, nil
	case "":
		return "", errors.New("flag type is required")
	default:
		return "", fmt.Errorf("flag type %s not supported", typeAsStr)
	}
}
