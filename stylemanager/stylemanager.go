package stylemanager

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// Style represents CSS properties for an element.
type Style map[string]string

// StyleSheet represents a collection of styles mapped to class names.
type StyleSheet map[string]Style

// StyleManager manages styles and generates CSS classes.
type StyleManager struct {
	styles     StyleSheet
	classCache map[string]string
}

// NewStyleManager creates a new instance of StyleManager.
func NewStyleManager() *StyleManager {
	return &StyleManager{
		styles:     make(StyleSheet),
		classCache: make(map[string]string),
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

	sm.classCache[styleStr] = className

	return className
}

// GenerateCSS generates the CSS string for all styles managed by StyleManager.
func (sm *StyleManager) GenerateCSS() string {
	var builder strings.Builder
	for className, style := range sm.styles {
		builder.WriteString(fmt.Sprintf(".%s { ", className))
		for prop, value := range style {
			builder.WriteString(fmt.Sprintf("%s: %s; ", prop, value))
		}
		builder.WriteString("} ")
	}
	return builder.String()
}
