package template

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"
)

// Change utc time to human tim
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

// Decode html tags
func decodeHTML(s string) template.HTML {
	return template.HTML(s)
}

// Generate HTML for pagination
func generatePaginationHTML(totalPages, currentPage int) template.HTML {
	var builder strings.Builder

	builder.WriteString(`<div class=article-pagination><ul>`)

	for i := 1; i < totalPages; i++ {
		if i == currentPage {
			builder.WriteString(fmt.Sprintf(`<li class="active"><a href="#news" class="active">%d</a></li>`, i))
		} else {
			builder.WriteString(fmt.Sprintf(`<li><a href="/?PAGEN_1=%d#news">%d</a></li>`, i, i))
		}
	}

	// Добавление кнопки "Следующая страница", если это не последняя страница
	// Add button "next page", if it is not the end page
	if currentPage < totalPages {
		builder.WriteString(`<li><a href="?PAGEN_1=` + strconv.Itoa(currentPage+1) + `#news` + `"><i class="fa fa-angle-right"></i></a></li>`)
	}

	builder.WriteString(`</ul></div>`)
	return template.HTML(builder.String())
}

// Initialize a template.FuncMap object and store it in a global variable. This
// essentially a string-keyed map which acts as a lookup between the names of o
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate":              humanDate,
	"decodeHTML":             decodeHTML,
	"generatePaginationHTML": generatePaginationHTML,
}
