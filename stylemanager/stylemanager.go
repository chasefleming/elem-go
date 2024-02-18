package stylemanager

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

// Style represents CSS properties for an element.
type Style map[string]string

// Keyframes represents CSS keyframes for an animation.
type Keyframes map[string]Style

// CompositeStyle represents a collection of styles.
type CompositeStyle struct {
	Default       Style
	PseudoClasses map[string]Style
	MediaQueries  map[string]Style
}

// StyleSheet represents a collection of styles mapped to class names.
type StyleSheet map[string]Style

// StyleManager manages styles and generates CSS classes.
type StyleManager struct {
	styles          StyleSheet
	compositeStyles map[string]CompositeStyle
	animations      map[string]Keyframes
	mediaQueries    map[string]Style
}

// NewStyleManager creates a new instance of StyleManager.
func NewStyleManager() *StyleManager {
	return &StyleManager{
		styles:          make(StyleSheet),
		animations:      make(map[string]Keyframes),
		compositeStyles: make(map[string]CompositeStyle),
		mediaQueries:    make(map[string]Style),
	}
}

// AddStyle adds a new style to the manager and returns a class name.
func (sm *StyleManager) AddStyle(style Style) string {
	// Convert style to a string to generate hash.
	styleStr := fmt.Sprintf("%v", style)
	hash := sha1.Sum([]byte(styleStr))
	className := fmt.Sprintf("cls_%x", hash[:5]) // Use first 5 bytes of hash for class name.

	if _, exists := sm.styles[className]; !exists {
		sm.styles[className] = style
	}

	return className
}

// AddAnimation adds a new keyframes animation to the manager and returns an animation name.
func (sm *StyleManager) AddAnimation(keyframes Keyframes) string {
	// Convert keyframes to a string to generate hash.
	keyframesStr := fmt.Sprintf("%v", keyframes)
	hash := sha1.Sum([]byte(keyframesStr))
	animationName := fmt.Sprintf("anim_%x", hash[:5]) // Use first 5 bytes of hash for animation name.

	if _, exists := sm.animations[animationName]; !exists {
		sm.animations[animationName] = keyframes
	}

	return animationName
}

func (sm *StyleManager) AddCompositeStyle(composite CompositeStyle) string {
	// Convert composite to a string to generate hash.
	compositeStr := fmt.Sprintf("%v", composite)
	hash := sha1.Sum([]byte(compositeStr))
	className := fmt.Sprintf("cls_%x", hash[:5]) // Use first 5 bytes of hash for class name.

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
		for key, style := range keyframes {
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

// ensureMediaPrefix ensures that the mediaQuery starts with "@media".
func ensureMediaPrefix(mediaQuery string) string {
	if !strings.HasPrefix(mediaQuery, "@media") {
		return "@media " + mediaQuery
	}
	return mediaQuery
}

// sortedKeys returns the keys of the map sorted alphabetically.
func sortedKeys(style Style) []string {
	keys := make([]string, 0, len(style))
	for key := range style {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
