package model

import ()

type Image struct {
	ImageID   int
	ImagePath string
	Size      int64
	Name      string
	Alt       string
}

type ImageModel interface {
	SaveImageToDB(*Image) error
}
