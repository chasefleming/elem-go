package elem

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
