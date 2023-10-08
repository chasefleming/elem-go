package main

import (
	"fmt"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define a simple structure to hold our form data
	var formData struct {
		Name  string
		Email string
	}

	app.Post("/submit-form", func(c *fiber.Ctx) error {
		// Capture the form data
		formData.Name = c.FormValue("name")
		formData.Email = c.FormValue("email")

		// Send a response with the captured data
		return c.SendString(fmt.Sprintf("Name: %s, Email: %s", formData.Name, formData.Email))
	})

	app.Get("/", func(c *fiber.Ctx) error {

		// Define the head with the htmx script and styling
		head := elem.Head(nil,
			elem.Script(elem.Attrs{attrs.Src: "https://unpkg.com/htmx.org@1.9.6"}),
		)

		// Define the body content for our form page
		body := elem.Body(nil,
			elem.H1(nil, elem.Text("Simple Form App")),
			elem.Form(elem.Attrs{
				attrs.Action: "/submit-form",
				attrs.Method: "POST",
				htmx.HXPost:  "/submit-form",
				htmx.HXSwap:  "outerHTML",
			},
				elem.Label(elem.Attrs{attrs.For: "name"}, "Name: "),
				elem.Input(elem.Attrs{
					attrs.Type: "text",
					attrs.Name: "name",
					attrs.ID:   "name",
				}),
				elem.Br(nil),
				elem.Label(elem.Attrs{attrs.For: "email"}, "Email: "),
				elem.Input(elem.Attrs{
					attrs.Type: "email",
					attrs.Name: "email",
					attrs.ID:   "email",
				}),
				elem.Br(nil),
				elem.Input(elem.Attrs{
					attrs.Type:  "submit",
					attrs.Value: "Submit",
				}),
			),
			elem.Div(elem.Attrs{attrs.ID: "response"}, elem.Text("")),
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
