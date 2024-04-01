package template

import (
	"html/template"
	"strings"
	"time"
)

// change utc time to human tim
func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	// create Location object for Moscow time
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return ""
	}

	t = t.In(loc)

	monthNames := map[string]string{
		"Jan": "Янв",
		"Feb": "Фев",
		"Mar": "Мар",
		"Apr": "Апр",
		"May": "Май",
		"Jun": "Июн",
		"Jul": "Июл",
		"Aug": "Авг",
		"Sep": "Сен",
		"Oct": "Окт",
		"Nov": "Ноя",
		"Dec": "Дек",
	}

	formatted := t.Format("02 Jan 2006 в 15:04")

	for eng, rus := range monthNames {
		formatted = strings.Replace(formatted, eng, rus, 1)
	}

	return formatted
}

// decode html tags
func decodeHTML(s string) template.HTML {
	return template.HTML(s)
}

// Initialize a template.FuncMap object and store it in a global variable. This
// essentially a string-keyed map which acts as a lookup between the names of o
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate":  humanDate,
	"decodeHTML": decodeHTML,
}
