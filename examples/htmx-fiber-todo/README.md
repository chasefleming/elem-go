# Todo List Application with `elem-go`, `htmx`, and `Go Fiber`

This project is a Todo List web application built with Go using the `Go Fiber` framework. It features server-side rendering of HTML with `elem-go` and uses `htmx` to handle asynchronous actions without needing to reload the page.

## Prerequisites

Before you begin, ensure you have Go installed on your system.

## Installation

Clone or download the repository, then run the following commands to download the necessary packages:

```bash
go mod tidy
```

This will install elem-go along with Go Fiber and the `htmx` subpackage required to run the application.

## Running the App

Navigate to the directory containing the application's source code and execute the following command:

```bash
go run main.go
```

This command compiles and runs the Go program, starting the `Go Fiber` server on port `3000`.

## Features

- **Add Todos**: Enter a new task in the text box and click "Add" or press enter to add a new Todo to the list.
- **Toggle Todos**: Click on a Todo item to toggle its Done status. Items marked as done will have a line-through style.

## Structure

- `GET /`: Displays the list of Todo items.
- `POST /toggle/:id`: Toggles the Done status of a Todo based on its ID.
- `POST /add`: Adds a new Todo to the list.

## HTML Generation

HTML content is programmatically generated using the `elem-go` library. This includes:

- A form for adding new Todos.
- A dynamic list that renders each Todo as an `li` element with an `input` checkbox to toggle completion.
- Styling for elements is handled inline using the `styles` subpackage.

## Asynchronous Behavior

Using `htmx`, the application performs actions like adding a new Todo or toggling the completion status without a full page reload. This provides a smoother user experience.

## Stopping the Server

To stop the application, use `Ctrl + C` in the terminal.