# Counter Application with elem-go and htmx

A basic web application demonstrating a simple counter that can be incremented without reloading the page. It uses `elem-go` to construct HTML elements programmatically and `htmx` for handling the asynchronous increment of the counter.

## Prerequisites

Ensure that Go is installed on your system.

## Installation

Clone or download the repository, then run the following commands to download the necessary packages:

```bash
go mod tidy
```

This will install elem-go and the htmx subpackage required to run the application.

## Running the Application

To run the application, execute the following command:

```bash
go run main.go
```

The server will start on localhost at port 8080. You can view the application by navigating to `http://localhost:8080` in your web browser.

## Features

**Increment Counter**: Click the "Increment" button to increase the counter value. This action will send a POST request to `/increment` and update the counter display asynchronously.

## Structure

- `GET /`: Renders the home page with the current value of the counter and an "Increment" button.
- `POST /increment`: Increases the counter by one and returns the updated value.

## Asynchronous Updates with htmx

The counter is updated asynchronously using htmx. When the "Increment" button is clicked, htmx sends a POST request to `/increment`, receives the new counter value, and updates the relevant part of the page without reloading.

## HTML Generation

The HTML for the page is generated on the server using the elem-go library, with elements such as `button` and `div` created programmatically. htmx attributes are applied directly to elements to enable the asynchronous behavior.

## Stopping the Server

To stop the application, press `Ctrl + C` in the terminal where the server is running.