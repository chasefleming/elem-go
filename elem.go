package elem

import (
	"sort"
	"strings"
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

type Attrs map[string]string

type Element struct {
	Tag      string
	Attrs    Attrs
	Children []interface{} // Can be either string (for text) or another Element
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
		builder.WriteString(` `)
		builder.WriteString(k)
		builder.WriteString(`="`)
		builder.WriteString(e.Attrs[k])
		builder.WriteString(`"`)
	}

	// If it's a void element, close it and return
	if _, exists := voidElements[e.Tag]; exists {
		builder.WriteString(`>`)
		return
	}

	// Close opening tag
	builder.WriteString(`>`)

	// Build the content (either child text or nested elements)
	for _, child := range e.Children {
		switch c := child.(type) {
		case string:
			builder.WriteString(c)
		case *Element:
			c.RenderTo(builder)
		}
	}

	// Append closing tag
	builder.WriteString(`</`)
	builder.WriteString(e.Tag)
	builder.WriteString(`>`)
}

func (e *Element) Render() string {
	var builder strings.Builder
	e.RenderTo(&builder)
	return builder.String()
}

func NewElement(tag string, attrs Attrs, children ...interface{}) *Element {
	return &Element{
		Tag:      tag,
		Attrs:    attrs,
		Children: children,
	}
}
