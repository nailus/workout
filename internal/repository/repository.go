package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nailus/workout/internal/entity"
	"fmt"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllExercises() ([]entity.Exercise, error) {
	exerciseList := []entity.Exercise{}
	err := r.db.Select(&exerciseList, "SELECT * FROM exercises")
	return exerciseList, err
}

func (r *Repository) CreateUser(user *entity.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (email, encrypted_password) values ($1, $2) RETURNING id", "users")

	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetUser(email string) (*entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", "users")
	err := r.db.Get(&user, query, email)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateExercise(exercise *entity.Exercise, userId int) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (title, body, author_id) values ($1, $2, $3) RETURNING id", "exercises")

	row := r.db.QueryRow(query, exercise.Title, exercise.Body, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}