package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Db *sqlx.DB
}

type UserList struct {
	Id          int    `json:"id" db:"id"`
}

func (r *Repository) GetUserAll() ([]UserList, error)  {
	var list []UserList
	query := "SELECT id FROM users"
	if err := r.Db.Select(&list, query); err != nil {
		return nil, err
	}
	return list, nil
}