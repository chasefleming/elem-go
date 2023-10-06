# htmx Integration with `elem`

The `htmx` subpackage within `elem` provides a seamless way to integrate the [htmx](https://htmx.org/) library, allowing for easy creation of dynamic web applications. This document outlines how to use the `htmx` subpackage and its features.

## Table of Contents

- [Introduction](#introduction)
- [Usage](#usage)
- [Creating Elements with htmx Attributes](#creating-elements-with-htmx-attributes)
- [Supported htmx Attributes](#supported-htmx-attributes)
- [Examples](#examples)

## Introduction

The `htmx` subpackage offers constants and utility functions tailored for htmx-specific attributes. This makes it easier to add dynamic behaviors to your web elements without writing verbose attribute strings.

## Usage

To utilize the `htmx` subpackage, import it alongside the main `elem` package:

```go
import (
    "github.com/chasefleming/elem-go"
    "github.com/chasefleming/elem-go/htmx"
)
```

## Creating Elements with htmx Attributes

With the `htmx` subpackage, you can effortlessly add htmx-specific attributes to your elements:

```go
content := elem.Div(elem.Props{
    elem.ID:       "container",
    elem.Class:    "my-class",
    htmx.HXGet:    "/path-to-new-content",
    htmx.HXTarget: "#content-div",
},
    elem.H1(nil, elem.Text("Hello, Elem!")),
    elem.Div(elem.Props{elem.ID: "content-div"}, elem.Text("Initial content")),
)
```

In this example, the main `div` has htmx attributes set to fetch content from `/path-to-new-content` and place it inside a `div` with the ID `content-div`.

## Supported htmx Attributes

The `htmx` subpackage provides constants for commonly used `htmx` attributes:

- **HXGet**: Represents the `hx-get` attribute.
- **HXPost**: Represents the `hx-post` attribute.
- **HXPut**: Represents the `hx-put` attribute.
- **HXDelete**: Represents the `hx-delete` attribute.
- **HXPatch**: Represents the `hx-patch` attribute.
- **HXParams**: Represents the `hx-params` attribute.
- **HXValues**: Represents the `hx-values` attribute.
- **HXSwap**: Represents the `hx-swap` attribute.
- **HXTarget**: Represents the `hx-target` attribute.
- **HXSwapOOB**: Represents the `hx-swap-oob` attribute.
- **HXTrigger**: Represents the `hx-trigger` attribute.
- **HXConfirm**: Represents the `hx-confirm` attribute.
- **HXIndicator**: Represents the `hx-indicator` attribute.
- **HXPushURL**: Represents the `hx-push-url` attribute.
- **HXBoost**: Represents the `hx-boost` attribute.
- **HXError**: Represents the `hx-error` attribute.

## Examples

### Loading Content on Button Click

To create a button that loads content into a specific div when clicked:

```go
button := elem.Button(elem.Props{
    htmx.HXGet:    "/fetch-content",
    htmx.HXTarget: "#result-div",
}, elem.Text("Load Content"))

contentDiv := elem.Div(elem.Props{elem.ID: "result-div"}, elem.Text("Initial content"))

pageContent := elem.Div(nil, button, contentDiv)
```

When the button is clicked, htmx will fetch content from the `/fetch-content` endpoint and replace the content inside the `#result-div`.