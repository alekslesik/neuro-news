package template

import (
	"html/template"
	"time"
)

// HumanDate return nicely formatted string of time.Time object
func HumanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	// Convert the time to UTC before formatting it.
	return t.UTC().Format("02 Янв 2006 в 15:04")
}

// Initialize a template.FuncMap object and store it in a global variable. This
// essentially a string-keyed map which acts as a lookup between the names of o
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate": HumanDate,
}