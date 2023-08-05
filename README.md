# 🧛grr: Go Reactish Rendering

## Write HTML directly in pure Go functions. Less Typos, More Types!

grr is an open-source library designed to simplify the writing and debugging of Go templates. Its goal is to enhance type safety and reduce the number of typos that can occur when crafting templates. The name grr stands for "Go React-like Rendering".

Drawing from extensive experience with React and TypeScript, as well as a love for the simplicity and type system of Go, grr was created to provide a more ergonomic way to work with the built-in template system of Go. The library is designed to reduce external dependencies as much as possible and leans heavily on the Go standard library.

Key features and benefits of the grr library include:

- **Enhanced Type Safety**: grr helps you reduce typos and make your code more type-safe.
- **No Embeddings Needed**: Everything is part of the binary by default, making deployment super easy without annoying additions.
- **Simplified Debugging**: Debugging Go templates can be a pain, especially when it comes to avoiding nil pointers. grr helps simplify this process.
- **Component-based Testing**: grr allows you to test your templates at the component level.
- **Improved Syntax Highlighting and Autocompletion**: As grr is written directly in Go, you can take advantage of Go's syntax highlighting and autocompletion features in your development environment.

## Installation

To install the grr library, simply run the following command:

`go get github.com/klaus-breyer/grr`

## Usage

The grr library provides a number of functions for working with Go templates. Here are some of the key functions:

- `Render()`: This function allows you to render a template with given data. Here's a basic example:

```go
type Person struct {
     Name string
}

p := Person{Name: "John"}
renderOutput := grr.Render(`<h1>Hello, {{.Name}}!</h1>`, p)
fmt.Println(renderOutput)

// Output: <h1>Hello, John!</h1>``
```

- `Map()`: This function allows you to pass a map of data to a template. It's perfect for rendering arrays, similar to the JavaScript `map` method. Here's a basic example:

```go
persons := []Person{
     {Name: "John"},
     {Name: "Jane"},
}

mapOutput := grr.Render(`<p>{{.Name}}</p>`, persons)
fmt.Println(mapOutput)

// Output: <p>John</p><p>Jane</p>
```

- `Yield()`: This function allows you to render a specific part of a template. It's perfect for rendering several HTML elements in sequence. Here's a basic example:

```go
fmt.Println(grr.Yield(`<body>{{yield}}</body>`, renderOutput, "Family: ", mapOutput))

// Output: <body><h1>Hello, John!</h1>Family: <p>John</p><p>Jane</p></body>
```

Please refer to the examples in the 'example' folder for more detailed usage scenarios. You can find them [here](https://github.com/klausbreyer/grr/tree/main/example).

## Testing

To run tests, simply execute the following command:

bash

`go test`

## Contribution

Contributions from the community are welcome! If you find a bug or want to add a feature, don't hesitate to open an issue or create a pull request.

Please ensure you run the tests before submitting a pull request to ensure your changes don't break existing functionality.
