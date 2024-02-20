package repository

import (
	"database/sql"

	"github.com/alekslesik/neuro-news/internal/app/model"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

type MySQLImageRepository struct {
	db *sql.DB
	l  *logger.Logger
}

type ImageQueries struct {
	insert string
}

var imageQueries = ImageQueries{
	insert: `INSERT INTO image
	(image_path, image_size, image_name, image_alt)
	VALUES(?, ?, ?, ?);`,
}

func (ir *MySQLImageRepository) SaveImageToDB(model.Image) error {
	return nil
}
