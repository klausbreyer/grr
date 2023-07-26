package grr

import (
	"bytes"
	"html/template"
	"log"
	"reflect"
	"strings"
)

// joinHTML joins template.HTMLs with a separator.
func joinHTML(htmls []template.HTML, sep string) template.HTML {
	s := make([]string, len(htmls))
	for i, h := range htmls {
		s[i] = string(h)
	}
	return template.HTML(strings.Join(s, sep))
}

// Extend converts a struct to a map[string]interface{}.
func Extend(s interface{}) map[string]interface{} {
	if s == nil {
		s = struct{}{}
	}
	v := reflect.ValueOf(s)
	t := v.Type()

	m := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Name
		value := v.Field(i).Interface()
		m[key] = value
	}
	return m
}

// Render returns a template.HTML from a template string and data.
// data interface{} so that it can use a map or a struct.
func Render(tmpl string, data interface{}, yield ...template.HTML) template.HTML {
	t := template.Must(template.New("").Funcs(template.FuncMap{
		"yield": func() template.HTML {
			return joinHTML(yield, "")
		},
	}).Parse(tmpl))

	var buf bytes.Buffer
	err := t.Execute(&buf, data)
	if err != nil {
		log.Fatal(err)
	}
	return template.HTML(buf.String())
}

type DataNav struct {
	InputVariable string
	OtherInput    string
}

func Flatten(items []template.HTML) template.HTML {
	var all template.HTML
	for _, row := range items {
		all += row
	}
	return all
}

func getNav(data DataNav, yield ...template.HTML) template.HTML {
	dataMap := Extend(data)
	dataMap["Foot"] = getFoot(DataFoot{Copy: "© 2021"})
	return Render(`
	<nav class="shadow sticky top-0 z-10">
		{{.Foot}}
		{{.InputVariable}}
		{{.OtherInput}}
		{{yield}}
	</nav>
	`, dataMap, yield...)
}

type DataFoot struct {
	Copy string
}

func getFoot(data DataFoot, yield ...template.HTML) template.HTML {
	return Render(`
	<footer>
		{{.Copy}}
		{{yield}}
	</footer>
	`, Extend(data), yield...)
}

func main() {
	navData := DataNav{
		InputVariable: "Eingabe 1",
		OtherInput:    "Eingabe 2",
	}
	html := getNav(navData, "Yield 1", "Yield 2")
	log.Println(html)
}
