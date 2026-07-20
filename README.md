![elem-go logo](./logo.png)

**Type-safe HTML for Go: no templates, no code generation, no build step.**

`elem-go` lets you build HTML views in pure Go: every element is a Go function and every attribute a typed constant, so typos fail at compile time instead of at runtime in production. And because views are plain Go code rather than a separate template language, the tools you already use work on them unchanged: autocomplete, refactoring, tests, the debugger.

## Features

- Type-safe constructors for nearly every standard HTML5 element and attribute.
- Utilities like `If` and `TransformEach` for conditional and list rendering.
- Automatic HTML escaping of text content (use `Raw` to opt out).
- Zero runtime dependencies.
- Inline CSS styling with the [styles](styles/README.md) subpackage.
- Advanced CSS features (pseudo-classes, animations, media queries) with [`StyleManager`](styles/STYLEMANAGER.md).
- htmx attribute helpers in the [htmx](htmx/README.md) subpackage.

## Installation

To install `elem-go`, use `go get`:

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

Here's an example of creating a `<div>` element with nested `<h1>`, `<h2>`, and `<p>` elements using `elem-go`:

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

Calling the `.Render()` method on this element produces the following HTML:

```html
<div id="container" class="my-class">
    <h1>Hello, Elem!</h1>
    <h2>Subheading</h2>
    <p>This is a paragraph.</p>
</div>
```

### Attributes and Styles

The [`attrs`](attrs/README.md) subpackage provides type-safe attribute functions that ensure you're setting valid attributes for your elements. This helps eliminate potential issues at runtime due to misspelled or unsupported attribute names.

For boolean attributes like `checked` and `selected`, assign the value `"true"` or `"false"`. When set to `"true"`, the attribute renders without an explicit value:

```go
// Using boolean attributes
checkbox := elem.Input(attrs.Props{
    attrs.Type:    "checkbox",
    attrs.Checked: "true",  // This will render as <input type="checkbox" checked>
})
```

For setting styles, the [`styles`](styles/README.md) subpackage enables you to create style objects and convert them to inline CSS strings:

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

The `.Render()` method converts an element tree into an HTML string, ready to serve to a browser.

```go
html := content.Render()
```

> NOTE: When using an <html> element, this method automatically includes a <!DOCTYPE html> preamble in the rendered HTML, ensuring compliance with modern web standards.

#### Custom Rendering Options

For more control over rendering, such as disabling the `<!DOCTYPE html>` preamble, use `RenderWithOptions` with a `RenderOptions` struct:

```go
options := elem.RenderOptions{DisableHtmlPreamble: true}
htmlString := myHtmlElement.RenderWithOptions(options)
```

### Generating Lists of Elements with `TransformEach`

The `TransformEach` function turns a slice of data into a slice of elements:

```go
items := []string{"Item 1", "Item 2", "Item 3"}

liElements := elem.TransformEach(items, func(item string) elem.Node {
    return elem.Li(nil, elem.Text(item))
})

ulElement := elem.Ul(nil, liElements...)
```

In this example, we transformed a slice of strings into a list of `li` elements and then wrapped them in a `ul` element.

### Conditional Rendering with `If`

`elem-go` provides a utility function `If` for conditional rendering of elements.

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

`elem-go` provides a specialized node `None` that implements the `Node` interface but produces no output, for cases where a condition should render nothing.

```go
showWelcomeMessage := false
welcomeMessage := elem.Div(nil, elem.Text("Welcome to our website!"))

content := elem.Div(nil,
    elem.If[elem.Node](showWelcomeMessage, welcomeMessage, elem.None()),
)
```

In this example, `welcomeMessage` is rendered only if `showWelcomeMessage` is `true`. If it's `false`, `None` is rendered instead, which produces no visible output.

Additionally, `None` can be used to create an empty element, as in `elem.Div(nil, elem.None())`, which results in `<div></div>`.

### Supported Elements

`elem-go` provides constructor functions for HTML elements:

