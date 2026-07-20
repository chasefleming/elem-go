package styles

import (
	"fmt"
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
	style := Props{"color": "red"}
	className := sm.AddStyle(style)

	assert.NotEmpty(t, className)
	assert.Contains(t, sm.styles, className)
	assert.Equal(t, style, sm.styles[className])
}

func TestAnimationKeyframes(t *testing.T) {
	sm := NewStyleManager()
	keyframes := Keyframes{
		"from": Props{Color: "red"},
		"to":   Props{Color: "blue"},
	}
	exampleAnimation := sm.AddAnimation(keyframes)

	assert.NotEmpty(t, exampleAnimation)
	assert.Contains(t, sm.animations, exampleAnimation)
	assert.Equal(t, keyframes, sm.animations[exampleAnimation])
}

func TestAddCompositeStyle(t *testing.T) {
	sm := NewStyleManager()

	compositeStyle := CompositeStyle{
		Default: Props{Color: "red"},
		PseudoClasses: map[string]Props{
			":hover": Props{Color: "blue"},
		},
	}

	className := sm.AddCompositeStyle(compositeStyle)

	assert.NotEmpty(t, className)
	assert.Contains(t, sm.compositeStyles, className)
	assert.Equal(t, compositeStyle.Default, sm.compositeStyles[className].Default)
	assert.Equal(t, compositeStyle.PseudoClasses, sm.compositeStyles[className].PseudoClasses)
}

func TestMediaQuery(t *testing.T) {
	sm := NewStyleManager()

	// Define a media query for tablets
	tabletQuery := "@media (min-width: 768px) and (max-width: 1024px)"

	compositeStyle := CompositeStyle{
		Default: Props{
			"padding": "10px",
		},
		MediaQueries: map[string]Props{
			tabletQuery: Props{
				"padding": "20px",
			},
		},
	}

	compositeClassName := sm.AddCompositeStyle(compositeStyle)

	assert.NotEmpty(t, compositeClassName)
	assert.Contains(t, sm.compositeStyles, compositeClassName)
	assert.Equal(t, compositeStyle.Default, sm.compositeStyles[compositeClassName].Default)
	assert.Equal(t, compositeStyle.MediaQueries, sm.compositeStyles[compositeClassName].MediaQueries)
}

func TestGenerateCSS(t *testing.T) {
	sm := NewStyleManager()

	// Define and add keyframes for an animation.
	keyframes := Keyframes{
		"from": {"color": "red"},
		"to":   {"color": "blue"},
	}
	animationName := sm.AddAnimation(keyframes)

	// Add a style that uses the animation.
	styleUsingAnimation := Props{
		"animation-name":     animationName,
		"animation-duration": "2s",
	}
	className := sm.AddStyle(styleUsingAnimation)

	// Add some additional basic styles.
	styleOneClass := sm.AddStyle(Props{"color": "red"})
	styleTwoClass := sm.AddStyle(Props{"background": "blue"})

	// Add a composite style
	compositeClassName := sm.AddCompositeStyle(CompositeStyle{
		Default: Props{"color": "pink"},
		PseudoClasses: map[string]Props{
			"hover": Props{"color": "blue"},
		},
		PseudoElements: map[string]Props{
			"before": Props{"content": "'[x] '"},
		},
		MediaQueries: map[string]Props{
			"@media (min-width: 768px)":  Props{"color": "green", "background": "yellow"},
			"@media (min-width: 1024px)": Props{"color": "purple", "background": "orange"},
		},
	})

	// Generate the CSS.
	css := sm.GenerateCSS()

	// Assertions for basic styles.
	assert.Contains(t, css, fmt.Sprintf(".%s { color: red; }", styleOneClass), "CSS should contain style one")
	assert.Contains(t, css, fmt.Sprintf(".%s { background: blue; }", styleTwoClass), "CSS should contain style two")

	// Assertions for animations.
	assert.Contains(t, css, fmt.Sprintf("@keyframes %s", animationName), "CSS should contain the animation")
	assert.Contains(t, css, "from { color: red; }", "CSS should contain the 'from' keyframe")
	assert.Contains(t, css, "to { color: blue; }", "CSS should contain the 'to' keyframe")
	// Assertions for the usage of the animation in a style.
	assert.Contains(t, css, fmt.Sprintf(".%s { animation-duration: 2s; animation-name: %s; }", className, animationName), "CSS should contain the style using the animation")

	// Assertions for the usage of composite styles.
	assert.Contains(t, css, fmt.Sprintf(".%s { color: pink; }", compositeClassName), "CSS should contain the composite default")
	assert.Contains(t, css, fmt.Sprintf(".%s:hover { color: blue; }", compositeClassName), "CSS should contain the composite hover")

	// Assertion for the usage of pseudo elements.
	assert.Contains(t, css, fmt.Sprintf(".%s::before { content: '[x] '; }", compositeClassName), "CSS should contain the composite before")

	// Assertions for media queries.
	assert.Contains(t, css, fmt.Sprintf("@media (min-width: 768px) { .%s { background: yellow; color: green; } }", compositeClassName), "CSS should contain the first media query")
	assert.Contains(t, css, fmt.Sprintf("@media (min-width: 1024px) { .%s { background: orange; color: purple; } }", compositeClassName), "CSS should contain the second media query")
}

func TestAddStyleDeduplication(t *testing.T) {
	sm := NewStyleManager()
	style := Props{"color": "red"}
	firstClassName := sm.AddStyle(style)
	secondClassName := sm.AddStyle(style)

	assert.Equal(t, firstClassName, secondClassName, "Identical styles should return the same class name")
	assert.Len(t, sm.styles, 1, "StyleManager should not duplicate identical styles")
}
