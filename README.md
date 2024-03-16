# elem-go

`elem` is a lightweight Go library for creating HTML elements programmatically. Utilizing the strong typing features of Go, `elem` ensures type safety in defining and manipulating HTML elements, minimizing potential runtime errors. It simplifies the generation of HTML views by providing a simple and intuitive way to create elements and set their attributes, properties, and content.

## Features

- Easily create HTML elements with Go code.
- Type-safe definition and manipulation of elements, attributes, and properties.
- Supports common HTML elements and attributes.
- Utilities for simplified element generation and manipulation.
- Advanced CSS styling capabilities with the [styles](styles/README.md) subpackage.

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

Here's an example of creating a `<div>` element with nested `<h1>`, `<h2>`, and `<p>` elements using elem:

```go
content := elem.Div(attrs.Props{
    attrs.ID:    "container",
    attrs.Class: "my-class",
},
    elem.H1(nil, elem.Text("Hello, Elem!")),
    elem.H2(nil, elem.Text("Subheading")),
    elem.P(nil, elem.Text("This is a paragraph.")),
)
```

When the above Go code is executed and the `.Render()` method is called, it produces the following HTML:

```html
<div id="container" class="my-class">
    <h1>Hello, Elem!</h1>
    <h2>Subheading</h2>
    <p>This is a paragraph.</p>
</div>
```

### Attributes and Styles

The [`attrs`](attrs/README.md) subpackage provides type-safe attribute functions that ensure you're setting valid attributes for your elements. This helps eliminate potential issues at runtime due to misspelled or unsupported attribute names.

For boolean attributes like `checked` and `selected`, you can simply assign them the value `"true"` or `"false"`. When set to `"true"`, the library will correctly render these attributes without needing an explicit value. For instance:

```go
// Using boolean attributes
checkbox := elem.Input(attrs.Props{
    attrs.Type:    "checkbox",
    attrs.Checked: "true",  // This will render as <input type="checkbox" checked>
})
```

For setting styles, the `styles` subpackage enables you to create style objects and convert them to inline CSS strings:

```go
// Define a style
buttonStyle := styles.Props{
    styles.BackgroundColor: "blue",
    styles.Color:           "white",
}

// Convert style to inline CSS and apply it
button := elem.Button(
    attrs.Props{
        attrs.Style: buttonStyle.ToInline(),
    },
    elem.Text("Click Me"),
)
```

See the complete list of supported attributes in [the `attrs` package](./attrs/attrs.go), and for a full overview of style properties and information on using the `styles` subpackage, see the [styles README](styles/README.md).

### Rendering Elements

The `.Render()` method is used to convert the structured Go elements into HTML strings. This method is essential for generating the final HTML output that can be served to a web browser or integrated into templates.

```go
html := content.Render()
```

In this example, `content` refers to an `elem` element structure. When the `.Render()` method is called on content, it generates the HTML representation of the constructed elements.

> NOTE: When using an <html> element, this method automatically includes a <!DOCTYPE html> preamble in the rendered HTML, ensuring compliance with modern web standards.

#### Custom Rendering Options

For more control over the rendering process, such as disabling the HTML preamble, use the `RenderWithOptions` method. This method accepts a `RenderOptions` struct, allowing you to specify various rendering preferences.

```go
options := RenderOptions{DisableHtmlPreamble: true}
htmlString := myHtmlElement.RenderWithOptions(options)
```

This flexibility is particularly useful in scenarios where default rendering behaviors need to be overridden or customized.

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
adminLink := elem.A(attrs.Props{attrs.Href: "/admin"}, elem.Text("Admin Panel"))
guestLink := elem.A(attrs.Props{attrs.Href: "/login"}, elem.Text("Login"))

content := elem.Div(nil,
    elem.H1(nil, elem.Text("Dashboard")),
    elem.If(isAdmin, adminLink, guestLink),
)
```

In this example, if `isAdmin` is `true`, the `Admin Panel` link is rendered. Otherwise, the `Login` link is rendered.

#### `None` in Conditional Rendering

`elem` provides a specialized node `None` that implements the `Node` interface but does not produce any visible output. It's particularly useful in scenarios where rendering nothing for a specific condition is required.

```go
showWelcomeMessage := false
welcomeMessage := elem.Div(nil, elem.Text("Welcome to our website!"))

