package main

import (
	"html/template"
	"log"

	"github.com/klausbreyer/grr"
)

// GetHTML generates the main HTML structure
func GetHTML() template.HTML {
	// Define main sections
	main := GetMain([]SectionData{
		{Title: "Welcome", Content: "Welcome to our website!"},
		{Title: "About", Content: "We are a company that does amazing things."},
		{Title: "Contact", Content: "Please reach out to us through the contact form."},
	})

	// Define footer data
	footData := DataFoot{Copy: "© 2023"}

	// Generate HTML for each part
	nav := GetNav()
	foot := GetFoot(footData)

	// Assemble parts into final HTML
	return grr.Yield(`
	<html>
		<head>
			<title>grr-example</title>
		</head>
		<body>
		{{yield}}
		</body>
	</html>
	`, nav, main, foot)
}

// Link represents a link in the navigation
type Link struct {
	Label string
	URL   string
}

// GetNav generates the HTML for the navigation
func GetNav() template.HTML {
	// Define navigation data
	navData := []Link{
		{Label: "Home", URL: "/"},
		{Label: "About", URL: "/about"},
		{Label: "Contact", URL: "/contact"},
	}

	return grr.Map(`
    <nav>
		<a href="{{.URL}}">{{.Label}}</a>
    </nav>
    `, navData)
}

// SectionData represents the data for a section
type SectionData struct {
	Title   string
	Content string
}

// GetMain generates the HTML for the main content
func GetMain(data []SectionData) template.HTML {
	return grr.Map(`
    <section>
        <h2>{{.Title}}</h2>
		<p>{{.Content}}</p>
    </section>
    `, data)
}

// DataFoot represents the data for the footer
type DataFoot struct {
	Copy string
}

// GetFoot generates the HTML for the footer
func GetFoot(data DataFoot) template.HTML {
	return grr.Render(`
	<footer>
		{{.Copy}}
	</footer>
	`, data)
}

func main() {
	html := GetHTML()
	log.Println(html)
}
