package main

import (
	"github.com/chasefleming/elem-go/htmx"
	"strconv"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/styles"
	"github.com/gofiber/fiber/v2"
)

// Todo model
type Todo struct {
	ID    int
	Title string
	Done  bool
}

// Global todos slice (for simplicity)
var todos = []Todo{
	{ID: 1, Title: "First task", Done: false},
	{ID: 2, Title: "Second task", Done: true},
}

func main() {
	app := fiber.New()

	// Routes
	app.Get("/", renderTodosRoute)
	app.Post("/toggle/:id", toggleTodoRoute)
	app.Post("/add", addTodoRoute)

	app.Listen(":3000")
}

func renderTodosRoute(c *fiber.Ctx) error {
	c.Type("html")
	return c.SendString(renderTodos(todos))
}

func toggleTodoRoute(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var updatedTodo Todo
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Done = !todo.Done
			updatedTodo = todos[i]
			break
		}
	}
	c.Type("html")
	return c.SendString(renderSingleTodo(updatedTodo))
}

func addTodoRoute(c *fiber.Ctx) error {
	newTitle := c.FormValue("newTodo")
	if newTitle != "" {
		todos = append(todos, Todo{ID: len(todos) + 1, Title: newTitle, Done: false})
	}
	return c.Redirect("/")
}

func getTodoAttributes(todo Todo) (elem.Attrs, elem.Attrs) {
	checkboxAttributes := elem.Attrs{
		attrs.Type:       "checkbox",
		htmx.HXPost:      "/toggle/" + strconv.Itoa(todo.ID),
		htmx.HXTrigger:   "change",
		htmx.HXIndicator: "#saving-indicator",
	}

	textAttributes := elem.Attrs{}
	if todo.Done {
		checkboxAttributes[attrs.Checked] = "checked"
		textAttributes[attrs.Style] = elem.ApplyStyle(elem.Style{
			styles.TextDecoration: "line-through",
		})
	}

	return checkboxAttributes, textAttributes
}

func renderSingleTodo(todo Todo) string {
	checkboxAttributes, textAttributes := getTodoAttributes(todo)
	checkbox := elem.Input(checkboxAttributes)
	todoItem := elem.Li(nil, checkbox, elem.Span(textAttributes, elem.Text(todo.Title)))
	return todoItem.Render()
}

func renderTodos(todos []Todo) string {
	inputButtonStyle := elem.Style{
		styles.Width:           "100%",
		styles.Padding:         "10px",
		styles.MarginBottom:    "10px",
		styles.Border:          "1px solid #ccc",
		styles.BorderRadius:    "4px",
		styles.BackgroundColor: "#f9f9f9",
	}

	buttonStyle := elem.ApplyStyle(elem.Style{
		styles.BackgroundColor: "#007BFF",
		styles.Color:           "white",
		styles.BorderStyle:     "none",
		styles.BorderRadius:    "4px",
		styles.Cursor:          "pointer",
		styles.Width:           "100%",
		styles.Padding:         "8px 12px",
		styles.FontSize:        "14px",
		styles.Height:          "36px",
		styles.MarginRight:     "10px",
	})

	listContainerStyle := elem.Style{
		styles.ListStyleType: "none",
		styles.Padding:       "0",
		styles.Width:         "100%",
	}

	centerContainerStyle := elem.Style{
		styles.MaxWidth:        "300px",
		styles.Margin:          "40px auto",
		styles.Padding:         "20px",
		styles.Border:          "1px solid #ccc",
		styles.BoxShadow:       "0px 0px 10px rgba(0,0,0,0.1)",
		styles.BackgroundColor: "#f9f9f9",
	}

	headContent := elem.Head(nil,
		elem.Script(elem.Attrs{attrs.Src: "https://unpkg.com/htmx.org"}),
	)
	bodyContent := elem.Div(
		elem.Attrs{attrs.Style: elem.ApplyStyle(centerContainerStyle)},
		elem.H1(nil, elem.Text("Todo List")),
		elem.Form(
			elem.Attrs{attrs.Method: "post", attrs.Action: "/add"},
			elem.Input(
				elem.Attrs{
					attrs.Type:        "text",
					attrs.Name:        "newTodo",
					attrs.Placeholder: "Add new task...",
					attrs.Style:       elem.ApplyStyle(inputButtonStyle),
				},
			),
			elem.Button(
				elem.Attrs{
					attrs.Type:  "submit",
					attrs.Style: buttonStyle,
				},
				elem.Text("Add"),
			),
		),
		elem.Ul(
			elem.Attrs{attrs.Style: elem.ApplyStyle(listContainerStyle)},
			renderTodoItems(todos)...,
		),
		elem.Span(
			elem.Attrs{
				attrs.ID:    "saving-indicator",
				attrs.Style: elem.ApplyStyle(elem.Style{styles.Display: "none"}),
			},
			elem.Text("Saving..."),
		),
	)
	htmlContent := elem.Html(nil, headContent, bodyContent)

	return htmlContent.Render()
}

func renderTodoItems(todos []Todo) []elem.Node {
	return elem.TransformEach(todos, func(todo Todo) elem.Node {
		checkboxAttributes, textAttributes := getTodoAttributes(todo)
		checkbox := elem.Input(checkboxAttributes)
		return elem.Li(nil, checkbox, elem.Span(textAttributes, elem.Text(todo.Title)))
	})
}
