package stylemanager

import (
	"fmt"
	"github.com/chasefleming/elem-go/styles"
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
	style := Style{"color": "red"}
	className := sm.AddStyle(style)

	assert.NotEmpty(t, className)
	assert.Contains(t, sm.styles, className)
	assert.Equal(t, style, sm.styles[className])
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

func TestGenerateCSS(t *testing.T) {
	sm := NewStyleManager()

	// Define and add keyframes for an animation.
	keyframes := Keyframes{
		"from": {styles.Color: "red"},
		"to":   {styles.Color: "blue"},
	}
	animationName := sm.AddAnimation(keyframes)

	// Add a style that uses the animation.
	styleUsingAnimation := Style{
		"animation-name":     animationName,
		"animation-duration": "2s",
	}
	className := sm.AddStyle(styleUsingAnimation)

	// Add some additional basic styles.
	sm.AddStyle(Style{styles.Color: "red"})
	sm.AddStyle(Style{styles.Background: "blue"})

	// Generate the CSS.
	css := sm.GenerateCSS()

	// Assertions for basic styles.
	assert.Contains(t, css, "color: red", "CSS should contain color style")
	assert.Contains(t, css, "background: blue", "CSS should contain background style")

	// Assertions for animations.
	assert.Contains(t, css, fmt.Sprintf("@keyframes %s {", animationName), "CSS should contain the animation keyframes")
	assert.Contains(t, css, "from { color: red; }", "CSS should contain the 'from' keyframe")
	assert.Contains(t, css, "to { color: blue; }", "CSS should contain the 'to' keyframe")

	// Assertions for the usage of the animation in a style.
	assert.Contains(t, css, fmt.Sprintf(".%s {", className), "CSS should contain the class that uses the animation")
	assert.Contains(t, css, fmt.Sprintf("animation-name: %s;", animationName), "CSS should apply the animation name to the class")
	assert.Contains(t, css, "animation-duration: 2s;", "CSS should apply the animation duration to the class")
}

func TestAddStyleDeduplication(t *testing.T) {
	sm := NewStyleManager()
	style := Style{"color": "red"}
	firstClassName := sm.AddStyle(style)
	secondClassName := sm.AddStyle(style)

	assert.Equal(t, firstClassName, secondClassName, "Identical styles should return the same class name")
	assert.Len(t, sm.styles, 1, "StyleManager should not duplicate identical styles")
}
