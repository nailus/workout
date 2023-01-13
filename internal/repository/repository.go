package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nailus/workout/internal/entity"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository {db: db}
}

func (r *Repository) GetAllExercises() ([]entity.Exercise, error) {
	exerciseList := []entity.Exercise{}
  err := r.db.Select(&exerciseList, "SELECT * FROM exercises")
	return exerciseList, err
}