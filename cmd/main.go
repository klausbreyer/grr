package main

import (
	"html/template"
	"log"

	"github.com/klausbreyer/grr"
)

type DataNav struct {
	InputVariable string
	OtherInput    string
}

// my favorite. Because it is the most type save and the most readable and it does no hack and
// i dont like typing "" all the time.
func getNav(data DataNav) template.HTML {
	return grr.Render(struct {
		Foot          template.HTML
		InputVariable string
		OtherInput    string
	}{
		getFoot(DataFoot{Copy: "© 2021"}),
		data.InputVariable,
		data.OtherInput,
	},
		`
    <nav class="shadow sticky top-0 z-10"">
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
	return grr.Render(struct {
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
