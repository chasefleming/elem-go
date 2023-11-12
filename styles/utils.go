package styles

import (
	"strconv"
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

// Percent returns a string representation of the given integer as a percentage.
func Percent(value int) string {
	return strconv.Itoa(value) + "%"
}

// Pixels returns a string representation of the given integer as a pixel value.
func Pixels(value int) string {
	return strconv.Itoa(value) + "px"
}

// ViewportHeight returns a string representation of the given integer as a viewport height value.
func ViewportHeight(value int) string {
	return strconv.Itoa(value) + "vh"
}

// ViewportWidth returns a string representation of the given integer as a viewport width value.
func ViewportWidth(value int) string {
	return strconv.Itoa(value) + "vw"
}

// Em returns a string representation of the given float as an em value.
func Em(value float64) string {
	return strconv.FormatFloat(value, 'f', 2, 64) + "em"
}

// Rem returns a string representation of the given float as a rem value.
func Rem(value float64) string {
	return strconv.FormatFloat(value, 'f', 2, 64) + "rem"
}
