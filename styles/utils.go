package styles

import (
	"strconv"
	"strings"
)

// Merge combines multiple styles.Props maps into one, with later styles overriding earlier ones.
func Merge(styleMaps ...Props) Props {
	mergedStyles := Props{}
	for _, styleMap := range styleMaps {
		for key, value := range styleMap {
			mergedStyles[key] = value
		}
	}
	return mergedStyles
}

// Em returns a string representation of the given float as an em value.
func Em(value float64) string {
	return strconv.FormatFloat(value, 'f', 2, 64) + "em"
}

// Percent returns a string representation of the given integer as a percentage.
func Percent(value int) string {
	return strconv.Itoa(value) + "%"
}

// Pixels returns a string representation of the given integer as a pixel value.
func Pixels(value int) string {
	return strconv.Itoa(value) + "px"
}

// Rem returns a string representation of the given float as a rem value.
func Rem(value float64) string {
	return strconv.FormatFloat(value, 'f', 2, 64) + "rem"
}

// RGB returns a string representation of the given RGB values as a CSS RGB value.
func RGB(r, g, b int) string {
	var builder strings.Builder
	builder.WriteString("rgb(")
	builder.WriteString(strconv.Itoa(r))
	builder.WriteString(",")
	builder.WriteString(strconv.Itoa(g))
	builder.WriteString(",")
	builder.WriteString(strconv.Itoa(b))
	builder.WriteString(")")
	return builder.String()
}

// RGBA returns a string representation of the given RGBA values as a CSS RGBA value.
func RGBA(r, g, b, a int) string {
	var builder strings.Builder
	builder.WriteString("rgba(")
	builder.WriteString(strconv.Itoa(r))
	builder.WriteString(",")
	builder.WriteString(strconv.Itoa(g))
	builder.WriteString(",")
	builder.WriteString(strconv.Itoa(b))
	builder.WriteString(",")
	builder.WriteString(strconv.Itoa(a))
	builder.WriteString(")")
	return builder.String()
}

// ViewportHeight returns a string representation of the given integer as a viewport height value.
func ViewportHeight(value int) string {
	return strconv.Itoa(value) + "vh"
}

// ViewportWidth returns a string representation of the given integer as a viewport width value.
func ViewportWidth(value int) string {
	return strconv.Itoa(value) + "vw"
}
