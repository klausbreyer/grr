package grr

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"runtime"
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

// Render returns a template.HTML from a template string and data.
// data interface{} so that it can use a map or a struct.
func Render(data interface{}, tmpl string) template.HTML {
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

func Flatten(items []template.HTML) template.HTML {
	var all template.HTML
	for _, row := range items {
		all += row
	}
	return all
}

type DataNav struct {
	InputVariable string
	OtherInput    string
}

// my favorite. Because it is the most type save and the most readable and it does no hack and
// i dont like typing "" all the time.
func getNav(data DataNav) template.HTML {
	return Render(struct {
		Foot          template.HTML
		InputVariable string
		OtherInput    string
	}{
		getFoot(DataFoot{Copy: "© 2021"}),
		data.InputVariable,
		data.OtherInput,
	},
		`
    <nav class="shadow sticky top-0 z-10">
        {{.InputVariable}}
        {{.OtherInput}}
		{{.Foot}}
    </nav>
    `)
}

type DataFoot struct {
	Copy string
}

func getFoot(data DataFoot) template.HTML {
	return Render(struct {
		Copy string
	}{
		data.Copy,
	}, `
	<footer>
		{{.Copy}}
	</footer>
	`)
}

func main() {
	navData := DataNav{
		InputVariable: "Eingabe 1",
		OtherInput:    "Eingabe 2",
	}
	html := getNav(navData)
	log.Println(html)
}
