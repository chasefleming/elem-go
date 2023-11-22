package elem

import (
	"sort"
	"strings"

	"github.com/chasefleming/elem-go/attrs"
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

// List of boolean attributes. Boolean attributes can't have literal values. The presence of an boolean
// attribute represents the "true" value. To represent the "false" value, the attribute has to be omitted.
// See https://html.spec.whatwg.org/multipage/indices.html#attributes-3 for reference
var booleanAttrs = map[string]struct{}{
	attrs.AllowFullscreen: {},
	attrs.Async:           {},
	attrs.Autofocus:       {},
	attrs.Autoplay:        {},
	attrs.Checked:         {},
	attrs.Controls:        {},
	attrs.Defer:           {},
	attrs.Disabled:        {},
	attrs.Ismap:           {},
	attrs.Loop:            {},
	attrs.Multiple:        {},
	attrs.Muted:           {},
	attrs.Novalidate:      {},
	attrs.Open:            {},
	attrs.Playsinline:     {},
	attrs.Readonly:        {},
	attrs.Required:        {},
	attrs.Selected:        {},
}

type Node interface {
	RenderTo(builder *strings.Builder)
	Render() string
}

type TextNode string

func (t TextNode) RenderTo(builder *strings.Builder) {
	builder.WriteString(string(t))
}

func (t TextNode) Render() string {
	return string(t)
}

type RawNode string

func (r RawNode) RenderTo(builder *strings.Builder) {
	builder.WriteString(string(r))
}

func (r RawNode) Render() string {
	return string(r)
}

type CommentNode string

func (c CommentNode) RenderTo(builder *strings.Builder) {
	builder.WriteString("<!-- ")
	builder.WriteString(string(c))
	builder.WriteString(" -->")
}

func (c CommentNode) Render() string {
	var builder strings.Builder
	c.RenderTo(&builder)
	return builder.String()
}

type Element struct {
	Tag      string
	Attrs    attrs.Props
	Children []Node
}

func (e *Element) RenderTo(builder *strings.Builder) {
	// Start with opening tag
	builder.WriteString("<")
	builder.WriteString(e.Tag)

	// Sort the keys for consistent order
	keys := make([]string, 0, len(e.Attrs))
	for k := range e.Attrs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Append the attributes to the builder
	for _, k := range keys {
		e.renderAttrTo(k, builder)
	}

	// If it's a void element, close it and return
	if _, exists := voidElements[e.Tag]; exists {
		builder.WriteString(`>`)
		return
	}

	// Close opening tag
	builder.WriteString(`>`)

	// Build the content
	for _, child := range e.Children {
		child.RenderTo(builder)
	}

	// Append closing tag
	builder.WriteString(`</`)
	builder.WriteString(e.Tag)
	builder.WriteString(`>`)
}

// return string representation of given attribute with its value
func (e *Element) renderAttrTo(attrName string, builder *strings.Builder) {
	if _, exists := booleanAttrs[attrName]; exists {
		// boolean attribute presents its name only if the value is "true"
		if e.Attrs[attrName] == "true" {
			builder.WriteString(` `)
			builder.WriteString(attrName)
		}
	} else {
		// regular attribute has a name and a value
		builder.WriteString(` `)
		builder.WriteString(attrName)
		builder.WriteString(`="`)
		builder.WriteString(e.Attrs[attrName])
		builder.WriteString(`"`)
	}
}

func (e *Element) Render() string {
	var builder strings.Builder
	e.RenderTo(&builder)
	return builder.String()
}

func NewElement(tag string, attrs attrs.Props, children ...Node) *Element {
	return &Element{
		Tag:      tag,
		Attrs:    attrs,
		Children: children,
	}
}
