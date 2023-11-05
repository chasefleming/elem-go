package main

import (
	"fmt"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
	"github.com/chasefleming/elem-go/styles"
	"net/http"
	"strconv"
)

var counter int

func main() {
	http.HandleFunc("/", counterHomeHandler)
	http.HandleFunc("/increment", incrementCounterHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func counterHomeHandler(w http.ResponseWriter, r *http.Request) {
	content := generateCounterContent()
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(content))
}

func incrementCounterHandler(w http.ResponseWriter, r *http.Request) {
	counter++
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(counter)))
}

func generateCounterContent() string {
	head := elem.Head(nil,
		elem.Script(attrs.Props{attrs.Src: "https://unpkg.com/htmx.org@1.6.1"}),
	)

	buttonStyle := styles.Props{
		styles.BackgroundColor: "blue",
		styles.Color:           "white",
	}

	body := elem.Body(nil,
		elem.Button(attrs.Props{
			htmx.HXPost:   "/increment",
			htmx.HXTarget: "#counter-div",
			htmx.HXSwap:   "innerText",
			attrs.Style:   buttonStyle.ToInline(),
		}, elem.Text("Increment")),
		elem.Div(attrs.Props{attrs.ID: "counter-div"}, elem.Text(fmt.Sprintf("%d", counter))),
	)

	htmlDoc := elem.Html(nil, head, body)

	return htmlDoc.Render()
}
