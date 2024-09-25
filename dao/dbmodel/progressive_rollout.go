package dbmodel

import (
	"github.com/google/uuid"
	"time"
)

type ProgressiveRollout struct {
	ID                uuid.UUID `db:"id"`
	InitialVariation  string    `db:"initial_variation"`
	EndVariation      string    `db:"end_variation"`
	InitialPercentage float64   `db:"initial_percentage"`
	EndPercentage     float64   `db:"end_percentage"`
	StartDate         time.Time `db:"start_date"`
	EndDate           time.Time `db:"end_date"`
	CreatedDate       time.Time `db:"created_date"`
	LastUpdatedDate   time.Time `db:"last_updated_date"`
}
