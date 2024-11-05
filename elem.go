package elem

import (
	"fmt"
	"sort"
	"strings"

	"github.com/chasefleming/elem-go/attrs"
)

// List of HTML5 void elements. Void elements, also known as self-closing or empty elements,
// are elements that don't have a closing tag because they can't contain any content.
// For example, the <img> tag cannot wrap text or other tags, it stands alone, so it doesn't have a closing tag.
var voidElements = map[string]struct{}{
	"area":    {},
	"base":    {},
	"br":      {},
	"col":     {},
	"command": {},
	"embed":   {},
	"hr":      {},
	"img":     {},
	"input":   {},
	"keygen":  {},
	"link":    {},
	"meta":    {},
	"param":   {},
	"source":  {},
	"track":   {},
	"wbr":     {},
}

// List of boolean attributes. Boolean attributes can't have literal values. The presence of an boolean
// attribute represents the "true" value. To represent the "false" value, the attribute has to be omitted.
// See https://html.spec.whatwg.org/multipage/indices.html#attributes-3 for reference
var booleanAttrs = map[string]struct{}{
	attrs.AllowFullscreen: {},
	attrs.Async:           {},
	attrs.Autofocus:       {},
	attrs.Autoplay:        {},
	attrs.Checked:         {},
	attrs.Controls:        {},
	attrs.Defer:           {},
	attrs.Disabled:        {},
	attrs.Ismap:           {},
	attrs.Loop:            {},
	attrs.Multiple:        {},
	attrs.Muted:           {},
	attrs.Nomodule:        {},
	attrs.Novalidate:      {},
	attrs.Open:            {},
	attrs.Playsinline:     {},
	attrs.Readonly:        {},
	attrs.Required:        {},
	attrs.Selected:        {},
}

type CSSGenerator interface {
	GenerateCSS() string // TODO: Change to CSS()
}

type RenderOptions struct {
	// DisableHtmlPreamble disables the doctype preamble for the HTML tag if it exists in the rendering tree
	DisableHtmlPreamble bool
	StyleManager        CSSGenerator
}

type Node interface {
	RenderTo(builder *strings.Builder, opts RenderOptions)
	Render() string
	RenderWithOptions(opts RenderOptions) string
}

// NoneNode represents a node that renders nothing.
type NoneNode struct{}

// RenderTo for NoneNode does nothing.
func (n NoneNode) RenderTo(builder *strings.Builder, opts RenderOptions) {
	// Intentionally left blank to render nothing
}

// Render for NoneNode returns an empty string.
func (n NoneNode) Render() string {
	return ""
}

// RenderWithOptions for NoneNode returns an empty string.
func (n NoneNode) RenderWithOptions(opts RenderOptions) string {
	return ""
}

type TextNode string

func (t TextNode) RenderTo(builder *strings.Builder, opts RenderOptions) {
	builder.WriteString(string(t))
}

func (t TextNode) Render() string {
	return string(t)
}

func (t TextNode) RenderWithOptions(opts RenderOptions) string {
	return string(t)
}

type RawNode string

func (r RawNode) RenderTo(builder *strings.Builder, opts RenderOptions) {
	builder.WriteString(string(r))
}

func (r RawNode) Render() string {
	return string(r)
}

func (t RawNode) RenderWithOptions(opts RenderOptions) string {
	return string(t)
}

type CommentNode string

func (c CommentNode) RenderTo(builder *strings.Builder, opts RenderOptions) {
	builder.WriteString("<!-- ")
	builder.WriteString(string(c))
	builder.WriteString(" -->")
}

func (c CommentNode) Render() string {
	return c.RenderWithOptions(RenderOptions{})
}

func (c CommentNode) RenderWithOptions(opts RenderOptions) string {
	var builder strings.Builder
	c.RenderTo(&builder, opts)
	return builder.String()
}

type Element struct {
	Tag      string
	Attrs    attrs.Props
	Children []Node
}

