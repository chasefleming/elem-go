package elem

import (
	"regexp"
	"strings"
)

// nodeContentReplacer handles escaping of HTML special characters in text nodes
// note that single and double quotes (', ") do not need escaping in HTML5 text nodes
var nodeContentReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	// technically, > does not need escaping in HTML5, but we do it for consistency
	">", "&gt;",
)

// commentContentsReplacer handles escaping of disallowed sequences in HTML comments
var commentContentsReplacer = strings.NewReplacer(
	"<!--", "&lt;!--",
	"-->", "--&gt;",
	"--!>", "--!&gt;",
)

// If conditionally renders one of the provided elements based on the condition
func If[T any](condition bool, ifTrue, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// TransformEach maps a slice of items to a slice of Elements using the provided function
func TransformEach[T any](items []T, fn func(T) Node) []Node {
	var nodes []Node
	for _, item := range items {
		nodes = append(nodes, fn(item))
	}
	return nodes
}

// EscapeNodeContents escapes HTML5 special characters in a string to ensure safe rendering as a text node
func EscapeNodeContents(s string) string {
	return nodeContentReplacer.Replace(s)
}

// EscapeCdataContents escapes the contents of a CDATA section to ensure safe rendering
func EscapeCdataContents(s string) string {
	// text in cdata must not contain the string "]]>"
	if strings.Contains(s, "]]>") {
		s = strings.ReplaceAll(s, "]]>", "]]&gt;")
	}
	return s
}

// EscapeCommentContents escapes the contents of a comment node to ensure safe rendering according to https://html.spec.whatwg.org/multipage/syntax.html#comments
func EscapeCommentContents(s string) string {
	s = commentContentsReplacer.Replace(s)

	// comments cannot start with specific character sequences
	if strings.HasPrefix(s, "->") {
		s = "-&gt;" + s[2:]
	} else if strings.HasPrefix(s, ">") {
		s = "&gt;" + s[1:]
	}

	// comments cannot end with specific character sequences
	if strings.HasSuffix(s, "<!-") {
		s = s[:len(s)-3] + "&lt;!-"
	}

	return s
}

// scriptContentsRegexp matches <script and </script case-insensitively
var scriptContentsRegexp = regexp.MustCompile(`(?i)<(/?)(script)`)

// EscapeScriptContents escapes the contents of a script node according to https://html.spec.whatwg.org/multipage/scripting.html#restrictions-for-contents-of-script-elements
func EscapeScriptContents(s string) string {
	s = strings.ReplaceAll(s, `<!--`, `\x3C!--`)
	s = scriptContentsRegexp.ReplaceAllString(s, `\x3C$1$2`)
	return s
}
