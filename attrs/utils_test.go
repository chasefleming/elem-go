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

func TestClassNames(t *testing.T) {
	cases := []struct {
		name string
		in   []string
		want string
	}{
		{"none", nil, ""},
		{"single", []string{"btn"}, "btn"},
		{"two", []string{"btn", "btn-primary"}, "btn btn-primary"},
		{"skip empty", []string{"btn", "", "btn-primary"}, "btn btn-primary"},
		{"trim whitespace", []string{"  btn  ", "\tbtn-primary\n"}, "btn btn-primary"},
		{"skip whitespace-only", []string{"btn", "   ", "btn-primary"}, "btn btn-primary"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, ClassNames(tc.in...))
		})
	}
}
