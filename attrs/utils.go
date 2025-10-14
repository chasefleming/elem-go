package attrs

import "strings"

// DataAttr returns the name for a data attribute.
func DataAttr(name string) string {
	var builder strings.Builder
	builder.WriteString(DataPrefix)
	builder.WriteString(name)
	return builder.String()
}

// Merge merges multiple attribute maps into a single map, with later maps overriding earlier ones.
func Merge(attrMaps ...Props) Props {
	mergedAttrs := Props{}
	for _, attrMap := range attrMaps {
		for key, value := range attrMap {
			mergedAttrs[key] = value
		}
	}
	return mergedAttrs
}
