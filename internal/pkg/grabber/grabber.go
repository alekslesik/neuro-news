package grabber

import "github.com/alekslesik/neuro-news/pkg/logger"

// Grabber struct
type Grabber struct {
	l *logger.Logger
}

// New return new instance of Grabber struct
func New(l *logger.Logger) *Grabber {
	return &Grabber{l: l}
}

//
// func (g *Grabber)  {

// }