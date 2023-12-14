package repository

import (
	"database/sql"

	"github.com/alekslesik/neuro-news/internal/app/model"
)

type MySQLRepository struct {
    db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
    return &MySQLRepository{
        db: db,
    }
}

func (r *MySQLRepository) GetArticleRepository() model.ArticleRepository {
    return &MySQLArticleRepository{db: r.db}
}

func (r *MySQLRepository) GetUserRepository() model.UserRepository {
    return &MySQLUserRepository{db: r.db}
}

