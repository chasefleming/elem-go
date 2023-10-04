package elem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShow(t *testing.T) {
	trueElement := Div(nil, Text("True Condition"))
	falseElement := Div(nil, Text("False Condition"))

	resultTrue := Show(true, trueElement, falseElement)
	assert.Equal(t, trueElement.Render(), resultTrue.Render())

	resultFalse := Show(false, trueElement, falseElement)
	assert.Equal(t, falseElement.Render(), resultFalse.Render())
}

func TestTransformEach(t *testing.T) {
	items := []string{"Item 1", "Item 2", "Item 3"}

	elements := TransformEach(items, func(item string) *Element {
		return Li(nil, Text(item))
	})

	assert.Equal(t, len(items), len(elements))

	for i, el := range elements {
		expected := `<li>` + items[i] + `</li>`
		assert.Equal(t, expected, el.Render())
	}
}
