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
		Padding: "10px",
		Margin:  "10px", // From secondaryStyle
		Color:   "blue", // From tertiaryStyle
		Border:  "1px solid black",
	}

	assert.Equal(t, expectedStyle, mergedStyle)
}

func TestPercent(t *testing.T) {
	assert.Equal(t, "100%", Percent(100))
	assert.Equal(t, "0%", Percent(0))
	assert.Equal(t, "50%", Percent(50))
	assert.Equal(t, "25%", Percent(25))
}

func TestPixels(t *testing.T) {
	assert.Equal(t, "100px", Pixels(100))
	assert.Equal(t, "0px", Pixels(0))
	assert.Equal(t, "50px", Pixels(50))
	assert.Equal(t, "25px", Pixels(25))
}

func TestViewportHeight(t *testing.T) {
	assert.Equal(t, "100vh", ViewportHeight(100))
	assert.Equal(t, "0vh", ViewportHeight(0))
	assert.Equal(t, "50vh", ViewportHeight(50))
	assert.Equal(t, "25vh", ViewportHeight(25))
}

func TestViewportWidth(t *testing.T) {
	assert.Equal(t, "100vw", ViewportWidth(100))
	assert.Equal(t, "0vw", ViewportWidth(0))
	assert.Equal(t, "50vw", ViewportWidth(50))
	assert.Equal(t, "25vw", ViewportWidth(25))
}

func TestEm(t *testing.T) {
	assert.Equal(t, "100.00em", Em(100))
	assert.Equal(t, "0.00em", Em(0))
	assert.Equal(t, "50.00em", Em(50))
	assert.Equal(t, "25.00em", Em(25))
}

func TestRem(t *testing.T) {
	assert.Equal(t, "100.00rem", Rem(100))
	assert.Equal(t, "0.00rem", Rem(0))
	assert.Equal(t, "50.00rem", Rem(50))
	assert.Equal(t, "25.00rem", Rem(25))
}
