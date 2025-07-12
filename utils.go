package elem

import "strings"

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

// EscapeNodeContents escapes HTML5 special characters in a string to ensure safe rendering as a text node
func EscapeNodeContents(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	// technically, > does not need escaping in HTML5, but we do it for consistency
	s = strings.ReplaceAll(s, ">", "&gt;")
	// note that single and double quotes (', ") do not need escaping in HTML5 text nodes

	return s
}
