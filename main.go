package main

import (
	"bytes"
	"html/template"
	"log"
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

// render returns a template.HTML from a template string and data.
func render(tmpl string, data interface{}, yield ...template.HTML) template.HTML {
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

func getNav(data DataNav, yield ...template.HTML) template.HTML {
	return render(`
	<nav class="shadow sticky top-0 z-10">

		{{.InputVariable}}
		{{.OtherInput}}
		{{yield}}
	</nav>
	`, data, yield...)
}

type DataFoot struct {
	Copy string
}

func getFoot(data DataFoot, yield ...template.HTML) template.HTML {
	return render(`
	<footer>
		{{.Copy}}
		{{yield}}
	</footer>
	`, data, yield...)
}

func main() {
	nav := DataNav{
		InputVariable: "Eingabe 1",
		OtherInput:    "Eingabe 2",
	}
	foot := DataFoot{
		Copy: "© 2021",
	}
	html := getNav(nav, "Yield 1", getFoot(foot, "Yield 1", "Yield 2"))
	log.Println(html)
}
