package styles

import (
	"sort"
	"strings"

	"github.com/chasefleming/elem-go"
)

// Props is a map of CSS properties
type Props map[string]string

// ToInline converts the Props to an inline style string
func (p Props) ToInline() string {
	// Extract the keys and sort them for deterministic order
	keys := make([]string, 0, len(p))
	for key := range p {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var builder strings.Builder
	for _, key := range keys {
		builder.WriteString(key)
		builder.WriteString(": ")
		builder.WriteString(p[key])
		builder.WriteString("; ")
	}

	// Trim the last space
	var styleStr = builder.String()
	if len(styleStr) > 0 {
		styleStr = styleStr[:len(styleStr)-1]
	}

	return styleStr
}

// CSSNode is a Node that renders to a CSS string
type CSSNode string

// RenderTo satisfies part of the Node interface by allowing CSSNode to be written to a strings.Builder
func (cn CSSNode) RenderTo(builder *strings.Builder, opts elem.RenderOptions) {
	builder.WriteString(string(cn))
}

// Render satisfies part of the Node interface by returning the CSS as a string
func (cn CSSNode) Render() string {
	return string(cn)
}

// CSS creates a new CSSNode from the given CSS content string
func CSS(content string) CSSNode {
	return CSSNode(content)
}
