package elem

import (
	"github.com/chasefleming/elem-go/styles"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStyleString(t *testing.T) {
	style := Style{
		"background-color": "blue",
		"color":            "white",
		"font-size":        "16px",
	}

	assert.Equal(t, "background-color: blue; color: white; font-size: 16px;", style.String())
}

func TestApplyStyle(t *testing.T) {
	style := Style{
		styles.BackgroundColor: "blue",
		styles.Color:           "white",
		styles.FontSize:        "16px",
	}

	expected := "background-color: blue; color: white; font-size: 16px;"
	actual := ApplyStyle(style)

	assert.Equal(t, expected, actual)
}

func TestIf(t *testing.T) {
	trueElement := Div(nil, Text("True Condition"))
	falseElement := Div(nil, Text("False Condition"))

	resultTrue := If(true, trueElement, falseElement)
	assert.Equal(t, trueElement.Render(), resultTrue.Render())

	resultFalse := If(false, trueElement, falseElement)
	assert.Equal(t, falseElement.Render(), resultFalse.Render())
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
