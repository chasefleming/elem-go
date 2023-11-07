package styles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	baseStyle := Props{
		Width: "100px",
		Color: "blue",
	}

	additionalStyle := Props{
		Color:           "red", // This should override the blue color in baseStyle
		BackgroundColor: "yellow",
	}

	expectedMergedStyle := Props{
		Width:           "100px",
		Color:           "red",
		BackgroundColor: "yellow",
	}

	mergedStyle := Merge(baseStyle, additionalStyle)

	assert.Equal(t, expectedMergedStyle, mergedStyle)
}

func TestMergeThreeStyles(t *testing.T) {
	baseStyle := Props{
		Padding: "10px",
		Margin:  "5px",
	}

	secondaryStyle := Props{
		Margin: "10px", // This should override the baseStyle margin
		Color:  "red",
	}

	tertiaryStyle := Props{
		Color:  "blue", // This should override the secondaryStyle color
		Border: "1px solid black",
	}

	mergedStyle := Merge(baseStyle, secondaryStyle, tertiaryStyle)

	expectedStyle := Props{
		"padding": "10px",
		"margin":  "10px", // From secondaryStyle
		"color":   "blue", // From tertiaryStyle
		"border":  "1px solid black",
	}

	assert.Equal(t, expectedStyle, mergedStyle)
}
