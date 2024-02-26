package service

import "github.com/alekslesik/neuro-news/internal/app/model"

type ImageService interface {
	SaveImageToDB(model.Image) error
}

type imageService struct {
	ir model.ImageModel
}

func (is *imageService) SaveImageToDB(model.Image) error {
	return nil
}
