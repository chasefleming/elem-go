# `htmx` Integration with `elem-go`

The `htmx` subpackage within `elem` provides a seamless way to integrate the [htmx](https://htmx.org/) library, allowing for easy creation of dynamic web applications. This document outlines how to use the `htmx` subpackage and its features.

## Table of Contents

- [Introduction](#introduction)
- [htmx Version Compatibility](#htmx-version-compatibility)
- [Usage](#usage)
- [Creating Elements with htmx Attributes](#creating-elements-with-htmx-attributes)
- [Supported htmx Attributes](#supported-htmx-attributes)
- [Examples](#examples)

## Introduction

The `htmx` subpackage offers constants and utility functions tailored for htmx-specific attributes. This makes it easier to add dynamic behaviors to your web elements without writing verbose attribute strings.

## htmx Version Compatibility

This package targets **htmx 2.x**. To include htmx in your page, use the CDN script recommended by the [htmx docs](https://htmx.org/docs/#installing):

```go
elem.Script(attrs.Props{
    attrs.Src: "https://cdn.jsdelivr.net/npm/htmx.org@2.0.10/dist/htmx.min.js",
})
```

A few constants from the htmx 1.x era (such as `HXSSE`, `HXWS`, and the bare `HXOn`) are deprecated because the corresponding attributes were removed in htmx 2.0 — their godoc comments point to the 2.x replacements. If you are upgrading an app from htmx 1.x, see the official [htmx migration guide](https://htmx.org/migration-guide-htmx-1/).

## Usage

To utilize the `htmx` subpackage, import it alongside the main `elem` package:

```go
import (
    "github.com/chasefleming/elem-go"
    "github.com/chasefleming/elem-go/attrs"
    "github.com/chasefleming/elem-go/htmx"
)
```

## Creating Elements with htmx Attributes

With the `htmx` subpackage, you can effortlessly add htmx-specific attributes to your elements:

```go
content := elem.Div(attrs.Props{
    attrs.ID:       "container",
    attrs.Class:    "my-class",
    htmx.HXGet:    "/path-to-new-content",
    htmx.HXTarget: "#content-div",
},
    elem.H1(nil, elem.Text("Hello, Elem!")),
    elem.Div(attrs.Props{attrs.ID: "content-div"}, elem.Text("Initial content")),
)
```

In this example, the main `div` has htmx attributes set to fetch content from `/path-to-new-content` and place it inside a `div` with the ID `content-div`.

## Supported `htmx` Attributes

The `htmx` subpackage provides constants for the `htmx` attributes, mirroring the [htmx reference](https://htmx.org/reference/):

- **Core Attributes**
    - `HXGet`: URL for GET requests.
    - `HXPost`: URL for POST requests.
    - `HXPushURL`: Pushes a new URL into the browser history.
    - `HXSelect`: CSS selector for the content to swap in from the response.
    - `HXSelectOOB`: CSS selectors for out-of-band content in the response.
    - `HXSwap`: How the response is swapped in (`innerHTML`, `outerHTML`, `textContent`, etc.).
    - `HXSwapOOB`: Marks response content for out-of-band swapping.
    - `HXTarget`: Where to place the content in the DOM.
    - `HXTrigger`: Event that triggers the request.
    - `HXVals`: Values to submit with the request.

- **Additional Attributes**
    - `HXBoost`: Enhances links and forms with AJAX.
    - `HXConfirm`: Confirmation message before sending the request.
    - `HXDelete`: URL for DELETE requests.
    - `HXDisable`: Disables htmx processing for an element and its children.
    - `HXDisabledElt`: Elements to disable while a request is in flight.
    - `HXDisinherit`: Controls which attributes are not inherited by children.
    - `HXEncoding`: Changes the request encoding (e.g. `multipart/form-data`).
    - `HXExt`: htmx extensions to use.
    - `HXHeaders`: Specifies request headers.
    - `HXHistory`: Prevents snapshotting the page into the history cache.
    - `HXHistoryElt`: Element to snapshot and restore for history navigation.
    - `HXInclude`: Additional elements whose values are included in the request.
    - `HXIndicator`: Element displayed as an indicator while processing.
    - `HXInherit`: Controls which attributes are inherited by children (htmx 2.0+).
    - `HXParams`: Filters the parameters submitted with the request.
    - `HXPatch`: URL for PATCH requests.
    - `HXPreserve`: Keeps an element unchanged between requests.
    - `HXPrompt`: Prompts the user for input before the request.
    - `HXPut`: URL for PUT requests.
    - `HXReplaceURL`: Replaces the current URL in the browser history.
    - `HXRequest`: Configures the request (timeout, credentials, headers).
    - `HXSync`: Synchronizes requests between elements.
    - `HXValidate`: Forces validation before the request.

- **Extension Attributes** (require the [SSE](https://htmx.org/extensions/sse/) / [WebSockets](https://htmx.org/extensions/ws/) extensions)
    - `SSEConnect`, `SSESwap`, `SSEClose`: Server Sent Events.
    - `WSConnect`, `WSSend`: WebSockets.

- **Event Handlers**
    - `HXOnBeforeRequest`, `HXOnAfterSwap`, and the other `HXOn*` constants: inline handlers for [htmx events](https://htmx.org/reference/#events), rendered in the `hx-on--<event>` form.

## Examples

### Loading Content on Button Click

To create a button that loads content into a specific div when clicked:

```go
button := elem.Button(attrs.Props{
    htmx.HXGet:    "/fetch-content",
    htmx.HXTarget: "#result-div",
}, elem.Text("Load Content"))

contentDiv := elem.Div(attrs.Props{attrs.ID: "result-div"}, elem.Text("Initial content"))

pageContent := elem.Div(nil, button, contentDiv)
```

When the button is clicked, `htmx` will fetch content from the `/fetch-content` endpoint and replace the content inside the `#result-div`.

## Handling JSON Strings in Attributes

When using attributes like `hx-vals` that require JSON strings, ensure the JSON is wrapped in single quotes. This is necessary for correct rendering by `htmx`. For example:

```go
content := elem.Div(attrs.Props{
    htmx.HXGet:  "/example",
    htmx.HXVals: `'{"myVal": "My Value"}'`,
}, elem.Text("Get Some HTML, Including A Value in the Request"))
```