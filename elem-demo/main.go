package main

import (
	"fmt"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/styles"
)

func main() {
	content := elem.Div(attrs.Props{
		attrs.ID:    "container",
		attrs.Class: "my-class",
	},
		elem.H1(nil, elem.Text("Hello, Elem!")),
		elem.H2(nil, elem.Text("Subheading")),
		elem.P(nil, elem.Text("This is a paragraph.")),
	)

	// Convert the Go elements into HTML
	html := content.Render()

	// Print the generated HTML
	fmt.Println(html)

	// Example with styling
	buttonStyle := styles.Props{
		styles.BackgroundColor: "blue",
		styles.Color:           "white",
	}

	button := elem.Button(attrs.Props{
		attrs.Style: buttonStyle.ToInline(),
	}, elem.Text("Click Me"))

	fmt.Println(button.Render())
}
