package stylemanager

import (
	"github.com/chasefleming/elem-go/styles"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStyleManager(t *testing.T) {
	sm := NewStyleManager()
	assert.NotNil(t, sm)
	assert.Empty(t, sm.styles)
}

func TestAddStyle(t *testing.T) {
	sm := NewStyleManager()
	style := Style{styles.Color: "red"}
	className := sm.AddStyle(style)

	assert.NotEmpty(t, className)
	assert.Contains(t, sm.styles, className)
	assert.Equal(t, style, sm.styles[className])
}

func TestGenerateCSS(t *testing.T) {
	sm := NewStyleManager()
	sm.AddStyle(Style{styles.Color: "red"})
	sm.AddStyle(Style{styles.Background: "blue"})

	css := sm.GenerateCSS()

	assert.Contains(t, css, "color: red")
	assert.Contains(t, css, "background: blue")
	assert.True(t, strings.HasSuffix(css, "} "), "CSS should end with } ")
}

func TestAnimationKeyframes(t *testing.T) {
	sm := NewStyleManager()
	keyframes := Keyframes{
		"from": Style{styles.Color: "red"},
		"to":   Style{styles.Color: "blue"},
	}
	exampleAnimation := sm.AddAnimation(keyframes)

	assert.NotEmpty(t, exampleAnimation)
	assert.Contains(t, sm.animationCache, exampleAnimation)
	assert.Equal(t, keyframes, sm.animationCache[exampleAnimation])
}

func TestAddStyleDeduplication(t *testing.T) {
	sm := NewStyleManager()
	style := Style{"color": "red"}
	firstClassName := sm.AddStyle(style)
	secondClassName := sm.AddStyle(style)

	assert.Equal(t, firstClassName, secondClassName, "Identical styles should return the same class name")
	assert.Len(t, sm.styles, 1, "StyleManager should not duplicate identical styles")
}
