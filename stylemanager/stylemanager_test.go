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
	assert.Contains(t, sm.animations, exampleAnimation)
	assert.Equal(t, keyframes, sm.animations[exampleAnimation])
}

func TestAddCompositeStyle(t *testing.T) {
	sm := NewStyleManager()

	compositeStyle := CompositeStyle{
		Default: Style{styles.Color: "red"},
		PseudoClasses: map[string]Style{
			":hover": Style{styles.Color: "blue"},
		},
	}

	className := sm.AddCompositeStyle(compositeStyle)

	assert.NotEmpty(t, className)
	assert.Contains(t, sm.compositeStyles, className)
	assert.Equal(t, compositeStyle.Default, sm.compositeStyles[className].Default)
	assert.Equal(t, compositeStyle.PseudoClasses, sm.compositeStyles[className].PseudoClasses)
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
	styleOneClass := sm.AddStyle(Style{styles.Color: "red"})
	styleTwoClass := sm.AddStyle(Style{styles.Background: "blue"})

	// Add a composite style
	compositeClassName := sm.AddCompositeStyle(CompositeStyle{
		Default: Style{styles.Color: "pink"},
		PseudoClasses: map[string]Style{
			"hover": Style{styles.Color: "blue"},
		},
	})

	// Generate the CSS.
	css := sm.GenerateCSS()
	fmt.Println(css)

	// Assertions for basic styles.
	assert.Contains(t, css, fmt.Sprintf(".%s { color: red; }", styleOneClass), "CSS should contain style one")
	assert.Contains(t, css, fmt.Sprintf(".%s { background: blue; }", styleTwoClass), "CSS should contain style two")

	// Assertions for animations.
	assert.Contains(t, css, fmt.Sprintf("@keyframes %s { from { color: red; } to { color: blue; } }", animationName), "CSS should contain the keyframes for the animation")

	// Assertions for the usage of the animation in a style.
	assert.Contains(t, css, fmt.Sprintf(".%s { animation-duration: 2s; animation-name: %s; }", className, animationName), "CSS should contain the style using the animation")

	// Assertions for the usage of composite styles.
	assert.Contains(t, css, fmt.Sprintf(".%s { color: pink; }", compositeClassName), "CSS should contain the composite default")
	assert.Contains(t, css, fmt.Sprintf(".%s:hover { color: blue; }", compositeClassName), "CSS should contain the composite hover")
}

func TestAddStyleDeduplication(t *testing.T) {
	sm := NewStyleManager()
	style := Style{"color": "red"}
	firstClassName := sm.AddStyle(style)
	secondClassName := sm.AddStyle(style)

	assert.Equal(t, firstClassName, secondClassName, "Identical styles should return the same class name")
	assert.Len(t, sm.styles, 1, "StyleManager should not duplicate identical styles")
}
