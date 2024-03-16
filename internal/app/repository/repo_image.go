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

func (ir *MySQLImageRepository) SaveImageToDB(image *model.Image) error {
	const op = "repository.SaveImageToDB()"

	result, err :=  ir.db.Exec(imageQueries.insert, image.ImagePath, image.Size, image.Name, image.Alt)
	if err != nil {
		ir.l.Warn().Msgf("%s: query exec save image error > %s", op, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		ir.l.Warn().Msgf("%s: query exec save image row affected error > %s", op, err)
	}

	if rows != 1 {
		ir.l.Warn().Msgf("%s: query exec save image number affected rows is > %d", op, rows)
	}

	return nil
}
