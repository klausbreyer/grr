package grr

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"reflect"
	"runtime"
	"strings"
)

// joinHTML joins multiple template.HTMLs into a single template.HTML string, separated by a provided separator string.
func joinHTML(htmls []template.HTML, sep string) template.HTML {
	s := make([]string, len(htmls))
	for i, h := range htmls {
		s[i] = string(h)
	}
	return template.HTML(strings.Join(s, sep))
}

// Render takes a template string and a data object, and returns a template.HTML after rendering the template with the provided data.
// It also handles any error during the rendering process and logs a fatal error with stack trace if any error occurs.
func Render(tmpl string, data interface{}) template.HTML {
	t := template.Must(template.New("").Funcs(template.FuncMap{}).Parse(tmpl))

	var buf bytes.Buffer
	err := t.Execute(&buf, data)
	if err != nil {
		stack := make([]byte, 1024)
		length := runtime.Stack(stack, true)
		log.Fatal(fmt.Errorf("error rendering template: %w. Template was: %s. Stack trace: \n%s", err, tmpl, stack[:length]))
	}
	return template.HTML(buf.String())
}

// Yield takes a template string and a variable number of template.HTML arguments and returns a rendered template.HTML.
// It provides a "yield" function that can be used within the template to insert the provided arguments.
func Yield(tmpl string, yield ...template.HTML) template.HTML {
	t := template.Must(template.New("").Funcs(template.FuncMap{
		"yield": func() template.HTML {
			return joinHTML(yield, "")
		},
	}).Parse(tmpl))
	var buf bytes.Buffer
	err := t.Execute(&buf, nil)
	if err != nil {
		stack := make([]byte, 1024)
		length := runtime.Stack(stack, true)
		log.Fatal(fmt.Errorf("error rendering template: %w. Template was: %s. Stack trace: \n%s", err, tmpl, stack[:length]))
	}
	return template.HTML(buf.String())
}

// Flatten takes a slice of template.HTML and concatenates them into a single template.HTML.
func Flatten(items []template.HTML) template.HTML {
	var all template.HTML
	for _, row := range items {
		all += row
	}
	return all
}

// Map takes a template string and a slice of any type, and returns a template.HTML after rendering the template with each item of the slice.
// It converts the slice to a slice of interfaces to handle slices of any type, and panics if the provided object is not a slice.
func Map(tmpl string, items interface{}) template.HTML {
	// We need to convert items to []interface{} to iterate over it
	var itemsInterface []interface{}

	// Get value of items (which should be a slice)
	v := reflect.ValueOf(items)

	// Check if items is a slice
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Items is not a slice, but a %s", v.Kind()))
	}

	// Create slice of interfaces with same length
	itemsInterface = make([]interface{}, v.Len())

	// Fill slice of interfaces with values from items
	for i := 0; i < v.Len(); i++ {
		itemsInterface[i] = v.Index(i).Interface()
	}

	var all template.HTML
	for _, row := range itemsInterface {
		all += Render(tmpl, row)
	}
	return all
}
