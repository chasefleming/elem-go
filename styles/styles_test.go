package styles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyleToInline(t *testing.T) {
	style := Props{
		BackgroundColor: "blue",
		Color:           "white",
		FontSize:        "16px",
	}

	assert.Equal(t, "background-color: blue; color: white; font-size: 16px;", style.ToInline())
}

func TestStyleString_Empty(t *testing.T) {
	style := Props{}
	assert.Equal(t, "", style.ToInline())
}

func TestStyleString_SinglePair(t *testing.T) {
	style := Props{
		Color: "red",
	}
	assert.Equal(t, "color: red;", style.ToInline())
}

func TestStyleString_UnorderedKeys(t *testing.T) {
	style := Props{
		BackgroundColor: "blue",
		Color:           "white",
		FontSize:        "16px",
		OutlineStyle:    "solid",
	}
	assert.Equal(t, "background-color: blue; color: white; font-size: 16px; outline-style: solid;", style.ToInline())
}