func (e *Element) RenderTo(builder *strings.Builder, opts RenderOptions) {
	// The HTML tag needs a doctype preamble in order to ensure
	// browsers don't render in legacy/quirks mode
	// https://developer.mozilla.org/en-US/docs/Glossary/Doctype
	if !opts.DisableHtmlPreamble && e.Tag == "html" {
		builder.WriteString("<!DOCTYPE html>")
	}

	isFragment := e.Tag == "fragment"

	// Start with opening tag
	if !isFragment {
		builder.WriteString("<")
		builder.WriteString(e.Tag)
	}

	// Sort the keys for consistent order
	keys := make([]string, 0, len(e.Attrs))
	for k := range e.Attrs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Append the attributes to the builder
	for _, k := range keys {
		e.renderAttrTo(k, builder)
	}

	// If it's a void element, close it and return
	if _, exists := voidElements[e.Tag]; exists {
		builder.WriteString(`>`)
		return
	}

	if !isFragment {
		// Close opening tag
		builder.WriteString(`>`)
	}

	// Build the content
	for _, child := range e.Children {
		child.RenderTo(builder, opts)
	}

	if !isFragment {
		// Append closing tag
		builder.WriteString(`</`)
		builder.WriteString(e.Tag)
		builder.WriteString(`>`)
	}
}

// return string representation of given attribute with its value
func (e *Element) renderAttrTo(attrName string, builder *strings.Builder) {
	if _, exists := booleanAttrs[attrName]; exists {
		// boolean attribute presents its name only if the value is "true"
		if e.Attrs[attrName] == "true" {
			builder.WriteString(` `)
			builder.WriteString(attrName)
		}
	} else {
		// regular attribute has a name and a value
		attrVal := e.Attrs[attrName]

		// A necessary check to to avoid adding extra quotes around values that are already single-quoted
		// An example is '{"quantity": 5}'
		isSingleQuoted := strings.HasPrefix(attrVal, "'") && strings.HasSuffix(attrVal, "'")

		builder.WriteString(` `)
		builder.WriteString(attrName)
		builder.WriteString(`=`)
		if !isSingleQuoted {
			builder.WriteString(`"`)
		}
		builder.WriteString(attrVal)
		if !isSingleQuoted {
			builder.WriteString(`"`)
		}
	}
}

func (e *Element) Render() string {
	return e.RenderWithOptions(RenderOptions{})
}

func (e *Element) RenderWithOptions(opts RenderOptions) string {
	var builder strings.Builder
	e.RenderTo(&builder, opts)

	if opts.StyleManager != nil {
		htmlContent := builder.String()
		cssContent := opts.StyleManager.GenerateCSS()

		// Define the <style> element with the generated CSS content
		styleElement := fmt.Sprintf("<style>%s</style>", cssContent)

		// Check if a <head> tag exists in the HTML content
		headStartIndex := strings.Index(htmlContent, "<head>")
		headEndIndex := strings.Index(htmlContent, "</head>")

		if headStartIndex != -1 && headEndIndex != -1 {
			// If <head> exists, inject the style content just before </head>
			beforeHead := htmlContent[:headEndIndex]
			afterHead := htmlContent[headEndIndex:]
			modifiedHTML := beforeHead + styleElement + afterHead
			return modifiedHTML
		} else {
			// If <head> does not exist, create it and inject the style content
			// Assuming <html> tag exists and injecting <head> immediately after <html>
			htmlTagEnd := strings.Index(htmlContent, ">") + 1
			if htmlTagEnd > 0 {
				beforeHTML := htmlContent[:htmlTagEnd]
				afterHTML := htmlContent[htmlTagEnd:]
				modifiedHTML := beforeHTML + "<head>" + styleElement + "</head>" + afterHTML
				return modifiedHTML
			}
		}
	}

	// Return the original HTML content if no modifications were made
	return builder.String()
}

func newElement(tag string, attrs attrs.Props, children ...Node) *Element {
	return &Element{
		Tag:      tag,
		Attrs:    attrs,
		Children: children,
	}
}
