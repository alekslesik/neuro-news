package repository

import (
	"database/sql"

	"github.com/alekslesik/neuro-news/internal/app/model"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func (r *MySQLUserRepository) GetUserByID(id int) (*model.User, error) {
	return nil, nil
}
