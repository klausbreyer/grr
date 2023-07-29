package main

import (
	"html/template"
	"log"

	"github.com/klausbreyer/grr"
)

func getHtml() template.HTML {
	navData := DataNav{
		InputVariable: "Input Var 1",
		OtherInput:    "Input Var 2",
	}
	nav := getNav(navData)
	foot := getFoot(DataFoot{Copy: "© 2021"})

	main := getMain([]SectionData{
		{Title: "Section 1"},
		{Title: "Section 2"},
		{Title: "Section 3"},
	})

	return grr.Yield(`
	<html>
		<head>
			<title>all-the-highlights</title>
		</head>
		<body>
		{{yield}}
		</body>
	</html>
	`, nav, main, foot)
}

type DataNav struct {
	InputVariable string
	OtherInput    string
}

func getNav(data DataNav) template.HTML {
	return grr.Render(`
    <nav class="shadow sticky top-0 z-10">
        {{.InputVariable}}
        {{.OtherInput}}
		{{.Foot}}
    </nav>
    `, struct {
		Foot          template.HTML
		InputVariable string
		OtherInput    string
	}{
		getFoot(DataFoot{Copy: "© 2021"}),
		data.InputVariable,
		data.OtherInput,
	})
}

type SectionData struct {
	Title string
}

func getMain(data []SectionData) template.HTML {
	return grr.Map(`
    <section>
        <h2>{{.Title}}</h2>
    </section>
    `, data)
}

type DataFoot struct {
	Copy string
}

func getFoot(data DataFoot) template.HTML {
	return grr.Render(`
	<footer>
		{{.Copy}}
	</footer>
	`, struct {
		Copy string
	}{
		data.Copy,
	})
}

func main() {
	html := getHtml()
	log.Println(html)
}