content := elem.Div(nil,
    elem.If[elem.Node](showWelcomeMessage, welcomeMessage, elem.None()),
)
```

In this example, `welcomeMessage` is rendered only if `showWelcomeMessage` is `true`. If it's `false`, `None` is rendered instead, which produces no visible output.

Additionally, `None` can be used to create an empty element, as in `elem.Div(nil, elem.None())`, which results in `<div></div>`. This can be handy for creating placeholders or structuring your HTML document without adding additional content.

### Supported Elements

`elem` provides utility functions for creating HTML elements:

- **Document Structure**: `Html`, `Head`, `Body`, `Title`, `Link`, `Meta`, `Style`, `Base`
- **Text Content**: `H1`, `H2`, `H3`, `H4`, `H5`, `H6`, `P`, `Blockquote`, `Pre`, `Code`, `I`, `Br`, `Hr`, `Small`, `Q`, `Cite`, `Abbr`, `Data`, `Time`, `Var`, `Samp`, `Kbd`
- **Sectioning & Semantic Layout**: `Article`, `Aside`, `FigCaption`, `Figure`, `Footer`, `Header`, `Hgroup`, `Main`, `Mark`, `Nav`, `Section`
- **Form Elements**: `Form`, `Input`, `Textarea`, `Button`, `Select`, `Optgroup`, `Option`, `Label`, `Fieldset`, `Legend`, `Datalist`, `Meter`, `Output`, `Progress`
- **Interactive Elements**: `Details`, `Dialog`, `Menu`, `Summary`
- **Grouping Content**: `Div`, `Span`, `Li`, `Ul`, `Ol`, `Dl`, `Dt`, `Dd`
- **Tables**: `Table`, `Tr`, `Td`, `Th`, `TBody`, `THead`, `TFoot`
- **Hyperlinks and Multimedia**: `Img`, `Map`, `Area`
- **Embedded Content**: `Audio`, `Iframe`, `Source`, `Video`
- **Script-supporting Elements**: `Script`, `Noscript`
- **Inline Semantic**: `A`, `Strong`, `Em`, `Code`, `I`, `B`, `U`, `Sub`, `Sup`, `Ruby`, `Rt`, `Rp`

### Raw HTML Insertion

The `Raw` function allows for the direct inclusion of raw HTML content within your document structure. This function can be used to insert HTML strings, which will be rendered as part of the final HTML output.

```go
rawHTML := `<div class="custom-html"><p>Custom HTML content</p></div>`
content := elem.Div(nil,
    elem.H1(nil, elem.Text("Welcome to Elem-Go")),
    elem.Raw(rawHTML), // Inserting the raw HTML
    elem.P(nil, elem.Text("More content here...")), 
)

htmlOutput := content.Render()
// Output: <div><h1>Welcome to Elem-Go</h1><div class="custom-html"><p>Custom HTML content</p></div><p>More content here...</p></div>
```
> **NOTE**: If you are passing HTML from an untrusted source, make sure to sanitize it to prevent potential security risks such as Cross-Site Scripting (XSS) attacks.

### HTML Comments

Apart from standard elements, `elem-go` also allows you to insert HTML comments using the `Comment` function:

```go
comment := elem.Comment("Section: Main Content Start")
// Generates: <!-- Section: Main Content Start -->
```

## HTMX Integration

We provide a subpackage for htmx integration. [Read more about htmx integration here](htmx/README.md).

## Examples

For hands-on examples showcasing the usage of `elem`, you can find sample implementations in the `examples/` folder of the repository. Dive into the examples to get a deeper understanding of how to leverage the library in various scenarios.

[Check out the examples here.](./examples)

## Tutorials & Guides

Dive deeper into the capabilities of `elem` and learn best practices through our collection of tutorials and guides:

- [Building a Counter App with htmx, Go Fiber, and elem-go](https://dev.to/chasefleming/building-a-counter-app-with-htmx-go-fiber-and-elem-go-9jd/)
- [Building a Go Static Site Generator Using elem-go](https://dev.to/chasefleming/building-a-go-static-site-generator-using-elem-go-3fhh)

Stay tuned for more tutorials and guides in the future!

## Contributing

Contributions are welcome! If you have ideas for improvements or new features, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
