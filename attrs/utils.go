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

// ClassNames joins the non-empty class names with a single space. Empty strings
// and strings consisting only of whitespace are skipped, which makes it easy
// to mix conditional branches inline:
//
//	attrs.Class: attrs.ClassNames(
//	    "btn",
//	    elem.If(isPrimary, "btn-primary", ""),
//	    elem.If(isDisabled, "btn-disabled", ""),
//	)
//
// Each entry is also trimmed of surrounding whitespace.
func ClassNames(classes ...string) string {
	var b strings.Builder
	for _, c := range classes {
		c = strings.TrimSpace(c)
		if c == "" {
			continue
		}
		if b.Len() > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(c)
	}
	return b.String()
}