- **Document Structure**: `Html`, `Head`, `Body`, `Title`, `Link`, `Meta`, `Style`, `Base`
- **Text Content**: `H1`, `H2`, `H3`, `H4`, `H5`, `H6`, `P`, `Blockquote`, `Pre`, `Code`, `I`, `Br`, `Wbr`, `Hr`, `Small`, `Q`, `Cite`, `Abbr`, `Data`, `Time`, `Var`, `Samp`, `Kbd`
- **Sectioning & Semantic Layout**: `Article`, `Aside`, `FigCaption`, `Figure`, `Footer`, `Header`, `Hgroup`, `Main`, `Mark`, `Nav`, `Search`, `Section`
- **Form Elements**: `Form`, `Input`, `Textarea`, `Button`, `Select`, `Optgroup`, `Option`, `Label`, `Fieldset`, `Legend`, `Datalist`, `Meter`, `Output`, `Progress`
- **Interactive Elements**: `Details`, `Dialog`, `Menu`, `Summary`
- **Grouping Content**: `Div`, `Span`, `Li`, `Ul`, `Ol`, `Dl`, `Dt`, `Dd`
- **Tables**: `Table`, `Caption`, `Col`, `Colgroup`, `Tr`, `Td`, `Th`, `TBody`, `THead`, `TFoot`
- **Hyperlinks and Multimedia**: `Img`, `Picture`, `Map`, `Area`
- **Embedded Content**: `Audio`, `Canvas`, `Embed`, `Iframe`, `Object`, `Source`, `Track`, `Video`
- **Script-supporting Elements**: `Script`, `Noscript`, `Slot`, `Template`
- **Inline Semantic**: `A`, `Strong`, `Em`, `Code`, `I`, `B`, `Bdi`, `Bdo`, `U`, `S`, `Sub`, `Sup`, `Del`, `Dfn`, `Ins`, `Ruby`, `Rt`, `Rp`

### Raw HTML Insertion

The `Raw` function inserts raw HTML verbatim into your document structure:

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

### Script Content Escaping

When you pass content to the `Script` function, `elem-go` automatically escapes the sequences that would otherwise terminate the `<script>` element early (`</script>`, `<script`, and `<!--`, all case-insensitive) by replacing the leading `<` with its `\x3C` escape, as required by the [HTML specification](https://html.spec.whatwg.org/multipage/scripting.html#restrictions-for-contents-of-script-elements). This keeps inline scripts from accidentally breaking out of their element.

```go
script := elem.Script(nil, elem.Raw(`alert("</script>")`))
// Renders: <script>alert("\x3C/script>")</script>
```

### Grouping Elements with Fragment

The `Fragment` function allows you to group multiple elements together without adding an extra wrapper element to the DOM. This is particularly useful when you want to merge multiple nodes into the same parent element without any additional structure.

```go
nodes := []elem.Node{
    elem.P(nil, elem.Text("1")),
    elem.P(nil, elem.Text("2")),
}

content := elem.Div(nil,
    elem.P(nil, elem.Text("0")),
    elem.Fragment(nodes...),
    elem.P(nil, elem.Text("3")),
)
```

Here, the nodes are inserted directly into the parent `div` with no additional wrapper elements in the output.

### Handling JSON Strings and Special Characters in Attributes

When using attributes that require JSON strings or special characters (like quotes), make sure to wrap these strings in single quotes. This prevents the library from adding extra quotes around your value. For example:

```go
content := elem.Div(attrs.Props{
    attrs.ID:    "my-div",
    attrs.Class: "special 'class'",
    attrs.Data:  `'{"key": "value"}'`,
}, elem.Text("Content"))
```

## Advanced CSS Styling with `StyleManager`

For advanced CSS styling, including animations, pseudo-classes, and responsive design via media queries, use `StyleManager` from the `styles` subpackage. It lets you create and manage complex CSS programmatically, with the same type safety as the rest of `elem-go`.

See the [`StyleManager` documentation](styles/STYLEMANAGER.md) for details.

## HTMX Integration

The [htmx subpackage](htmx/README.md) provides typed helpers for htmx attributes, so you can build dynamic server-rendered pages without writing JavaScript. It targets htmx 2.x, with deprecated constants preserved for code written against htmx 1.x.

## Examples

For hands-on sample implementations, see the [`examples/` folder](./examples).

## Tutorials & Guides

Dive deeper into `elem-go` with these tutorials and guides:

- [Building a Counter App with htmx, Go Fiber, and elem-go](https://dev.to/chasefleming/building-a-counter-app-with-htmx-go-fiber-and-elem-go-9jd/)
- [Building a Go Static Site Generator Using elem-go](https://dev.to/chasefleming/building-a-go-static-site-generator-using-elem-go-3fhh)

## Contributing

Contributions are welcome! If you have ideas for improvements or new features, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
