package elem

import (
	"strings"

	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/options"
)

// List of HTML5 void elements. Void elements, also known as self-closing or empty elements,
// are elements that don't have a closing tag because they can't contain any content.
// For example, the <img> tag cannot wrap text or other tags, it stands alone, so it doesn't have a closing tag.
var voidElements = map[string]struct{}{
	"area":    {},
	"base":    {},
	"br":      {},
	"col":     {},
	"command": {},
	"embed":   {},
	"hr":      {},
	"img":     {},
	"input":   {},
	"keygen":  {},
	"link":    {},
	"meta":    {},
	"param":   {},
	"source":  {},
	"track":   {},
	"wbr":     {},
}

type Node interface {
	RenderTo(builder *strings.Builder, opts options.RenderOptions)
	Render() string
	RenderWithOptions(opts options.RenderOptions) string
}

// NoneNode represents a node that renders nothing.
type NoneNode struct{}

// RenderTo for NoneNode does nothing.
func (n NoneNode) RenderTo(builder *strings.Builder, opts options.RenderOptions) {
	// Intentionally left blank to render nothing
}

// Render for NoneNode returns an empty string.
func (n NoneNode) Render() string {
	return ""
}

// RenderWithOptions for NoneNode returns an empty string.
func (n NoneNode) RenderWithOptions(opts options.RenderOptions) string {
	return ""
}

type TextNode string

func (t TextNode) RenderTo(builder *strings.Builder, opts options.RenderOptions) {
	builder.WriteString(string(t))
}

func (t TextNode) Render() string {
	return string(t)
}

func (t TextNode) RenderWithOptions(opts options.RenderOptions) string {
	return string(t)
}

type RawNode string

func (r RawNode) RenderTo(builder *strings.Builder, opts options.RenderOptions) {
	builder.WriteString(string(r))
}

func (r RawNode) Render() string {
	return string(r)
}

func (t RawNode) RenderWithOptions(opts options.RenderOptions) string {
	return string(t)
}

type CommentNode string

func (c CommentNode) RenderTo(builder *strings.Builder, opts options.RenderOptions) {
	builder.WriteString("<!-- ")
	builder.WriteString(string(c))
	builder.WriteString(" -->")
}

func (c CommentNode) Render() string {
	return c.RenderWithOptions(options.RenderOptions{})
}

func (c CommentNode) RenderWithOptions(opts options.RenderOptions) string {
	var builder strings.Builder
	c.RenderTo(&builder, opts)
	return builder.String()
}

type Element struct {
	Tag      string
	Attrs    attrs.Props
	Children []Node
}

func (e *Element) RenderTo(builder *strings.Builder, opts options.RenderOptions) {
	// The HTML tag needs a doctype preamble in order to ensure
	// browsers don't render in legacy/quirks mode
	// https://developer.mozilla.org/en-US/docs/Glossary/Doctype
	if !opts.DisableHtmlPreamble && e.Tag == "html" {
		builder.WriteString("<!DOCTYPE html>")
	}

	// Start with opening tag
	builder.WriteString("<")
	builder.WriteString(e.Tag)

	e.Attrs.RenderTo(builder, opts)
	// If it's a void element, close it and return
	if _, exists := voidElements[e.Tag]; exists {
		builder.WriteString(`>`)
		return
	}

	// Close opening tag
	builder.WriteString(`>`)

	// Build the content
	for _, child := range e.Children {
		if child == nil {
			continue
		}
		child.RenderTo(builder, opts)
	}

	// Append closing tag
	builder.WriteString(`</`)
	builder.WriteString(e.Tag)
	builder.WriteString(`>`)
}

func (e *Element) Render() string {
	return e.RenderWithOptions(options.RenderOptions{})
}

func (e *Element) RenderWithOptions(opts options.RenderOptions) string {
	var builder strings.Builder
	e.RenderTo(&builder, opts)
	return builder.String()
}

func newElement(tag string, children ...Node) *Element {
	attrs, ok := children[0].(attrs.Props)

	if ok {
		children = children[1:]
	}

	return &Element{
		Tag:      tag,
		Attrs:    attrs,
		Children: children,
	}
}
