package elem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	trueElement := Div(nil, Text("True Condition"))
	falseElement := Div(nil, Text("False Condition"))

	resultTrue := If(true, trueElement, falseElement)
	assert.Equal(t, trueElement.Render(), resultTrue.Render())

	resultFalse := If(false, trueElement, falseElement)
	assert.Equal(t, falseElement.Render(), resultFalse.Render())
}

func TestIfUsingNone(t *testing.T) {
	trueElement := Div(nil, Text("True Condition"))

	resultWithNone := If[Node](true, trueElement, None())
	assert.Equal(t, trueElement.Render(), resultWithNone.Render(), "If should render the true element when the condition is true")

	resultWithNoneFalse := If[Node](false, trueElement, None())
	assert.Equal(t, "", resultWithNoneFalse.Render(), "If should render nothing (empty string) when the condition is false and None is used")
}

func TestTransformEach(t *testing.T) {
	items := []string{"Item 1", "Item 2", "Item 3"}

	elements := TransformEach(items, func(item string) Node {
		return Li(nil, Text(item))
	})

	assert.Equal(t, len(items), len(elements))

	for i, el := range elements {
		expected := `<li>` + items[i] + `</li>`
		assert.Equal(t, expected, el.Render())
	}
}
