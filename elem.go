package elem

import (
	"sort"
)

type Element struct {
	Tag      string
	Props    map[string]string
	Children []interface{} // Can be either string (for text) or another Element
}

func NewElement(tag string, props Props, children ...interface{}) *Element {
	return &Element{
		Tag:      tag,
		Props:    props,
		Children: children,
	}
}

func (e *Element) Render() string {
	// Sort the keys for consistent order
	keys := make([]string, 0, len(e.Props))
	for k := range e.Props {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	props := ""
	for _, k := range keys {
		props += ` ` + k + `="` + e.Props[k] + `"`
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

	if content == "" {
		return `<` + e.Tag + props + ` />` // Self-closing tag
	}

	return `<` + e.Tag + props + `>` + content + `</` + e.Tag + `>`
}
