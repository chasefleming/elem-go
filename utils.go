package elem

import (
	"sort"
	"strings"
)

type Style map[string]string

func (s Style) String() string {
	// Extract the keys and sort them for deterministic order
	keys := make([]string, 0, len(s))
	for key := range s {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var builder strings.Builder
	for _, key := range keys {
		builder.WriteString(key)
		builder.WriteString(": ")
		builder.WriteString(s[key])
		builder.WriteString("; ")
	}

	// Trim the last space
	var styleStr = builder.String()
	if len(styleStr) > 0 {
		styleStr = styleStr[:len(styleStr)-1]
	}

	return styleStr
}

func ApplyStyle(s Style) string {
	return s.String()
}

// If conditionally renders one of the provided elements based on the condition
func If[T any](condition bool, ifTrue, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// TransformEach maps a slice of items to a slice of Elements using the provided function
func TransformEach[T any](items []T, fn func(T) Node) []Node {
	var nodes []Node
	for _, item := range items {
		nodes = append(nodes, fn(item))
	}
	return nodes
}
