package elem

import (
	"github.com/chasefleming/elem-go/attrs"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ========== Document Structure ==========

func TestBody(t *testing.T) {
	expected := `<body class="page-body"><p>Welcome to Elem!</p></body>`
	el := Body(Attrs{attrs.Class: "page-body"}, P(nil, Text("Welcome to Elem!")))
	assert.Equal(t, expected, el.Render())
}

func TestHead(t *testing.T) {
	// ... [Code for the Head test here]
}

func TestHtml(t *testing.T) {
	expected := `<html lang="en"><head><meta charset="UTF-8"><title>Elem Page</title></head><body><p>Welcome to Elem!</p></body></html>`
	el := Html(Attrs{attrs.Lang: "en"},
		Head(nil,
			Meta(Attrs{attrs.Charset: "UTF-8"}),
			Title(nil, Text("Elem Page")),
		),
		Body(nil, P(nil, Text("Welcome to Elem!"))),
	)
	assert.Equal(t, expected, el.Render())
}

// ========== Text Formatting and Structure ==========

func TestA(t *testing.T) {
	expected := `<a href="https://example.com">Visit Example</a>`
	el := A(Attrs{attrs.Href: "https://example.com"}, Text("Visit Example"))
	assert.Equal(t, expected, el.Render())
}

func TestBlockquote(t *testing.T) {
	expected := `<blockquote>Quote text</blockquote>`
	el := Blockquote(nil, Text("Quote text"))
	assert.Equal(t, expected, el.Render())
}

func TestBr(t *testing.T) {
	expected := `<br>`
	el := Br(nil)
	assert.Equal(t, expected, el.Render())
}

func TestCode(t *testing.T) {
	expected := `<code>Code snippet</code>`
	el := Code(nil, Text("Code snippet"))
	assert.Equal(t, expected, el.Render())
}

func TestDiv(t *testing.T) {
	expected := `<div class="container">Hello, Elem!</div>`
	el := Div(Attrs{attrs.Class: "container"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestEm(t *testing.T) {
	expected := `<em>Italic text</em>`
	el := Em(nil, Text("Italic text"))
	assert.Equal(t, expected, el.Render())
}

func TestH1(t *testing.T) {
	expected := `<h1 class="title">Hello, Elem!</h1>`
	el := H1(Attrs{attrs.Class: "title"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestH2(t *testing.T) {
	expected := `<h2 class="subtitle">Hello, Elem!</h2>`
	el := H2(Attrs{attrs.Class: "subtitle"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestH3(t *testing.T) {
	expected := `<h3>Hello, Elem!</h3>`
	el := H3(nil, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestHr(t *testing.T) {
	expected := `<hr>`
	el := Hr(nil)
	assert.Equal(t, expected, el.Render())
}

func TestP(t *testing.T) {
	expected := `<p>Hello, Elem!</p>`
	el := P(nil, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestPre(t *testing.T) {
	expected := `<pre>Preformatted text</pre>`
	el := Pre(nil, Text("Preformatted text"))
	assert.Equal(t, expected, el.Render())
}

func TestSpan(t *testing.T) {
	expected := `<span class="highlight">Hello, Elem!</span>`
	el := Span(Attrs{attrs.Class: "highlight"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestStrong(t *testing.T) {
	expected := `<strong>Bold text</strong>`
	el := Strong(nil, Text("Bold text"))
	assert.Equal(t, expected, el.Render())
}

// ========== Lists ==========

func TestLi(t *testing.T) {
	expected := `<li>Item 1</li>`
	el := Li(nil, Text("Item 1"))
	assert.Equal(t, expected, el.Render())
}

func TestUl(t *testing.T) {
	expected := `<ul><li>Item 1</li><li>Item 2</li></ul>`
	el := Ul(nil, Li(nil, Text("Item 1")), Li(nil, Text("Item 2")))
	assert.Equal(t, expected, el.Render())
}

// ========== Forms ==========

func TestButton(t *testing.T) {
	expected := `<button class="btn">Click Me</button>`
	el := Button(Attrs{attrs.Class: "btn"}, Text("Click Me"))
	assert.Equal(t, expected, el.Render())
}

// ========== Hyperlinks and Multimedia ==========

func TestImg(t *testing.T) {
	expected := `<img alt="An image" src="image.jpg">`
	el := Img(Attrs{attrs.Src: "image.jpg", attrs.Alt: "An image"})
	assert.Equal(t, expected, el.Render())
}

// ========== Meta Elements ==========

func TestScript(t *testing.T) {
	expected := `<script src="https://example.com/script.js"></script>`
	el := Script(Attrs{attrs.Src: "https://example.com/script.js"})
	assert.Equal(t, expected, el.Render())
}
