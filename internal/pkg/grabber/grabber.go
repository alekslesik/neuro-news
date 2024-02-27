package grabber

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/alekslesik/neuro-news/internal/app/model"
	"github.com/alekslesik/neuro-news/pkg/logger"

	"golang.org/x/net/html"
)

// Grabber struct
type Grabber struct {
	log  *logger.Logger
	home string
}

// New return new instance of Grabber struct
func New(log *logger.Logger, home string) *Grabber {
	return &Grabber{log: log, home: home}
}

type Tag struct {
	Name string
	Key  string
	Val  string
}

// GrabArticle grab article from
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
	const op = "grabber.GrabArticle()"

	// Написать код для извлечения списка новостей с сайта

	// Get last article from list
	last, err := g.LastArticle("parts/news/", Tag{
		Name: "ul",
		Key:  "class",
		Val:  "parts-page__body _parts-news",
	})
	if err != nil {
		g.log.Error().Msgf("%s: get last article error > %s", op, err)
		return nil, err
	}

	// Get article title
	title, err := g.TagInner(last, Tag{Name: "span", Key: "class", Val: "topic-body__title"})
	if err != nil {
		g.log.Error().Msgf("%s: get article title error > %s", op, err)
		return nil, err
	}

	// Get article PreviewText
	previewText, err := g.TagInner(last, Tag{Name: "div", Key: "class", Val: "topic-body__title-yandex"})
	if err != nil {
		g.log.Error().Msgf("%s: get article title error > %s", op, err)
		return nil, err
	}

	// Current article time
	articleTime := time.Now()

	// Get article tag
	tag, err := g.TagInner(last, Tag{Name: "a", Key: "class", Val: "topic-header__item topic-header__rubric"})
	if err != nil {
		g.log.Error().Msgf("%s: get article title error > %s", op, err)
		return nil, err
	}

	// Get detail text
	detailText, err := g.DetailText(last, Tag{Name: "p", Key: "class", Val: "topic-body__content-text"})
	if err != nil {
		g.log.Error().Msgf("%s: get article title error > %s", op, err)
		return nil, err
	}

	// Get article href (translit from title)
	href := translit(title)

	// TODO Извлечь Comments
	// Get category (translit from tag)
	category := translit(tag)

	// fill article model
	article := &model.Article{
		Title:       title,
		PreviewText: previewText,
		ArticleTime: articleTime,
		Tag:         tag,
		DetailText:  detailText,
		Href: href,
		Category: category,
	}

	return article, nil
}

// GetLastГКД return last element url from list
func (g *Grabber) LastArticle(url string, tag Tag) (string, error) {
	const op = "grabber.GetLast()"

	// get node from url
	node, err := getNodePage(g.home + url)
	if err != nil {
		g.log.Error().Msgf("%s: get node page error > %s", op, err)
		return "", err
	}

	// recursive find inner of tag with atr
	var result string
	var f func(*html.Node)
	f = func(n *html.Node) {
		// find ul
		if n.Type == html.ElementNode && n.Data == tag.Name {
			for _, attr := range n.Attr {
				// find class=value
				if attr.Key == tag.Key && strings.Contains(attr.Val, tag.Val) {
					for _, attr := range n.FirstChild.FirstChild.Attr {
						if attr.Key == "href" {
							result = attr.Val
							break
						}
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(node)

	return result, nil
}

// TagInner return inner of tag with attribute from url
func (g *Grabber) TagInner(url string, tag Tag) (string, error) {
	const op = "grabber.TagInner()"

	// get node from url
	node, err := getNodePage(g.home + url)
	if err != nil {
		g.log.Error().Msgf("%s: get node page error > %s", op, err)
		return "", err
	}

	// recursive find result of tag with atr
	var result string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tag.Name {
			for _, a := range n.Attr {
				if a.Key == tag.Key && a.Val == tag.Val {
					if d := n.FirstChild.Data; d != "" {
						result = d
					}
					break
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(node)

	return result, nil
}

// DetailText return detail article text
func (g *Grabber) DetailText(url string, tag Tag) (string, error) {
	const op = "grabber.DetailText()"

	var result string
	var resArr []string

	// get node from url
	node, err := getNodePage(g.home + url)
	if err != nil {
		g.log.Error().Msgf("%s: get node page error > %s", op, err)
		return "", err
	}

	// recursive find result of tag with atr
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tag.Name {
			for _, a := range n.Attr {
				if a.Key == tag.Key && a.Val == tag.Val {
					if p := n.FirstChild.Data; p != "" {
						p = "<p>" + p + "</p>"
						resArr = append(resArr, p)
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(node)

	result = strings.Join(resArr, "\n")

	return result, nil
}

// getNodePage return page like a Node
func getNodePage(url string) (*html.Node, error) {
	// response from url
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// read response body
	var body []byte
	if res.StatusCode == http.StatusOK {
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
	}

	// create node from response body
	node, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	return node, nil
}

// getTagHref return tag href with attr from url
func getTagHref(url, tag, atr, atrVal string) (string, error) {
	// get node from url
	node, err := getNodePage(url)
	if err != nil {
		return "", err
	}

	// recursive find inner of tag with atr
	var href string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "ul" {
			for _, a := range n.Attr {
				if a.Key == "class" && strings.Contains(a.Val, "parts-page__body _parts-news") {

					for _, attr := range n.FirstChild.Attr {
						if attr.Key == "href" {
							href = attr.Val
							break
						}
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(node)

	return href, nil
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

// translit
func translit(src string) string {
	var result string
	var resArr []string

	var dictionary = map[string]string{
		"А": "a", "а": "a",
		"Б": "b", "б": "b",
		"В": "v", "в": "v",
		"Г": "g", "г": "g",
		"Д": "d", "д": "d",
		"Е": "e", "е": "e",
		"Ё": "e", "ё": "e",
		"Ж": "zh", "ж": "zh",
		"З": "z", "з": "z",
		"И": "i", "и": "i",
		"Й": "i", "й": "i",
		"К": "k", "к": "k",
		"Л": "l", "л": "l",
		"М": "m", "м": "m",
		"Н": "n", "н": "n",
		"О": "o", "о": "o",
		"П": "p", "п": "p",
		"Р": "r", "р": "r",
		"С": "s", "с": "s",
		"Т": "t", "т": "t",
		"У": "u", "у": "u",
		"Ф": "f", "ф": "f",
		"Х": "h", "х": "h",
		"Ц": "c", "ц": "c",
		"Ч": "ch", "ч": "ch",
		"Ш": "sh", "ш": "sh",
		"Щ": "sh'", "щ": "sh'",
		"Ъ": "", "ъ": "",
		"Ы": "y", "ы": "y",
		"Ь": "", "ь": "",
		"Э": "e", "э": "e",
		"Ю": "yu", "ю": "yu",
		"Я": "ya", "я": "ya",
		" ": "-",
		"0": "0", "1": "1", "2": "2", "3": "3", "4": "4",
		"5": "5", "6": "6", "7": "7", "8": "8", "9": "9",
	}

	split := strings.Split(src, "")

	for _, s := range split {
		resArr = append(resArr, dictionary[s])
	}

	result = strings.Join(resArr, "")

	return result
}
