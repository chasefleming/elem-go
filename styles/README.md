# `styles` Subpackage in `elem-go`

The `styles` subpackage within `elem-go` offers enhanced functionality for CSS styling in Go-based web applications. This document provides a detailed overview of how to use the `styles` subpackage and its features.

## Table of Contents

- [Introduction](#introduction)
- [Usage](#usage)
- [Styling Elements with `styles.Props`](#styling-elements-with-stylesprops)
- [Features](#features)
    - [`Merge` Function](#merge-function)
    - [`Style` and `CSS` Functions](#style-and-css-functions)

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
styleTag := elem.Style(nil, styles.CSS(cssContent))

// Incorporating the <style> tag in an HTML document
document := elem.Html(nil, elem.Head(nil, styleTag), /* ... */)
```