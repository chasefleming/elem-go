# Advanced CSS Features with `StyleManager` in `elem-go`

`StyleManager` within the `styles` package, is a powerful addition to elem-go, providing developers with the tools to programmatically manage CSS styles with advanced features like pseudo-classes, animations, and media queries. This enables the creation of dynamic and responsive web applications directly in Go.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
    - [Creating a StyleManager](#creating-a-stylemanager)
    - [Adding Styles](#adding-styles)
    - [Pseudo-Classes](#pseudo-classes)
    - [Pseudo-Elements](#pseudo-elements)
    - [Animations](#animations)
    - [Media Queries](#media-queries)
- [Features](#features)
- [Integration with `elem-go`](#integration-with-elem-go)
- [Examples](#examples)

## Introduction

The `StyleManager` provides a robust solution for managing CSS in Go-based web applications. It simplifies the creation of styles, including complex scenarios involving pseudo-classes animations, and media queries, within the type-safe environment of Go.

## Installation

To leverage `StyleManager`, ensure you're importing the `styles` package along with the core `elem` package:

```go
import (
    "github.com/chasefleming/elem-go"
    "github.com/chasefleming/elem-go/styles"
)
```

## Usage

### Creating a `StyleManager`

Initialize a `StyleManager` using `NewStyleManager()` to start managing your styles:

```go
styleMgr := stylemanager.NewStyleManager()
```

Now, when you render your HTML elements, all you have to do is use `RenderWithOptions` instead of `Render` and provide the `StyleManager` instance:

```go
html := elem.Div(
    nil,
    elem.Text("Hello, World!"),
)

html.RenderWithOptions(elem.RenderOptions{
    StyleManager: styleMgr,
})
```

> Note: This will inject the generated CSS into the HTML output in a `<style>` tag. You cannot manually inject the CSS into your HTML, as the styles may not all exist until the time of rendering.

### Adding Styles

Define styles using `StyleManager` and receive a unique class name for application:

```go
className := styleMgr.AddStyle(styles.Props{
    styles.BackgroundColor: "blue",
    styles.Color:           "white",
})
```

You can then apply the generated class name to your HTML elements:

```go
button := elem.Button(
    attrs.Props{
        attrs.Class: className,
    },
    elem.Text("Click Me"),
)
```

### Pseudo-Classes

Define styles with pseudo-classes for interactive styling:

```go
compositeClassName := styleMgr.AddCompositeStyle(styles.CompositeStyle{
    Default: styles.Props{styles.Color: "pink"},
    PseudoClasses: map[string]styles.Props{
        styles.PseudoHover: {styles.Color: "blue"},
    },
})
```

With the resulting class name, you can apply the styles to your HTML elements:

```go
link := elem.A(
    attrs.Props{
        attrs.Class: compositeClassName,
    },
    elem.Text("Hover over me"),
)
```

### Pseudo-Elements

Define styles with pseudo-elements to insert content or style specific parts of an element:

```go
compositeClassName := styleMgr.AddCompositeStyle(styles.CompositeStyle{
    Default: styles.Props{
		styles.Position: "relative"
	},
    PseudoElements: map[string]styles.Props{
        styles.PseudoBefore: {
			styles.Content: "'Before Text'"
		},
    },
})
```

In this example, the `::before` pseudo-element will insert the text "Before Text" before the content of the element.

### Animations

Create animations and apply them to styles with the returned animation name:

```go
animationName := styleMgr.AddAnimation(styles.Keyframes{
    attrs.From: {styles.Color: "red"},
    attrs.To:   {styles.Color: "blue"},
})
```

Then, apply the animation to your styles using the generated animation name:

```go
styleWithAnimationClassName := styleMgr.AddStyle(styles.Props{
    styles.AnimationName: animationName,
	styles.AnimationDuration: "2s",
})
```

### Media Queries

Apply styles conditionally based on media queries:

```go
classNameWithMediaQuery := styleMgr.AddCompositeStyle(styles.CompositeStyle{
    Default: styles.Props{styles.Padding: "10px"},
    MediaQueries: map[string]styles.Props{
        "@media (min-width: 768px) and (max-width: 1024px)": {"padding": "20px"},
    },
})
```

Just like before, you can apply the generated class name to your HTML elements:

```go
div := elem.Div(
    attrs.Props{
        attrs.Class: classNameWithMediaQuery,
    },
    elem.Text("Responsive Design"),
)
```

## Features

## Why Use `StyleManager`?

- **Style Deduplication**: Automatically deduplicates styles to optimize CSS output, ensuring that your style sheets are as efficient and concise as possible.
- **Automatic Class Name Generation**: Generates unique class names based on style content, abstracting away the need for manually naming classes and reducing the likelihood of naming collisions.
- **Type-Safe Style References**: By utilizing variables for class names, animation names, and other identifiers instead of plain strings, the system enhances code readability and maintainability. This approach reduces errors, such as typos in class or animation names, and improves the developer experience by offering autocomplete and refactoring support in IDEs. It ensures that references to styles, animations, and media queries are checked at compile time, leading to fewer runtime errors and a more robust codebase.

## Integration with `elem-go`

`StyleManager` integrates smoothly with `elem-go`, enabling you to apply generated class names directly to your HTML elements, enhancing the dynamic capabilities of your web applications.

## Examples

For more examples and detailed usage of `StyleManager`, refer to the [`StyleManager` demo application](../examples/stylemanager-demo).