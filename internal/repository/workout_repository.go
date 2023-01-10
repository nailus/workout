package repository

import (
	"github.com/jmoiron/sqlx"
)

type Workout struct {
	db *sqlx.DB
}

func NewWorkout(db *sqlx.DB) *Workout {
	return &Workout{ db: db }
}

func