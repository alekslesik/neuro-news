package grabber

import (
	"github.com/alekslesik/neuro-news/internal/app/model"
	"github.com/alekslesik/neuro-news/pkg/logger"
)

// Grabber struct
type Grabber struct {
	l *logger.Logger
}

// New return new instance of Grabber struct
func New(l *logger.Logger) *Grabber {
	return &Grabber{l: l}
}

// GrabArticle
func (g *Grabber) GrabArticle() (*model.Article, error) {
	// type Article struct {
	// 	ArticleID   int
	// 	Title       string
	// 	PreviewText string
	// 	Image       string
	// 	ArticleTime time.Time
	// 	Tag         string
	// 	DetailText  string
	// 	Href        string
	// 	Comments    int
	// 	Category    string
	// 	Video       string
	// }

	// Написать код для извлечения списка новостей с сайта.
	// Выбрать последнюю новость из списка.
	// Извлечь заголовок Title
	// Извлечь превью текст PreviewText
	// Извлечь время статьи ArticleTime
	// Извлечь тег Tag
	// Извлечь детальный текст DetailText
	// Извлечь ссылку на статью Href (транслитерация заголовка статьи)
	// TODO Извлечь Comments
	// Извлечь категорию Category (транслитерация тега)

	return nil, nil
}

// GetGeneratedImage generate, save image and return image model
func (g *Grabber) GetGeneratedImage(title string) (model.Image, error) {
	// send news title to API and take generated image link

	// download image to website/static/img

	// create and fill image model
	imgModel := model.Image{}

	// return file
	return imgModel, nil
}
