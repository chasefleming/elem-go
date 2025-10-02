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

func TestClassNames(t *testing.T) {
	tests := []struct {
		name     string
		classes  []string
		expected string
	}{
		{
			name:     "multiple valid classes",
			classes:  []string{"btn", "btn-primary", "active"},
			expected: "btn btn-primary active",
		},
		{
			name:     "with empty strings",
			classes:  []string{"btn", "", "btn-primary"},
			expected: "btn btn-primary",
		},
		{
			name:     "with whitespace",
			classes:  []string{"btn", "  ", "btn-primary"},
			expected: "btn btn-primary",
		},
		{
			name:     "with leading/trailing spaces",
			classes:  []string{" btn ", "btn-primary  ", "  active"},
			expected: "btn btn-primary active",
		},
		{
			name:     "all empty",
			classes:  []string{"", "  ", ""},
			expected: "",
		},
		{
			name:     "single class",
			classes:  []string{"btn"},
			expected: "btn",
		},
		{
			name:     "no classes",
			classes:  []string{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClassNames(tt.classes...)
			assert.Equal(t, tt.expected, result)
		})
	}
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
