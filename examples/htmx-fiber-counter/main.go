package main

import (
	"fmt"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
	"github.com/chasefleming/elem-go/styles"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	count := 0

	app.Post("/increment", func(c *fiber.Ctx) error {
		count++
		return c.SendString(fmt.Sprintf("%d", count))
	})

	app.Post("/decrement", func(c *fiber.Ctx) error {
		count--
		return c.SendString(fmt.Sprintf("%d", count))
	})

	app.Get("/", func(c *fiber.Ctx) error {

		// Define the head with the htmx script and styling
		head := elem.Head(nil,
			elem.Script(attrs.Props{attrs.Src: "https://unpkg.com/htmx.org@1.9.6"}),
		)

		// Define styling
		bodyStyle := styles.Props{
			styles.BackgroundColor: "#f4f4f4",
			styles.FontFamily:      "Arial, sans-serif",
			styles.Height:          "100vh",
			styles.Display:         "flex",
			styles.FlexDirection:   "column",
			styles.AlignItems:      "center",
			styles.JustifyContent:  "center",
		}

		buttonStyle := styles.Props{
			styles.Padding:         "10px 20px",
			styles.BackgroundColor: "#007BFF",
			styles.Color:           "#fff",
			styles.BorderColor:     "#007BFF",
			styles.BorderRadius:    "5px",
			styles.Margin:          "10px",
			styles.Cursor:          "pointer",
		}

		// Define the body content for our counter page
		body := elem.Body(attrs.Props{
			attrs.Style: bodyStyle.ToInline(),
		},
			elem.H1(nil, elem.Text("Counter App")),
			elem.Div(attrs.Props{attrs.ID: "count"}, elem.Text(fmt.Sprintf("%d", count))),
			elem.Button(attrs.Props{
				htmx.HXPost:   "/increment",
				htmx.HXTarget: "#count",
				attrs.Style:   buttonStyle.ToInline(),
			}, elem.Text("+")),
			elem.Button(attrs.Props{
				htmx.HXPost:   "/decrement",
				htmx.HXTarget: "#count",
				attrs.Style:   buttonStyle.ToInline(),
			}, elem.Text("-")),
		)

		// Wrap the head and body content within an HTML tag
		pageContent := elem.Html(nil, head, body)

		html := pageContent.Render()

		// Specify that the response content type is HTML before sending the response
		c.Type("html")
		return c.SendString(html)
	})

	app.Listen(":3000")
}
