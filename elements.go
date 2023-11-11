package elem

import (
	"github.com/chasefleming/elem-go/attrs"
)

// ========== Document Structure ==========

// Body creates a <body> element.
func Body(attrs attrs.Props, children ...Node) *Element {
	return NewElement("body", attrs, children...)
}

// Head creates a <head> element.
func Head(attrs attrs.Props, children ...Node) *Element {
	return NewElement("head", attrs, children...)
}

// Html creates an <html> element.
func Html(attrs attrs.Props, children ...Node) *Element {
	return NewElement("html", attrs, children...)
}

// Title creates a <title> element.
func Title(attrs attrs.Props, children ...Node) *Element {
	return NewElement("title", attrs, children...)
}

// ========== Text Formatting and Structure ==========

// A creates an <a> element.
func A(attrs attrs.Props, children ...Node) *Element {
	return NewElement("a", attrs, children...)
}

// Br creates a <br> element.
func Br(attrs attrs.Props) *Element {
	return NewElement("br", attrs)
}

// Blockquote creates a <blockquote> element.
func Blockquote(attrs attrs.Props, children ...Node) *Element {
	return NewElement("blockquote", attrs, children...)
}

// Code creates a <code> element.
func Code(attrs attrs.Props, children ...Node) *Element {
	return NewElement("code", attrs, children...)
}

// Div creates a <div> element.
func Div(attrs attrs.Props, children ...Node) *Element {
	return NewElement("div", attrs, children...)
}

// Em creates an <em> element.
func Em(attrs attrs.Props, children ...Node) *Element {
	return NewElement("em", attrs, children...)
}

// H1 creates an <h1> element.
func H1(attrs attrs.Props, children ...Node) *Element {
	return NewElement("h1", attrs, children...)
}

// H2 creates an <h2> element.
func H2(attrs attrs.Props, children ...Node) *Element {
	return NewElement("h2", attrs, children...)
}

// H3 creates an <h3> element.
func H3(attrs attrs.Props, children ...Node) *Element {
	return NewElement("h3", attrs, children...)
}

// H4 creates an <h4> element.
func Hr(attrs attrs.Props) *Element {
	return NewElement("hr", attrs)
}

// I creates an <i> element.
func I(attrs attrs.Props, children ...Node) *Element {
	return NewElement("i", attrs, children...)
}

// P creates a <p> element.
func P(attrs attrs.Props, children ...Node) *Element {
	return NewElement("p", attrs, children...)
}

// Pre creates a <pre> element.
func Pre(attrs attrs.Props, children ...Node) *Element {
	return NewElement("pre", attrs, children...)
}

// Span creates a <span> element.
func Span(attrs attrs.Props, children ...Node) *Element {
	return NewElement("span", attrs, children...)
}

// Strong creates a <strong> element.
func Strong(attrs attrs.Props, children ...Node) *Element {
	return NewElement("strong", attrs, children...)
}

// Text creates a TextNode.
func Text(content string) TextNode {
	return TextNode(content)
}

// Comment creates a CommentNode.
func Comment(comment string) CommentNode {
	return CommentNode(comment)
}

// ========== Lists ==========

// Li creates an <li> element.
func Li(attrs attrs.Props, children ...Node) *Element {
	return NewElement("li", attrs, children...)
}

// Ul creates a <ul> element.
func Ul(attrs attrs.Props, children ...Node) *Element {
	return NewElement("ul", attrs, children...)
}

// Ol creates an <ol> element.
func Ol(attrs attrs.Props, children ...Node) *Element {
	return NewElement("ol", attrs, children...)
}

// Dl creates a <dl> element.
func Dl(attrs attrs.Props, children ...Node) *Element {
	return NewElement("dl", attrs, children...)
}

// Dt creates a <dt> element.
func Dt(attrs attrs.Props, children ...Node) *Element {
	return NewElement("dt", attrs, children...)
}

// Dd creates a <dd> element.
func Dd(attrs attrs.Props, children ...Node) *Element {
	return NewElement("dd", attrs, children...)
}

// ========== Forms ==========

// Button creates a <button> element.
func Button(attrs attrs.Props, children ...Node) *Element {
	return NewElement("button", attrs, children...)
}

// Form creates a <form> element.
func Form(attrs attrs.Props, children ...Node) *Element {
	return NewElement("form", attrs, children...)
}

// Input creates an <input> element.
func Input(attrs attrs.Props) *Element {
	return NewElement("input", attrs)
}

// Label creates a <label> element.
func Label(attrs attrs.Props, children ...Node) *Element {
	return NewElement("label", attrs, children...)
}

// Option creates an <option> element.
func Option(attrs attrs.Props, content TextNode) *Element {
	return NewElement("option", attrs, content)
}

// Select creates a <select> element.
func Select(attrs attrs.Props, children ...Node) *Element {
	return NewElement("select", attrs, children...)
}

// Textarea creates a <textarea> element.
func Textarea(attrs attrs.Props, content TextNode) *Element {
	return NewElement("textarea", attrs, content)
}

// ========== Hyperlinks and Multimedia ==========

// Img creates an <img> element.
func Img(attrs attrs.Props) *Element {
	return NewElement("img", attrs)
}

// ========== Meta Elements ==========

// Link creates a <link> element.
func Link(attrs attrs.Props) *Element {
	return NewElement("link", attrs)
}

// Meta creates a <meta> element.
func Meta(attrs attrs.Props) *Element {
	return NewElement("meta", attrs)
}

// Script creates a <script> element.
func Script(attrs attrs.Props, children ...Node) *Element {
	return NewElement("script", attrs, children...)
}

