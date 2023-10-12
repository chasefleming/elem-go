package elem

import (
	"sort"
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

func (e *Element) Render() string {

	// Sort the keys for consistent order
	keys := make([]string, 0, len(e.Attrs))
	for k := range e.Attrs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	props := ""
	for _, k := range keys {
		props += ` ` + k + `="` + e.Attrs[k] + `"`
	}

	content := ""
	for _, child := range e.Children {
		switch c := child.(type) {
		case string:
			content += c
		case *Element:
			content += c.Render()
		}
	}

	// Check if the tag is a void element and doesn't have any content
	if content == "" {
		_, exists := voidElements[e.Tag]
		if exists {
			return `<` + e.Tag + props + `>` // No closing tag for void elements
		}
	}

	return `<` + e.Tag + props + `>` + content + `</` + e.Tag + `>`
}

func NewElement(tag string, attrs Attrs, children ...interface{}) *Element {
	return &Element{
		Tag:      tag,
		Attrs:    attrs,
		Children: children,
	}
}
