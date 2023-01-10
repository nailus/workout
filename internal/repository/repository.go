package repository

import (
	"github.com/jmoiron/sqlx"
)

type Workout interface {
	Create() (int, error)
	Update(workoutId int) error
	Delete(workoutId int) error
}

type Repository struct {
	Workout
}

func New(db *sqlx.DB) *Repository {
	return &Repository {Workout: repository.NewWorkout(db)}
}