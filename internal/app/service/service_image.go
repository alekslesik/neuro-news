package service

import (
	"github.com/alekslesik/neuro-news/internal/app/model"
	"github.com/alekslesik/neuro-news/internal/pkg/grabber"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

type ImageService interface {
	SaveImageToDB(*model.Image) error
	GenerateImage(*model.Article) (*model.Image, error)
}

type imageService struct {
	ir model.ImageModel
	l  *logger.Logger
	g  *grabber.Grabber
}

func (is *imageService) SaveImageToDB(i *model.Image) error {
	const op = "service.SaveImageToDB()"

	err := is.ir.SaveImageToDB(i)
	if err != nil {
		is.l.Error().Msgf("%s: save image to DB error > %s", op, err)
		return err
	}
	return nil
}

func (is *imageService) GenerateImage(a *model.Article) (*model.Image, error) {
	const op = "service.GenerateImage()"
	// get image model
	image, err := is.g.GetGeneratedImage(a)
	if err != nil {
		is.l.Error().Msgf("%s: get generated image error > %s", op, err)
		return nil, err
	}
	return image, nil
}
