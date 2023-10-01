package elem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA(t *testing.T) {
	expected := `<a href="https://example.com">Visit Example</a>`
	el := A(Props{Href: "https://example.com"}, Text("Visit Example"))
	assert.Equal(t, expected, el.Render())
}

func TestDiv(t *testing.T) {
	expected := `<div class="container">Hello, Elem!</div>`
	el := Div(Props{Class: "container"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestH1(t *testing.T) {
	expected := `<h1 class="title">Hello, Elem!</h1>`
	el := H1(Props{Class: "title"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestH2(t *testing.T) {
	expected := `<h2 class="subtitle">Hello, Elem!</h2>`
	el := H2(Props{Class: "subtitle"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestH3(t *testing.T) {
	expected := `<h3>Hello, Elem!</h3>`
	el := H3(nil, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestImg(t *testing.T) {
	expected := `<img alt="An image" src="image.jpg" />`
	el := Img(Props{Src: "image.jpg", Alt: "An image"})
	assert.Equal(t, expected, el.Render())
}

func TestLi(t *testing.T) {
	expected := `<li>Item 1</li>`
	el := Li(nil, Text("Item 1"))
	assert.Equal(t, expected, el.Render())
}

func TestP(t *testing.T) {
	expected := `<p>Hello, Elem!</p>`
	el := P(nil, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestSpan(t *testing.T) {
	expected := `<span class="highlight">Hello, Elem!</span>`
	el := Span(Props{Class: "highlight"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestUl(t *testing.T) {
	expected := `<ul><li>Item 1</li><li>Item 2</li></ul>`
	el := Ul(nil, Li(nil, Text("Item 1")), Li(nil, Text("Item 2")))
	assert.Equal(t, expected, el.Render())
}
