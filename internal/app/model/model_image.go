package model

import ()

type Image struct {
	ImageID   int
	ImagePath string
	Size      string
	Name      string
	Alt       string
}

type ImageModel interface {
	SaveImageToDB(Image) error
}
