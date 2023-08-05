
# 🧛grr: Go Reactish Rendering

## Embed HTML directly in pure Go functions. Fewer Typos, More Types!

grr is an open-source library constructed with the aim of streamlining the creation and debugging of Go templates. By allowing HTML embedding directly within pure Go functions, grr enhances type safety and minimizes typos that can surface when crafting templates. grr is an acronym that represents "Go React-like Rendering".

With a foundation rooted in extensive experience with React and TypeScript, and an appreciation for Go's simplicity and type system, grr was born out of a desire to make working with Go's built-in template system more ergonomic. Aiming to limit external dependencies, grr is heavily reliant on the Go standard library.

Key attributes and advantages of the grr library include:

- **Embed HTML Directly in Pure Go Functions**: grr enables HTML templates to be written directly within your Go functions, leading to cleaner and more manageable code.
- **Amplified Type Safety**: grr assists in reducing typos and enhancing the type safety of your code.
- **No Need for Embeddings**: Everything is inherently part of the binary, ensuring deployment is a breeze without the need for annoying additions.
- **Simplified Debugging**: grr alleviates the often painful process of debugging Go templates, especially when trying to avoid nil pointers.
- **Component-based Testing**: grr facilitates the testing of your templates at the component level.
- **Enhanced Syntax Highlighting and Autocompletion**: grr's direct Go integration allows you to leverage Go's syntax highlighting and autocompletion features in your development environment.

## Installation

To install the grr library, execute the following command:

```shell
go get github.com/klaus-breyer/grr
```

## Usage

grr provides a suite of functions for enhancing the interaction with Go templates. Key functions include:

### Render

The `Render()` function facilitates the rendering of a template with provided data. Here's a straightforward example:

```go
type Person struct {
     Name string
}

p := Person{Name: "John"}
renderOutput := grr.Render(`<h1>Hello, {{.Name}}!</h1>`, p)
fmt.Println(renderOutput)

// Output: <h1>Hello, John!</h1>
```

An analogous outcome can be achieved with an anonymous struct, providing additional flexibility when an appropriate type doesn't already exist:

```go
renderOutput := grr.Render(`<h1>Hello, {{.Name}}!</h1>`,
    struct {
        Name string
    }{
        Name: "John"
    }
)
fmt.Println(renderOutput)

// Output: <h1>Hello, John!</h1>
```

Here's another scenario where grr can simplify your code. Consider the following Go template code that sets a background color based on a condition:

```go
{{ if eq $id 3}}bg-cyan-600{{else}}bg-white-900{{end}}
```

With grr, you can transfer this condition from the template to your Go code:

```go
bg := "bg-white-900"
if id == 3 {
     bg = "bg-cyan-600"
}
```

Then, within your template, you can simply use `{{.BgColor}}` to set the background color:

```go
renderOutput := grr.Render(`<div class="{{.BgColor}}">Hello, world!</div>`, struct{ BgColor string }{BgColor: bg})
fmt.Println(renderOutput)

// Output: <div class="bg-cyan-600">Hello, world!</div> or <div class="bg-white-900">Hello, world!</div>
```

By adopting this approach, your HTML templates become cleaner and easier to comprehend, and you can sidestep potential errors that may occur when writing complex conditions directly in your templates.

### Map

The `Map()` function enables you to pass a map of data to a template. It's perfect for rendering arrays, akin to the JavaScript `map` method. Here's a simple example:

```go
persons := []Person{
     {Name: "John"},
     {Name: "Jane"},
}

mapOutput := grr.Render(`<p>{{.Name}}</p>`, persons)
fmt.Println(mapOutput)

// Output: <p>John</p><p>Jane</p>
```

### Yield

The `Yield()` function empowers you to render a specific part of a template. It's ideal for rendering several HTML elements in sequence. Here's a basic example:

```go
fmt.Println(grr.Yield(`<body>{{yield}}</body>`, renderOutput, "Family: ", mapOutput))

// Output: <body><h1>Hello, John!</h1><p>John</p><p>Jane</p></body>
```

For a deeper dive into usage scenarios, please refer to the examples in the 'example' folder. You can find them [here](https://github.com/klausbreyer/grr/tree/main/example).

## Testing

To run tests, execute the following command:

```shell
go test
```

## Contribution

Community contributions are warmly welcomed! If you encounter a bug or wish to add a feature, don't hesitate to open an issue or create a pull request.

Before submitting a pull request, please ensure you run the tests to confirm your changes don't disrupt existing functionality.
