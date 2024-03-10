package attrs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataAttr(t *testing.T) {
	actual := DataAttr("foobar")
	expected := "data-foobar"
	assert.Equal(t, expected, actual)
}

func TestMerge(t *testing.T) {
	baseStyle := Props{
		"Width": "100px",
		"Color": "blue",
	}

	additionalStyle := Props{
		"Color":           "red", // This should override the blue color in baseStyle
		"BackgroundColor": "yellow",
	}

	expectedMergedStyle := Props{
		"Width":           "100px",
		"Color":           "red",
		"BackgroundColor": "yellow",
	}

	mergedStyle := Merge(baseStyle, additionalStyle)

	assert.Equal(t, expectedMergedStyle, mergedStyle)
}
