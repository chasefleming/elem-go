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
		elem.Script(elem.Attrs{attrs.Src: "https://unpkg.com/htmx.org@1.6.1"}),
	)

	buttonStyle := elem.Style{
		styles.BackgroundColor: "blue",
		styles.Color:           "white",
	}

	body := elem.Body(nil,
		elem.Button(elem.Attrs{
			htmx.HXPost:   "/increment",
			htmx.HXTarget: "#counter-div",
			htmx.HXSwap:   "innerText",
			attrs.Style:   elem.ApplyStyle(buttonStyle),
		}, elem.Text("Increment")),
		elem.Div(elem.Attrs{attrs.ID: "counter-div"}, elem.Text(fmt.Sprintf("%d", counter))),
	)

	htmlDoc := elem.Html(nil, head, body)

	return htmlDoc.Render()
}
