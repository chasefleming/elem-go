package styles

import (
	"fmt"
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

// Float returns a string representation of the float with 2 decimal places
func Float(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

// Int returns a string representation of the int
func Int(value int) string {
	return strconv.Itoa(value)
}

// Percent returns a string representation of the given integer as a percentage.
func Percent(value int) string {
	return strconv.Itoa(value) + "%"
}

// Pixels returns a string representation of the given integer as a pixel value.
func Pixels(value int) string {
	if value == 0 {
		return "0"
	}
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
func RGBA(r, g, b int, a float64) string {
	var builder strings.Builder
	builder.WriteString("rgba(")
	builder.WriteString(strconv.Itoa(r))
	builder.WriteString(",")
	builder.WriteString(strconv.Itoa(g))
	builder.WriteString(",")
	builder.WriteString(strconv.Itoa(b))
	builder.WriteString(",")

	// Convert the alpha float value to a string with desired precision
	alphaStr := strconv.FormatFloat(a, 'f', 1, 64)
	builder.WriteString(alphaStr)
	builder.WriteString(")")
	return builder.String()
}

// URL returns a string representation of the given string as a CSS URL value.
func URL(url string) string {
	var builder strings.Builder
	builder.WriteString("url('")
	builder.WriteString(url)
	builder.WriteString("')")
	return builder.String()
}

// Var returns a string representation of the given string as a CSS var value.
func Var(name string) string {
	var builder strings.Builder
	builder.WriteString("var(--")
	builder.WriteString(name)
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

// ViewportMin returns a string representation of the given integer as a viewport minimum value.
func ViewportMin(value int) string {
	return strconv.Itoa(value) + "vmin"
}

// ViewportMax returns a string representation of the given integer as a viewport maximum value.
func ViewportMax(value int) string {
	return strconv.Itoa(value) + "vmax"
}

// Seconds returns a string representation of the given float64 or int
func Seconds(value float64) string {
	return strconv.FormatFloat(value, 'f', 2, 64) + "s"
}

// Milliseconds returns a string representation of the given int
func Milliseconds(value int) string {
	return strconv.Itoa(value) + "ms"
}

func HSL(hue, saturation, lightness int) string {
	var builder strings.Builder
	builder.WriteString("hsl(")
	builder.WriteString(strconv.FormatInt(int64(hue), 10))
	builder.WriteString(",")
	builder.WriteString(strconv.FormatInt(int64(hue), 10))
	builder.WriteString("%,")
	builder.WriteString(strconv.FormatInt(int64(hue), 10))
	builder.WriteString("%")
	builder.WriteString(")")
	return builder.String()
}

func HSLA(hue, saturation, lightness int, alpha float64) string {
	var builder strings.Builder
	builder.WriteString("hsla(")
	builder.WriteString(strconv.FormatInt(int64(hue), 10))
	builder.WriteString(",")
	builder.WriteString(strconv.FormatInt(int64(hue), 10))
	builder.WriteString("%,")
	builder.WriteString(strconv.FormatInt(int64(hue), 10))
	builder.WriteString("%,")
	builder.WriteString(strconv.FormatFloat(alpha, 'f', 2, 64))
	builder.WriteString(")")
	return builder.String()
}
