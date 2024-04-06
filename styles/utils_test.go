package styles

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestEm(t *testing.T) {
	assert.Equal(t, "100.00em", Em(100))
	assert.Equal(t, "0.00em", Em(0))
	assert.Equal(t, "50.00em", Em(50))
	assert.Equal(t, "25.00em", Em(25))
}

func TestFloat(t *testing.T) {
	assert.Equal(t, "100.00", Float(100))
	assert.Equal(t, "0.00", Float(0))
	assert.Equal(t, "50.00", Float(50))
	assert.Equal(t, "25.00", Float(25))

	// Longer float
	assert.Equal(t, "100.12", Float(100.123456))
}

func TestInt(t *testing.T) {
	assert.Equal(t, "100", Int(100))
	assert.Equal(t, "0", Int(0))
	assert.Equal(t, "50", Int(50))
	assert.Equal(t, "25", Int(25))
}

func TestRem(t *testing.T) {
	assert.Equal(t, "100.00rem", Rem(100))
	assert.Equal(t, "0.00rem", Rem(0))
	assert.Equal(t, "50.00rem", Rem(50))
	assert.Equal(t, "50.00rem", Rem(50.0))
	assert.Equal(t, "50.00rem", Rem(50.00))
	assert.Equal(t, "25.00rem", Rem(25))
}

func TestPercent(t *testing.T) {
	assert.Equal(t, "100%", Percent(100))
	assert.Equal(t, "0%", Percent(0))
	assert.Equal(t, "50%", Percent(50))
	assert.Equal(t, "25%", Percent(25))
}

func TestPixels(t *testing.T) {
	assert.Equal(t, "100px", Pixels(100))
	assert.Equal(t, "0", Pixels(0))
	assert.Equal(t, "50px", Pixels(50))
	assert.Equal(t, "25px", Pixels(25))
}

func TestRGB(t *testing.T) {
	assert.Equal(t, "rgb(100,100,100)", RGB(100, 100, 100))
	assert.Equal(t, "rgb(0,0,0)", RGB(0, 0, 0))
	assert.Equal(t, "rgb(50,50,50)", RGB(50, 50, 50))
	assert.Equal(t, "rgb(25,25,25)", RGB(25, 25, 25))
}

func TestRGBA(t *testing.T) {
	assert.Equal(t, "rgba(100,100,100,100.0)", RGBA(100, 100, 100, 100))
	assert.Equal(t, "rgba(0,0,0,0.0)", RGBA(0, 0, 0, 0))
	assert.Equal(t, "rgba(50,50,50,0.5)", RGBA(50, 50, 50, 0.5))
	assert.Equal(t, "rgba(25,25,25,25.0)", RGBA(25, 25, 25, 25.0))
}

func TestURL(t *testing.T) {
	assert.Equal(t, "url('https://example.com')", URL("https://example.com"))
}

func TestVar(t *testing.T) {
	assert.Equal(t, "var(--primary-color)", Var("primary-color"))
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

func TestViewportMin(t *testing.T) {
	assert.Equal(t, "100vmin", ViewportMin(100))
	assert.Equal(t, "0vmin", ViewportMin(0))
	assert.Equal(t, "50vmin", ViewportMin(50))
	assert.Equal(t, "25vmin", ViewportMin(25))
}

func TestViewportMax(t *testing.T) {
	assert.Equal(t, "100vmax", ViewportMax(100))
	assert.Equal(t, "0vmax", ViewportMax(0))
	assert.Equal(t, "50vmax", ViewportMax(50))
	assert.Equal(t, "25vmax", ViewportMax(25))
}

func TestSeconds(t *testing.T) {
	assert.Equal(t, "100.00s", Seconds(100.00))
	assert.Equal(t, "1.25s", Seconds(1.25))
	assert.Equal(t, "0.00s", Seconds(0))
	assert.Equal(t, "0.25s", Seconds(0.25))
}

func TestMilliseconds(t *testing.T) {
	assert.Equal(t, "25ms", Milliseconds(25))
	assert.Equal(t, "1ms", Milliseconds(1))
	assert.Equal(t, "0ms", Milliseconds(0))
	assert.Equal(t, "110ms", Milliseconds(110))
}
