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

// func (r *Repository) GetUser(email, password) (User, error) {
// 	return nil, nil
// }

func (r *Repository) CreateUser(user *entity.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (email, encrypted_password) values ($1, $2) RETURNING id", "users")

	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetUser(user *entity.User) (*entity.User, error) {
	query := fmt.Sprintf("SELECT FROM %s WHERE email = $1 AND encrypted_password = $2", "users")
	err := r.db.Get(&user, query, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
} 