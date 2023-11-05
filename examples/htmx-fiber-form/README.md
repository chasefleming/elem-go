# Simple Form App with `elem-go`, `htmx`, and `Go Fiber`

This simple form application demonstrates the power of `elem-go` for HTML generation, `htmx` for dynamic UIs, and the `Go Fiber` web framework for handling server-side logic in Go.

## Overview

The application presents a basic form to the user for inputting their name and email. Upon submission, it uses `htmx` to send the data to the server without a full page reload and displays the captured data back to the user.

## Installation

Clone or download the repository, then run the following commands to download the necessary packages:

```bash
go mod tidy
```

This will install `elem-go` along with `Go Fiber` and the `htmx` subpackage required to run the application.

## Running the App

Navigate to the directory containing the application's source code and execute the following command:

```bash
go run main.go
```

This command compiles and runs the Go program, starting the `Go Fiber` server on port `3000`.

## Usage

Open your web browser and navigate to `http://localhost:3000`. You will see a form requesting your name and email. Fill in the details and click submit to see the captured data displayed without reloading the page, thanks to `htmx`.

## Code Explanation

- The Go Fiber app is initialized and set to listen on port 3000.
- Two routes are defined: `/` for serving the form and `/submit-form` for processing the form submission.
- When the form is submitted, the `/submit-form` route captures the input from the `name` and `email` fields and sends a response back.
- `htmx` is used to handle the form submission asynchronously, allowing the server response to update the content on the page without a full refresh.
- `elem-go` library is used to programmatically build the HTML content served when the root route is accessed.

## Stopping the Server

To stop the application, use `Ctrl + C` in the terminal.