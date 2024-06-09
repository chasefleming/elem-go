# `styles` Subpackage in `elem-go`

The `styles` subpackage within `elem-go` offers enhanced functionality for CSS styling in Go-based web applications. This document provides a detailed overview of how to use the `styles` subpackage and its features.

## Table of Contents

- [Introduction](#introduction)
- [Usage](#usage)
- [Styling Elements with `styles.Props`](#styling-elements-with-stylesprops)
- [Features](#features)
    - [`Style` and `CSS` Functions](#style-and-css-functions)
    - [`Merge` Function](#merge-function)
    - [Type-Safe CSS Values](#type-safe-css-values)
- [Advanced Styling with `StyleManager`](#advanced-styling-with-stylemanager)
    - [Key Features of `StyleManager`](#key-features-of-stylemanager)
    - [Example: Implementing a Hover Effect](#example-implementing-a-hover-effect)
    - [Detailed Usage](stylemanager/README.md)

## Introduction

The `styles` subpackage provides a convenient way to define and manipulate CSS styles in Go. With utility functions and constants tailored for styling, it simplifies the process of applying styles to HTML elements.

## Usage

To use the `styles` subpackage, import it alongside the main `elem` package:

```go
import (
    "github.com/chasefleming/elem-go"
    "github.com/chasefleming/elem-go/styles"
)
```

## Styling Elements with `styles.Props`

The `styles.Props` type allows for defining CSS properties in a structured, type-safe manner. It ensures that your style definitions are well-organized and easy to maintain.

### CSS Property Keys as Constants

To further enhance type safety and reduce errors, the styles subpackage provides constants for CSS property keys. This means you don't have to rely on writing raw string literals for CSS properties, which are prone to typos and errors. Instead, you can use predefined constants that the package offers, ensuring correctness and saving time on debugging.

```go
// Example of using constants for CSS properties
buttonStyle := styles.Props{
    styles.BackgroundColor: "#4CAF50", // Using constant instead of raw string
    styles.Border:          "none",
    // ... other properties
}
```

By using these constants, you can write more reliable and error-resistant style code in Go, making your development process smoother and more efficient.

For a full list of available constants, see the [styles.go file](styles.go).

### Applying Styles to Elements

Once you have defined your styles using `styles.Props`, you can convert them to an inline CSS string using the `ToInline` method. This inline CSS can then be applied directly to HTML elements.

```go
// Define styles using styles.Props
buttonStyle := styles.Props{
    styles.BackgroundColor: "#4CAF50",
    styles.Border:          "none",
    // ... other properties
}

// Convert styles to inline CSS
inlineStyle := buttonStyle.ToInline()

// Apply inline CSS to a button element
button := elem.Button(attrs.Props{attrs.Style: inlineStyle}, elem.Text("Click Me"))
```

In this example, `buttonStyle` is first defined using `styles.Props` and then converted into an inline CSS string using `ToInline`. This string is used to set the style attribute of a button element.

## Features

### `Merge` Function

The `Merge` method combines multiple style prop objects into one. It's useful for applying conditional styles or layering style sets.

```go
// Example style definitions
baseButtonStyle := styles.Props{
  Padding: "10px 15px",
  Border: "none",
  FontWeight: "bold",
}

primaryStyles := styles.Props{
  BackgroundColor: "blue",
  Color: "white",
}

secondaryStyles := styles.Props{
  BackgroundColor: "red",
  Color: "white",
}

// Merging styles with the new Merge function
primaryButtonStyles := styles.Merge(baseButtonStyle, primaryStyles)
secondaryButtonStyles := styles.Merge(baseButtonStyle, secondaryStyles)
```

In the `Merge` function, later style objects take precedence over earlier ones for properties defined in multiple style objects.

### `Style` and `CSS` Functions

These functions facilitate the embedding of CSS into HTML documents, particularly useful for creating <style> tags and including raw CSS.

```go
// CSS content
cssContent := `/* ... */`

// Creating a <style> tag
styleTag := elem.Style(nil, elem.CSS(cssContent))

// Incorporating the <style> tag in an HTML document
document := elem.Html(nil, elem.Head(nil, styleTag), /* ... */)
```

### Type-Safe CSS Values

This suite of functions provides a type-safe approach to defining CSS values, ensuring that only valid CSS values are used.

#### Length and Size Functions

##### `Em(value float64) string` and `Rem(value float64) string`

These functions return a string representation of the given value with the "em" and "rem" units, respectively.

```go
emValue := styles.Em(2.5) // Returns "2.5em"
remValue := styles.Rem(1.5) // Returns "1.5rem"
```

##### `Pixels(value int) string`

This function returns a string representation of the given value with the "px" unit.

```go
pxValue := styles.Pixels(10) // Returns "10px"
```

##### `Percent(value int) string`

This function returns a string representation of the given value with the "%" unit.

```go
percentValue := styles.Percent(50) // Returns "50%"
```

#### Viewport Functions

##### `ViewportHeight(value int) string` and `ViewportWidth(value int) string`

These functions return a string representation of the given value with the "vh" and "vw" units, respectively.

```go
vhValue := styles.ViewportHeight(50) // Returns "50vh"
vwValue := styles.ViewportWidth(25) // Returns "25vw"
```

##### `ViewportMin(value int) string` and `ViewportMax(value int) string`

These functions return a string representation of the given value with the "vmin" and "vmax" units, respectively.

```go
vminValue := styles.ViewportMin(10) // Returns "10vmin"
vmaxValue := styles.ViewportMax(20) // Returns "20vmax"
```

#### Color Functions

##### `RGB(r, g, b int) string`

This function returns a string representation of the given RGB color.

```go
rgbColor := styles.RGB(255, 0, 0) // Returns "rgb(255, 0, 0)"
```

##### `RGBA(r, g, b int, a float64) string`

This function returns a string representation of the given RGBA color.

```go
rgbaColor := styles.RGBA(255, 0, 0, 0.5) // Returns "rgba(255, 0, 0, 0.5)"
```

##### `HSL(h, s, l int) string`

This function returns a string representation of the given HSL color.

```go
hslColor := styles.HSL(120, 100, 50) // Returns "hsl(120, 100%, 50%)"
```

##### `HSLA(h, s, l int, a float64) string`

This function returns a string representation of the given HSLA color.

```go
hslaColor := styles.HSLA(120, 100, 50, 0.5) // Returns "hsla(120, 100%, 50%, 0.5)"
```

#### Time Duration Functions

##### `Seconds(value float64) string`

This function returns a string representation of the given time duration in seconds.

```go
secondsValue := styles.Seconds(2.5) // Returns "2.5s"
```

##### `Milliseconds(value int) string`

This function returns a string representation of the given time duration in milliseconds.

```go
millisecondsValue := styles.Milliseconds(500) // Returns "500ms"
```

#### Other Functions

##### `Int(value int) string`

This function returns a string representation of the given integer value.

```go
intValue := styles.Int(100) // Returns "100"
```

##### `Float(value float64) string`

This function returns a string representation of the given float value.

```go
floatValue := styles.Float(3.14) // Returns "3.14"
```

##### `URL(url string) string`

This function returns a string representation as a formatted CSS URL.

```go
urlValue := styles.URL("https://example.com/image.jpg") // Returns "url('https://example.com/image.jpg')"
```

##### `Var(name string) string`

This function returns a string representation as a CSS variable.

```go
varValue := styles.Var("primary-color") // Returns "var(--primary-color)"
```

## Advanced Styling with `StyleManager`

`StyleManager`, a component of the `styles` package, extends the capability of Go-based web application development by introducing a structured and type-safe approach to managing CSS styles. This integration supports dynamic styling features like pseudo-classes, animations, and responsive design through a Go-centric API, providing a novel way to apply CSS with the added benefits of Go's type system.

### Key Features of `StyleManager`

- **Pseudo-Classes & Animations**: Enables the application of CSS pseudo-classes and keyframe animations to elements for interactive and dynamic styling, leveraging Go's type safety for style definitions.
- **Media Queries**: Supports defining responsive styles that adapt to various screen sizes and orientations, crafted within Go's type-safe environment to enhance web application usability across devices.
- **Automatic Class Name Generation & Style Deduplication**: Improves stylesheet efficiency by automatically generating unique class names for styles and deduplicating CSS rules, all within a type-safe framework.

### Example: Implementing a Hover Effect

The following example showcases how to leverage `StyleManager` to add a hover effect to a button, illustrating the application of dynamic styles within a type-safe context:

```go
// Initialize StyleManager
styleMgr := styles.NewStyleManager()

// Define styles with a hover effect, utilizing Go's type safety
buttonClass := styleMgr.AddCompositeStyle(styles.CompositeStyle{
    Default: styles.Props{
        styles.BackgroundColor: "green",
        styles.Color:           "white",
        styles.Padding:         "10px 20px",
        styles.Border:          "none",
        styles.Cursor:          "pointer",
    },
    PseudoClasses: map[string]styles.Props{
        styles.PseudoHover: {
            styles.BackgroundColor: "darkgreen",
        },
    },
})

// Create a button and apply the generated class name
button := elem.Button(
    attrs.Props{attrs.Class: buttonClass},
    elem.Text("Hover Over Me"),
)

// Use RenderWithOptions to apply the style definitions effectively
htmlOutput := button.RenderWithOptions(elem.RenderOptions{StyleManager: styleMgr})
```

This example demonstrates the use of `StyleManager` for defining and applying a set of styles that include a dynamic hover effect, all within a type-safe framework. Utilizing `RenderWithOptions` is crucial for integrating the styles managed by `StyleManager` into the HTML output.

### Detailed Usage

For a comprehensive guide on using `StyleManager` and its advanced features, refer to the [`StyleManager` documentation](STYLEMANAGER.md).
