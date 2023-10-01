package elem

func TransformEach[T any](items []T, fn func(T) *Element) []*Element {
	var elements []*Element
	for _, item := range items {
		elements = append(elements, fn(item))
	}
	return elements
}
