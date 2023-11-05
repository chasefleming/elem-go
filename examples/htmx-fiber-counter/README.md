# Counter App Example with `elem-go`, `htmx`, `Go Fiber`

This is a simple counter application demonstrating the integration of a `Go Fiber` backend with `elem-go` for server-side HTML element creation and `htmx` for dynamic front-end interactions.

## Prerequisites

Before you begin, ensure you have Go installed on your system.

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

This command compiles and runs the Go program, starting the Go Fiber server on port `3000`.

## Using the Application

Open your web browser and go to `http://localhost:3000` to view and interact with the counter application.

- Click on "+" to increment the counter.
- Click on "-" to decrement the counter.

These actions will trigger `htmx` requests to the server at the `/increment` and `/decrement` endpoints, respectively, which will update the counter value displayed on the page without needing a full page refresh.

## Code Explanation

The Go application defines three routes:

- `POST /increment`: Increases the counter and returns the new value as a string.
- `POST /decrement`: Decreases the counter and returns the new value as a string.
- `GET /`: Serves the main HTML page which includes:
  - The `htmx` script for handling attributes like `hx-post` and `hx-target`.
  - A styled button to increment and decrement the counter.
  - A `div` element to display the current value of the counter, which is updated by `htmx` when either button is clicked.

The `elem-go` library is used to programmatically create the HTML content served on the root path. The HTML content includes a head element with the required `htmx` script and a body element with inline styling for the counter and buttons.

## Stopping the Server

To stop the application, switch to the terminal window where it's running and press `Ctrl + C`.