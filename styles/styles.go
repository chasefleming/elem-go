package styles

import (
	"sort"
	"strings"
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
