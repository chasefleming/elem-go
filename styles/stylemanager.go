package styles

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

// Keyframes represents CSS keyframes for an animation.
type Keyframes map[string]Props

// CompositeStyle represents a collection of styles.
type CompositeStyle struct {
	Default       Props
	PseudoClasses map[string]Props
	PseudoElements map[string]Props
	MediaQueries  map[string]Props
}

// StyleSheet represents a collection of styles mapped to class names.
type StyleSheet map[string]Props

// StyleManager manages styles and generates CSS classes.
type StyleManager struct {
	styles          StyleSheet
	compositeStyles map[string]CompositeStyle
	animations      map[string]Keyframes
	mediaQueries    map[string]Props
}

// NewStyleManager creates a new instance of StyleManager.
func NewStyleManager() *StyleManager {
	return &StyleManager{
		styles:          make(StyleSheet),
		animations:      make(map[string]Keyframes),
		compositeStyles: make(map[string]CompositeStyle),
		mediaQueries:    make(map[string]Props),
	}
}

// AddStyle adds a new style to the manager and returns a class name.
func (sm *StyleManager) AddStyle(style Props) string {
	className := fmt.Sprintf("cls_%x", entityHash(style))

	if _, exists := sm.styles[className]; !exists {
		sm.styles[className] = style
	}

	return className
}

// AddAnimation adds a new keyframes animation to the manager and returns an animation name.
func (sm *StyleManager) AddAnimation(keyframes Keyframes) string {
	animationName := fmt.Sprintf("anim_%x", entityHash(keyframes))

	if _, exists := sm.animations[animationName]; !exists {
		sm.animations[animationName] = keyframes
	}

	return animationName
}

func (sm *StyleManager) AddCompositeStyle(composite CompositeStyle) string {
	className := fmt.Sprintf("cls_%x", entityHash(composite))

	if _, exists := sm.compositeStyles[className]; !exists {
		sm.compositeStyles[className] = composite
	}

	return className
}

// GenerateCSS generates the CSS string for all styles managed by StyleManager.
func (sm *StyleManager) GenerateCSS() string {
	var builder strings.Builder

	for className, style := range sm.styles {
		keys := sortedKeys(style)
		builder.WriteString(fmt.Sprintf(".%s { ", className))
		for _, prop := range keys {
			builder.WriteString(fmt.Sprintf("%s: %s; ", prop, style[prop]))
		}
		builder.WriteString("} ")
	}

	for animationName, keyframes := range sm.animations {
		builder.WriteString(fmt.Sprintf("@keyframes %s { ", animationName))
		keys := sortedKeys(keyframes)
		for _, key := range keys {
			style := keyframes[key]
			builder.WriteString(fmt.Sprintf("%s { ", key))
			for prop, value := range style {
				builder.WriteString(fmt.Sprintf("%s: %s; ", prop, value))
			}
			builder.WriteString("} ")
		}
		builder.WriteString("} ")
	}

	for className, composite := range sm.compositeStyles {
		keys := sortedKeys(composite.Default)
		builder.WriteString(fmt.Sprintf(".%s { ", className))
		for _, prop := range keys {
			builder.WriteString(fmt.Sprintf("%s: %s; ", prop, composite.Default[prop]))
		}
		builder.WriteString("} ")

		for pseudoClass, style := range composite.PseudoClasses {
			// Ensure pseudoClass starts with a colon
			formattedPseudoClass := fmt.Sprintf("%s%s", className, ensureLeadingColon(pseudoClass))
			keys := sortedKeys(style)
			builder.WriteString(fmt.Sprintf(".%s { ", formattedPseudoClass))
			for _, prop := range keys {
				builder.WriteString(fmt.Sprintf("%s: %s; ", prop, style[prop]))
			}
			builder.WriteString("} ")
		}

		for pseudoElement, style := range composite.PseudoElements {
			// Ensure pseudoElement starts with a double colon
			formattedPseudoElement := fmt.Sprintf("%s%s", className, ensureDoubleLeadingColon(pseudoElement))
			keys := sortedKeys(style)
			builder.WriteString(fmt.Sprintf(".%s { ", formattedPseudoElement))
			for _, prop := range keys {
				builder.WriteString(fmt.Sprintf("%s: %s; ", prop, style[prop]))
			}
			builder.WriteString("} ")
		}

		for mediaQuery, style := range composite.MediaQueries {
			// Ensure mediaQuery is correctly prefixed
			formattedMediaQuery := ensureMediaPrefix(mediaQuery)
			builder.WriteString(fmt.Sprintf("%s { .%s { ", formattedMediaQuery, className))
			keys := sortedKeys(style)
			for _, prop := range keys {
				builder.WriteString(fmt.Sprintf("%s: %s; ", prop, style[prop]))
			}
			builder.WriteString("} } ")
		}
	}

	return builder.String()
}

// ensureLeadingColon ensures that the pseudoClass starts with a colon.
func ensureLeadingColon(pseudoClass string) string {
	if strings.HasPrefix(pseudoClass, ":") {
		return pseudoClass
	}
	return ":" + pseudoClass
}

// ensureDoubleLeadingColon ensures that the pseudoElement starts with a double colon.
func ensureDoubleLeadingColon(pseudoElement string) string {
	if strings.HasPrefix(pseudoElement, "::") {
		return pseudoElement
	}
	return "::" + pseudoElement
}

// ensureMediaPrefix ensures that the mediaQuery starts with "@media".
func ensureMediaPrefix(mediaQuery string) string {
	if !strings.HasPrefix(mediaQuery, "@media") {
		return "@media " + mediaQuery
	}
	return mediaQuery
}

// sortedKeys returns the keys of the map sorted alphanumerically
func sortedKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// entityHash generates a deterministic unique hash for CSS entities
// While result is deterministic, it is not guaranteed to be the same in future versions of the package
func entityHash[T any](x T) []byte {
	styleStr := fmt.Sprintf("%v", x) // since go1.12 this is deterministic
	hash := sha1.Sum([]byte(styleStr))
	// Use first 5 bytes of hash for brevity. The probability that these short hashes collide is very low
	return hash[:5]
}
