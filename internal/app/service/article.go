package service

import (
	"database/sql"
	"strings"

	"github.com/alekslesik/neuro-news/internal/app/model"
	"github.com/go-sql-driver/mysql"
)

type ArticleModel struct {
	DB *sql.DB
}

// Add a new record to the users table.
func (m *ArticleModel) Insert(title, body, image string) error {
	// Create a bcrypt hash of the plain-text password.
	// SQL request we wanted to execute.
	stmt := `INSERT INTO users (title, body, image, created)
    VALUES(?, ?, ?, UTC_TIMESTAMP())`

	// Use the Exec() method to insert the user details and hashed password
	// into the users table. If this returns an error, we try to type assert
	// it to a *mysql.MySQLError object so we can check if the error number is
	// 1062 and, if it is, we also check whether or not the error relates to
	// our users_uc_email key by checking the contents of the message string.
	// If it does, we return an ErrDuplicateEmail error. Otherwise, we just
	// return the original error (or nil if everything worked).
	_, err := m.DB.Exec(stmt, title, body, image)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			// if mysqlErr.Number == 1062 {
			// 	return models.ErrDuplicateEmail
			// }
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "Duplicate entry") {
				return models.ErrDuplicateEmail
			}
		}
	}

	return err
}


// Fetch details for a specific user based on their user ID.
func (m *ArticleModel) Get(id int) (*model.Article, error) {
	s := &model.Article{}

	stmt := `SELECT id, name, email, created FROM users WHERE id = ?`
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Body, &s.Created)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}
