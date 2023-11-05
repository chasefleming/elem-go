package styles

import (
	"sort"
	"strings"
)

type Props map[string]string

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

type CSSNode string

// RenderTo satisfies part of the Node interface by allowing CSSNode to be written to a strings.Builder
func (cn CSSNode) RenderTo(builder *strings.Builder) {
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
