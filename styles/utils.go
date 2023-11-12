package styles

import (
	"fmt"
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
	return fmt.Sprintf("%d%%", value)
}

// Pixels returns a string representation of the given integer as a pixel value.
func Pixels(value int) string {
	return fmt.Sprintf("%dpx", value)
}

// ViewportHeight returns a string representation of the given integer as a viewport height value.
func ViewportHeight(value int) string {
	return fmt.Sprintf("%dvh", value)
}

// ViewportWidth returns a string representation of the given integer as a viewport width value.
func ViewportWidth(value int) string {
	return fmt.Sprintf("%dvw", value)
}

// Em returns a string representation of the given float as an em value.
func Em(value float64) string {
	return fmt.Sprintf("%.2fem", value)
}

// Rem returns a string representation of the given float as a rem value.
func Rem(value float64) string {
	return fmt.Sprintf("%.2frem", value)
}
