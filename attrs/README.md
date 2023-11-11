# `attrs` Subpackage in `elem-go`

The `attrs` subpackage within `elem-go` offers a comprehensive set of constants representing HTML attributes, enhancing the process of setting attributes for HTML elements in a type-safe manner. This document outlines the usage and features of the `attrs` subpackage.

## Table of Contents

- [Introduction](#introduction)
- [Usage](#usage)
- [Available HTML Attributes](#available-html-attributes)
- [Using `Props` Type](#using-props-type)
- [Examples](#examples)

## Introduction

The `attrs` subpackage is designed to simplify the process of defining HTML attributes in Go. By providing constants for common attributes, it helps avoid errors due to typos and enhances code readability.

## Usage

To use the `attrs` subpackage, import it alongside the main `elem` package:

```go
import (
    "github.com/chasefleming/elem-go/attrs"
)
```

## Available HTML Attributes

The subpackage includes a wide range of constants representing universal attributes, link/script attributes, meta attributes, image/embed attributes, semantic text attributes, form/input attributes, interactive attributes, miscellaneous attributes, table attributes, iframe attributes, audio/video attributes, and video-specific attributes.

For instance, some of the constants are:

- Universal Attributes like `Class`, `ID`, `Style`
- Link/Script Attributes such as `Href`, `Src`
- Form/Input Attributes including `Type`, `Value`, `Placeholder`

For a full list of available constants, see the [attrs.go file](attrs.go).

## Using `Props` Type

The `Props` type is a map of strings that can be used to pass attribute values to HTML elements in a structured way. This type-safe approach ensures the correct assignment of attributes to elements.

## Examples

Here's an example of using `attrs` constants to set attributes for a button:

```go
buttonAttrs := attrs.Props{
    attrs.Type:  "button",
    attrs.Class: "btn btn-primary",
    attrs.ID:    "submitBtn",
}

button := elem.Button(buttonAttrs, elem.Text("Submit"))
```

In this example, attributes for the button element are defined using the attrs.Props map with attrs constants.