// Style creates a <style> element.
func Style(attrs attrs.Props, children ...Node) *Element {
	return NewElement("style", attrs, children...)
}

// ========== Semantic Elements ==========

// --- Semantic Sectioning Elements ---

// Article creates an <article> element.
func Article(attrs attrs.Props, children ...Node) *Element {
	return NewElement("article", attrs, children...)
}

// Aside creates an <aside> element.
func Aside(attrs attrs.Props, children ...Node) *Element {
	return NewElement("aside", attrs, children...)
}

// Footer creates a <footer> element.
func Footer(attrs attrs.Props, children ...Node) *Element {
	return NewElement("footer", attrs, children...)
}

// Header creates a <header> element.
func Header(attrs attrs.Props, children ...Node) *Element {
	return NewElement("header", attrs, children...)
}

// Main creates a <main> element.
func Main(attrs attrs.Props, children ...Node) *Element {
	return NewElement("main", attrs, children...)
}

// Nav creates a <nav> element.
func Nav(attrs attrs.Props, children ...Node) *Element {
	return NewElement("nav", attrs, children...)
}

// Section creates a <section> element.
func Section(attrs attrs.Props, children ...Node) *Element {
	return NewElement("section", attrs, children...)
}

// ========== Semantic Form Elements ==========

// Fieldset creates a <fieldset> element.
func Fieldset(attrs attrs.Props, children ...Node) *Element {
	return NewElement("fieldset", attrs, children...)
}

// Legend creates a <legend> element.
func Legend(attrs attrs.Props, children ...Node) *Element {
	return NewElement("legend", attrs, children...)
}

// Datalist creates a <datalist> element.
func Datalist(attrs attrs.Props, children ...Node) *Element {
	return NewElement("datalist", attrs, children...)
}

// Meter creates a <meter> element.
func Meter(attrs attrs.Props, children ...Node) *Element {
	return NewElement("meter", attrs, children...)
}

// Output creates an <output> element.
func Output(attrs attrs.Props, children ...Node) *Element {
	return NewElement("output", attrs, children...)
}

// Progress creates a <progress> element.
func Progress(attrs attrs.Props, children ...Node) *Element {
	return NewElement("progress", attrs, children...)
}

// --- Semantic Interactive Elements ---

// Dialog creates a <dialog> element.
func Dialog(attrs attrs.Props, children ...Node) *Element {
	return NewElement("dialog", attrs, children...)
}

// Menu creates a <menu> element.
func Menu(attrs attrs.Props, children ...Node) *Element {
	return NewElement("menu", attrs, children...)
}

// --- Semantic Script Supporting Elements ---

// NoScript creates a <noscript> element.
func NoScript(attrs attrs.Props, children ...Node) *Element {
	return NewElement("noscript", attrs, children...)
}

// --- Semantic Text Content Elements ---

// Address creates an <address> element.
func Address(attrs attrs.Props, children ...Node) *Element {
	return NewElement("address", attrs, children...)
}

// Details creates a <details> element.
func Details(attrs attrs.Props, children ...Node) *Element {
	return NewElement("details", attrs, children...)
}

// FigCaption creates a <figcaption> element.
func FigCaption(attrs attrs.Props, children ...Node) *Element {
	return NewElement("figcaption", attrs, children...)
}

// Figure creates a <figure> element.
func Figure(attrs attrs.Props, children ...Node) *Element {
	return NewElement("figure", attrs, children...)
}

// Mark creates a <mark> element.
func Mark(attrs attrs.Props, children ...Node) *Element {
	return NewElement("mark", attrs, children...)
}

// Summary creates a <summary> element.
func Summary(attrs attrs.Props, children ...Node) *Element {
	return NewElement("summary", attrs, children...)
}

// Time creates a <time> element.
func Time(attrs attrs.Props, children ...Node) *Element {
	return NewElement("time", attrs, children...)
}

// ========== Tables ==========

// Table creates a <table> element.
func Table(attrs attrs.Props, children ...Node) *Element {
	return NewElement("table", attrs, children...)
}

// THead creates a <thead> element.
func THead(attrs attrs.Props, children ...Node) *Element {
	return NewElement("thead", attrs, children...)
}

// TBody creates a <tbody> element.
func TBody(attrs attrs.Props, children ...Node) *Element {
	return NewElement("tbody", attrs, children...)
}

// TFoot creates a <tfoot> element.
func TFoot(attrs attrs.Props, children ...Node) *Element {
	return NewElement("tfoot", attrs, children...)
}

// Tr creates a <tr> element.
func Tr(attrs attrs.Props, children ...Node) *Element {
	return NewElement("tr", attrs, children...)
}

// Th creates a <th> element.
func Th(attrs attrs.Props, children ...Node) *Element {
	return NewElement("th", attrs, children...)
}

// Td creates a <td> element.
func Td(attrs attrs.Props, children ...Node) *Element {
	return NewElement("td", attrs, children...)
}

// ========== Embedded Content ==========

// IFrames creates an <iframe> element.
func IFrame(attrs attrs.Props, children ...Node) *Element {
	return NewElement("iframe", attrs, children...)
}

// Audio creates an <audio> element.
func Audio(attrs attrs.Props, children ...Node) *Element {
	return NewElement("audio", attrs, children...)
}

// Video creates a <video> element.
func Video(attrs attrs.Props, children ...Node) *Element {
	return NewElement("video", attrs, children...)
}

// Source creates a <source> element.
func Source(attrs attrs.Props, children ...Node) *Element {
	return NewElement("source", attrs, children...)
}
