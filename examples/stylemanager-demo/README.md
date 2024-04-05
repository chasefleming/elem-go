# Advanced Styling App with elem-go and StyleManager

This web application demonstrates the power of advanced CSS styling within a Go server environment, using the `elem-go` library and its `StyleManager` for dynamic styling. It features a button with hover effects, an animated element, and a responsive design element that adjusts styles based on the viewport width.

## Prerequisites

Ensure that Go is installed on your system.

## Installation

Clone or download the repository, then run the following commands to download the necessary packages:

```bash
go mod tidy
```

This will install `elem-go` and the `styles` subpackage required to run the application.

## Running the Application

To run the application, execute the following command:

```bash
go run main.go
```

The server will start on `localhost` at port `8080`. You can view the application by navigating to `http://localhost:8080` in your web browser.

## Features

**Button Hover Effect**: The button changes color and scale when hovered over, providing visual feedback to the user.
**Animated Element**: The animated element moves across the screen in a loop, demonstrating the use of animations with `StyleManager`.
**Responsive Design**: The application adjusts the background color based on the viewport width, showcasing the use of media queries for responsive styling.

## Advanced CSS Styling with `StyleManager`

The `StyleManager` within the `styles` package provides a powerful solution for managing CSS styles in Go-based web applications. It enables the creation of dynamic and responsive styles, including pseudo-classes, animations, and media queries, directly in Go.

To learn more about `StyleManager` and its advanced features, refer to the [StyleManager documentation](../../styles/STYLEMANAGER.md).

## Stopping the Server

To stop the application, press `Ctrl + C` in the terminal where the server is running.