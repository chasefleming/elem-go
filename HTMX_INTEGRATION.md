# `htmx` Integration with `elem-go`

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
	"github.com/chasefleming/elem-go/attrs"
    "github.com/chasefleming/elem-go/htmx"
)
```

## Creating Elements with htmx Attributes

With the `htmx` subpackage, you can effortlessly add htmx-specific attributes to your elements:

```go
content := elem.Div(attrs.Props{
    elem.ID:       "container",
    elem.Class:    "my-class",
    htmx.HXGet:    "/path-to-new-content",
    htmx.HXTarget: "#content-div",
},
    elem.H1(nil, elem.Text("Hello, Elem!")),
    elem.Div(attrs.Props{attrs.ID: "content-div"}, elem.Text("Initial content")),
)
```

In this example, the main `div` has htmx attributes set to fetch content from `/path-to-new-content` and place it inside a `div` with the ID `content-div`.

## Supported `htmx` Attributes

The `htmx` subpackage provides constants for commonly used `htmx` attributes:

- **Request Modifiers**
    - `HXGet`: URL for GET requests.
    - `HXPost`: URL for POST requests.
    - `HXPut`: URL for PUT requests.
    - `HXDelete`: URL for DELETE requests.
    - `HXPatch`: URL for PATCH requests.

- **Request Headers and Content-Type**
    - `HXHeaders`: Specifies request headers.
    - `HXContent`: Specifies the content type of the request.

- **Request Parameters**
    - `HXParams`: Parameters to include with the request.
    - `HXValues`: Values to include with the request.

- **Request Timeout and Retries**
    - `HXTimeout`: Timeout for the request.
    - `HXRetry`: Number of times to retry the request.
    - `HXRetryTimeout`: Timeout before retrying the request.

- **Response Processing**
    - `HXSwap`: How to swap the content.
    - `HXTarget`: Where to place the content in the DOM.
    - `HXSwapOOB`: Out-of-band swapping.
    - `HXSelect`: CSS selector for element in returned HTML.
    - `HXExt`: htmx extensions to use.
    - `HXVals`: Values to process in the response.

- **Events**
    - `HXTrigger`: Event that triggers the request.
    - `HXConfirm`: Confirmation message before sending the request.
    - `HXOn`: Event listener on the element.
    - `HXTriggeringElement`: Element that triggered the request.
    - `HXTriggeringEvent`: Event that triggered the request.

- **Indicators**
    - `HXIndicator`: Element displayed as an indicator while processing.

- **History**
    - `HXPushURL`: Pushes a new URL to the browser history.
    - `HXHistoryElt`: Element for history purposes.
    - `HXHistoryAttr`: Attribute for history purposes.

- **Error Handling**
    - `HXBoost`: Enhances links and forms with AJAX.
    - `HXError`: Element for displaying error messages.

- **Caching**
    - `HXCache`: Specifies caching behavior.

## Examples

### Loading Content on Button Click

To create a button that loads content into a specific div when clicked:

```go
button := elem.Button(attrs.Props{
    htmx.HXGet:    "/fetch-content",
    htmx.HXTarget: "#result-div",
}, elem.Text("Load Content"))

contentDiv := elem.Div(attrs.Props{elem.ID: "result-div"}, elem.Text("Initial content"))

pageContent := elem.Div(nil, button, contentDiv)
```

When the button is clicked, `htmx` will fetch content from the `/fetch-content` endpoint and replace the content inside the `#result-div`.