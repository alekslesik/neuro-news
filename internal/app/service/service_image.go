package service

import (
	"github.com/alekslesik/neuro-news/internal/app/model"
	"github.com/alekslesik/neuro-news/internal/pkg/grabber"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

type ImageService interface {
	InsertImage(*model.Image) error
	GenerateImageKand(*model.Article) (*model.Image, error)
	GenerateImageFruity(*model.Article) (*model.Image, error)
}

type imageService struct {
	ir model.ImageModel
	l  *logger.Logger
	g  *grabber.Grabber
}

// InsertImage insert image to DB
func (is *imageService) InsertImage(i *model.Image) error {
	const op = "service.SaveImageToDB()"

	err := is.ir.InsertImage(i)
	if err != nil {
		is.l.Error().Msgf("%s: save image to DB error > %s", op, err)
		return err
	}
	return nil
}

// GenerateImageKand generate image through Kandinsky API
func (is *imageService) GenerateImageKand(a *model.Article) (*model.Image, error) {
	const op = "service.GenerateImage()"
	// get image model
	image, err := is.g.GetGeneratedImage(a)
	if err != nil {
		is.l.Error().Msgf("%s: get generated image through Kandinsky error > %s", op, err)
		return nil, err
	}
	return image, nil
}

// GenerateImageFruity generate image through FruityBang/neuro-gen
func (is *imageService) GenerateImageFruity(a *model.Article) (*model.Image, error) {
	const op = "service.GenerateImageFruity()"
	// get image model
	image, err := is.g.GenerateImageFruity(a)
	if err != nil {
		is.l.Error().Msgf("%s: get generated image through FruityBang error > %s", op, err)
		return nil, err
	}
	return image, nil
}
