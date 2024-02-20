package model

import ()

type Image struct {
	ImageId   int
	ImagePath string
	Size      string
	Name      string
	Alt       string
}

type ImageModel interface {
	SaveImageToDB(Image) error
}
