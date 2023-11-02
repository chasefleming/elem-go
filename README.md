# elem-go

`elem` is a lightweight Go library for creating HTML elements programmatically. Utilizing the strong typing features of Go, `elem` ensures type safety in defining and manipulating HTML elements, minimizing potential runtime errors. It simplifies the generation of HTML views by providing a simple and intuitive way to create elements and set their attributes, properties, and content.

## Features

- Easily create HTML elements with Go code.
- Type-safe definition and manipulation of elements, attributes, and properties.
- Supports common HTML elements and attributes.
- Utilities for simplified element generation and manipulation.

## Installation

To install `elem`, use `go get`:

```bash
go get github.com/chasefleming/elem-go
```

## Usage

Import the `elem` package in your Go code:

```go
import (
    "github.com/chasefleming/elem-go"
    "github.com/chasefleming/elem-go/attrs"
    "github.com/chasefleming/elem-go/styles"
)
```

### Creating Elements

```go
content := elem.Div(elem.Attrs{
    attrs.ID:    "container",
    attrs.Class: "my-class",
},
    elem.H1(nil, elem.Text("Hello, Elem!")),
    elem.H2(nil, elem.Text("Subheading")),
    elem.P(nil, elem.Text("This is a paragraph.")),
)
```

### Attributes

The `attrs` subpackage provides type-safe attribute functions that ensure you're setting valid attributes for your elements. This helps eliminate potential issues at runtime due to misspelled or unsupported attribute names.

For boolean attributes like `checked` and `selected`, you can simply assign them the value `"true"` or `"false"`. When set to `"true"`, the library will correctly render these attributes without needing an explicit value. For instance:

```go
// Using boolean attributes
checkbox := elem.Input(elem.Attrs{
    attrs.Type:    "checkbox",
    attrs.Checked: "true",  // This will render as <input type="checkbox" checked>
})
```

See the complete list of supported attributes in the [`attrs` package](./attrs/attrs.go).

### Rendering Elements

```go
html := content.Render()
```

### Generating Lists of Elements with `TransformEach`

With `elem`, you can easily generate lists of elements from slices of data using the `TransformEach` function. This function abstracts the repetitive task of iterating over a slice and transforming its items into elements.

```go
items := []string{"Item 1", "Item 2", "Item 3"}

liElements := elem.TransformEach(items, func(item string) elem.Node {
    return elem.Li(nil, elem.Text(item))
})

ulElement := elem.Ul(nil, liElements)
```
In this example, we transformed a slice of strings into a list of `li` elements and then wrapped them in a `ul` element.

### Conditional Rendering with `If`

`elem` provides a utility function `If` for conditional rendering of elements.

```go
isAdmin := true
adminLink := elem.A(elem.Attrs{attrs.Href: "/admin"}, elem.Text("Admin Panel"))
guestLink := elem.A(elem.Attrs{attrs.Href: "/login"}, elem.Text("Login"))

content := elem.Div(nil,
    elem.H1(nil, elem.Text("Dashboard")),
    elem.If(isAdmin, adminLink, guestLink),
)
```

In this example, if `isAdmin` is `true`, the `Admin Panel` link is rendered. Otherwise, the `Login` link is rendered.

### Supported Elements

`elem` provides utility functions for creating common HTML elements:

- `A`: Anchor `<a>`
- `Article`: Article `<article>`
- `Aside`: Aside `<aside>`
- `Blockquote`: Blockquote `<blockquote>`
- `Body`: Body `<body>`
- `Br`: Break `<br>`
- `Button`: Button `<button>`
- `Code`: Code `<code>`
- `Div`: Division `<div>`
- `Dl`: Description List `<dl>`
- `Dt`: Description Term `<dt>`
- `Dd`: Description Description `<dd>`
- `Em`: Emphasis `<em>`
- `Footer`: Footer `<footer>`
- `Form`: Form `<form>`
- `H1`: Heading 1 `<h1>`
- `H2`: Heading 2 `<h2>`
- `H3`: Heading 3 `<h3>`
- `Head`: Head `<head>`
- `Header`: Header `<header>`
- `Html`: HTML `<html>`
- `Hr`: Horizontal Rule `<hr>`
- `I`: Idiomatic Text `<i>`
- `Img`: Image `<img>`
- `Input`: Input `<input>`
- `Label`: Label `<label>`
- `Li`: List Item `<li>`
- `Link`: Link `<link>`
- `Main`: Main `<main>`
- `Meta`: Meta `<meta>`
- `Nav`: Navigation `<nav>`
- `Ol`: Ordered List `<ol>`
- `Option`: Dropdown option `<option>`
- `P`: Paragraph `<p>`
- `Pre`: Preformatted Text `<pre>`
- `Script`: Script `<script>`
- `Section`: Section `<section>`
- `Select`: Dropdown select `<select>`
- `Span`: Span `<span>`
- `Strong`: Strong `<strong>`
- `Table`: Table `<table>`
- `TBody`: Table Body `<tbody>`
- `Td`: Table Cell `<td>`
- `Textarea`: Textarea `<textarea>`
- `TFoot`: Table Footer `<tfoot>`
- `Th`: Table Header Cell `<th>`
- `THead`: Table Header `<thead>`
- `Title`: Title `<title>`
- `Tr`: Table Row `<tr>`
- `Ul`: Unordered List `<ul>`

### Setting Styles with the `styles` Subpackage

With the `elem` library, you can also programmatically define and apply CSS styles to your HTML elements using the `styles` subpackage. This approach leverages Go's type system to ensure that your style property names and values are correctly defined.

```go
// Define a style
buttonStyle := elem.Style{
    styles.BackgroundColor: "blue",
    styles.Color:           "white",
}

// Apply the style to an element as an attribute value
button := elem.Button(
    elem.Attrs{
        attrs.Style: elem.ApplyStyle(buttonStyle),
    },
    elem.Text("Click Me"),
)
```

See the complete list of supported properties in the [`styles` package](./styles/styles.go).

## HTMX Integration

We provide a subpackage for htmx integration. [Read more about htmx integration here](HTMX_INTEGRATION.md).

## Examples

For hands-on examples showcasing the usage of `elem`, you can find sample implementations in the `examples/` folder of the repository. Dive into the examples to get a deeper understanding of how to leverage the library in various scenarios.

[Check out the examples here.](./examples)

## Tutorials & Guides

Dive deeper into the capabilities of `elem` and learn best practices through our collection of tutorials and guides:

- [Building a Counter App with htmx, Go Fiber, and elem-go](https://dev.to/chasefleming/building-a-counter-app-with-htmx-go-fiber-and-elem-go-9jd/)

Stay tuned for more tutorials and guides in the future!

## Contributing

Contributions are welcome! If you have ideas for improvements or new features, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.