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

func TestMergeDataAttrs(t *testing.T) {
	dataAttrs := map[string]string{
		"user-id":   "123",
		"user-name": "john",
		"role":      "admin",
	}

	result := MergeDataAttrs(dataAttrs)

	expected := Props{
		"data-user-id":   "123",
		"data-user-name": "john",
		"data-role":      "admin",
	}

	assert.Equal(t, expected, result)
}

func TestMergeDataAttrs_Empty(t *testing.T) {
	result := MergeDataAttrs(map[string]string{})
	assert.Empty(t, result)
}
