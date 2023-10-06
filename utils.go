package elem

import (
	"sort"
)

type Style map[string]string

func (s Style) String() string {
	// Extract the keys and sort them for deterministic order
	keys := make([]string, 0, len(s))
	for key := range s {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var styleStr string
	for _, key := range keys {
		styleStr += key + ": " + s[key] + "; "
	}

	// Trim the last space
	if len(styleStr) > 0 {
		styleStr = styleStr[:len(styleStr)-1]
	}
	return styleStr
}

// Show conditionally renders one of the provided elements based on the condition
func Show(condition bool, ifTrue, ifFalse *Element) *Element {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// TransformEach maps a slice of items to a slice of Elements using the provided function
func TransformEach[T any](items []T, fn func(T) *Element) []*Element {
	var elements []*Element
	for _, item := range items {
		elements = append(elements, fn(item))
	}
	return elements
}
