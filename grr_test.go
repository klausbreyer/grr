// This test suite tests the functions in the grr package.
// Note that we cannot test negative cases for some of the functions, because they call log.Fatal or panic
// when they encounter an error. This stops the entire program (or in this case, the entire test suite),
// making it impossible to continue with further tests.
// In a production environment, it would be better to modify these functions to return an error,
// which can be tested and handled appropriately.

package grr

import (
	"html/template"
	"testing"
)

// Test data for Struct function
type Person struct {
	Name string
	Age  int
}

// Test for Struct function
func TestStruct(t *testing.T) {
	p := Person{Name: "John", Age: 23}
	Struct(p)
	// This function prints output to stdout, so it's difficult to verify its behavior in a unit test.
	// In a production environment, you might want to modify it to return a string instead.
}

// Test for joinHTML function
func TestJoinHTML(t *testing.T) {
	htmls := []template.HTML{"<p>Hello</p>", "<p>World</p>"}
	sep := ""
	result := joinHTML(htmls, sep)
	expected := template.HTML("<p>Hello</p><p>World</p>")
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

// Test for Render function
func TestRender(t *testing.T) {
	tmpl := "<h1>{{.}}</h1>"
	data := "Hello, World"
	result := Render(tmpl, data)
	expected := template.HTML("<h1>Hello, World</h1>")
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

// Test for Yield function
func TestYield(t *testing.T) {
	tmpl := "<h1>{{yield}}</h1>"
	yield := []template.HTML{"Hello, World"}
	result := Yield(tmpl, yield...)
	expected := template.HTML("<h1>Hello, World</h1>")
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

// Test for Flatten function
func TestFlatten(t *testing.T) {
	items := []template.HTML{"<p>Hello</p>", "<p>World</p>"}
	result := Flatten(items)
	expected := template.HTML("<p>Hello</p><p>World</p>")
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

// Test for Map function
func TestMap(t *testing.T) {
	tmpl := "<h1>{{.}}</h1>"
	items := []string{"Hello", "World"}
	result := Map(tmpl, items)
	expected := template.HTML("<h1>Hello</h1><h1>World</h1>")
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
