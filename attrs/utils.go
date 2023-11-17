package attrs

import "strings"

// DataAttr returns the name for a data attribute
func DataAttr(name string) string {
	var builder strings.Builder
	builder.WriteString(DataPrefix)
	builder.WriteString(name)
	return builder.String()
}
