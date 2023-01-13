package repository

// import (
// 	"github.com/jmoiron/sqlx"
// )

// type Exercise interface {
// 	Create() (int, error)
// 	Update(exerciseId int) error
// 	Delete(exerciseId int) error
// }

// type Repository struct {
// 	ExerciseRepository
// }

// func New(db *sqlx.DB) *Repository {
// 	return &Repository {ExerciseRepository: NewExerciseRepository(db)}
// }