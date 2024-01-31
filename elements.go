package elem

import (
	"github.com/chasefleming/elem-go/attrs"
)

// ========== Document Structure ==========

// Body creates a <body> element.
func Body(attrs attrs.Props, children ...Node) *Element {
	return newElement("body", attrs, children...)
}

// Head creates a <head> element.
func Head(attrs attrs.Props, children ...Node) *Element {
	return newElement("head", attrs, children...)
}

// Html creates an <html> element.
func Html(attrs attrs.Props, children ...Node) *Element {
	return newElement("html", attrs, children...)
}

// Title creates a <title> element.
func Title(attrs attrs.Props, children ...Node) *Element {
	return newElement("title", attrs, children...)
}

// ========== Text Formatting and Structure ==========

// A creates an <a> element.
func A(attrs attrs.Props, children ...Node) *Element {
	return newElement("a", attrs, children...)
}

// Br creates a <br> element.
func Br(attrs attrs.Props) *Element {
	return newElement("br", attrs)
}

// Blockquote creates a <blockquote> element.
func Blockquote(attrs attrs.Props, children ...Node) *Element {
	return newElement("blockquote", attrs, children...)
}

// Code creates a <code> element.
func Code(attrs attrs.Props, children ...Node) *Element {
	return newElement("code", attrs, children...)
}

// Div creates a <div> element.
func Div(attrs attrs.Props, children ...Node) *Element {
	return newElement("div", attrs, children...)
}

// Em creates an <em> element.
func Em(attrs attrs.Props, children ...Node) *Element {
	return newElement("em", attrs, children...)
}

// H1 creates an <h1> element.
func H1(attrs attrs.Props, children ...Node) *Element {
	return newElement("h1", attrs, children...)
}

// H2 creates an <h2> element.
func H2(attrs attrs.Props, children ...Node) *Element {
	return newElement("h2", attrs, children...)
}

// H3 creates an <h3> element.
func H3(attrs attrs.Props, children ...Node) *Element {
	return newElement("h3", attrs, children...)
}

// H4 creates an <h4> element.
func H4(attrs attrs.Props, children ...Node) *Element {
	return newElement("h4", attrs, children...)
}

// H5 creates an <h5> element.
func H5(attrs attrs.Props, children ...Node) *Element {
	return newElement("h5", attrs, children...)
}

// H6 creates an <h6> element.
func H6(attrs attrs.Props, children ...Node) *Element {
	return newElement("h6", attrs, children...)
}

// Hr creates an <hr> element.
func Hr(attrs attrs.Props) *Element {
	return newElement("hr", attrs)
}

// I creates an <i> element.
func I(attrs attrs.Props, children ...Node) *Element {
	return newElement("i", attrs, children...)
}

// P creates a <p> element.
func P(attrs attrs.Props, children ...Node) *Element {
	return newElement("p", attrs, children...)
}

// Pre creates a <pre> element.
func Pre(attrs attrs.Props, children ...Node) *Element {
	return newElement("pre", attrs, children...)
}

// Span creates a <span> element.
func Span(attrs attrs.Props, children ...Node) *Element {
	return newElement("span", attrs, children...)
}

// Strong creates a <strong> element.
func Strong(attrs attrs.Props, children ...Node) *Element {
	return newElement("strong", attrs, children...)
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
	return newElement("li", attrs, children...)
}

// Ul creates a <ul> element.
func Ul(attrs attrs.Props, children ...Node) *Element {
	return newElement("ul", attrs, children...)
}

// Ol creates an <ol> element.
func Ol(attrs attrs.Props, children ...Node) *Element {
	return newElement("ol", attrs, children...)
}

// Dl creates a <dl> element.
func Dl(attrs attrs.Props, children ...Node) *Element {
	return newElement("dl", attrs, children...)
}

// Dt creates a <dt> element.
func Dt(attrs attrs.Props, children ...Node) *Element {
	return newElement("dt", attrs, children...)
}

// Dd creates a <dd> element.
func Dd(attrs attrs.Props, children ...Node) *Element {
	return newElement("dd", attrs, children...)
}

// ========== Forms ==========

// Button creates a <button> element.
func Button(attrs attrs.Props, children ...Node) *Element {
	return newElement("button", attrs, children...)
}

// Form creates a <form> element.
func Form(attrs attrs.Props, children ...Node) *Element {
	return newElement("form", attrs, children...)
}

// Input creates an <input> element.
func Input(attrs attrs.Props) *Element {
	return newElement("input", attrs)
}

// Label creates a <label> element.
func Label(attrs attrs.Props, children ...Node) *Element {
	return newElement("label", attrs, children...)
}

// Optgroup creates an <optgroup> element to group <option>s within a <select> element.
func Optgroup(attrs attrs.Props, children ...Node) *Element {
	return newElement("optgroup", attrs, children...)
}

// Option creates an <option> element.
func Option(attrs attrs.Props, content TextNode) *Element {
	return newElement("option", attrs, content)
}

// Select creates a <select> element.
func Select(attrs attrs.Props, children ...Node) *Element {
	return newElement("select", attrs, children...)
}

// Textarea creates a <textarea> element.
func Textarea(attrs attrs.Props, content TextNode) *Element {
	return newElement("textarea", attrs, content)
}

// ========== Hyperlinks and Multimedia ==========

// Img creates an <img> element.
func Img(attrs attrs.Props) *Element {
	return newElement("img", attrs)
}

// ========== Meta Elements ==========

// Link creates a <link> element.
func Link(attrs attrs.Props) *Element {
	return newElement("link", attrs)
}

// Meta creates a <meta> element.
func Meta(attrs attrs.Props) *Element {
	return newElement("meta", attrs)
}

// Script creates a <script> element.
func Script(attrs attrs.Props, children ...Node) *Element {
	return newElement("script", attrs, children...)
}

// Style creates a <style> element.
func Style(attrs attrs.Props, children ...Node) *Element {
	return newElement("style", attrs, children...)
}

// ========== Semantic Elements ==========

// --- Semantic Sectioning Elements ---

// Article creates an <article> element.
func Article(attrs attrs.Props, children ...Node) *Element {
	return newElement("article", attrs, children...)
}

// Aside creates an <aside> element.
func Aside(attrs attrs.Props, children ...Node) *Element {
	return newElement("aside", attrs, children...)
}

// Footer creates a <footer> element.
func Footer(attrs attrs.Props, children ...Node) *Element {
	return newElement("footer", attrs, children...)
}

// Header creates a <header> element.
func Header(attrs attrs.Props, children ...Node) *Element {
	return newElement("header", attrs, children...)
}

// Main creates a <main> element.
func Main(attrs attrs.Props, children ...Node) *Element {
	return newElement("main", attrs, children...)
}

// Nav creates a <nav> element.
func Nav(attrs attrs.Props, children ...Node) *Element {
	return newElement("nav", attrs, children...)
}

// Section creates a <section> element.
func Section(attrs attrs.Props, children ...Node) *Element {
	return newElement("section", attrs, children...)
}

// ========== Semantic Form Elements ==========

// Fieldset creates a <fieldset> element.
func Fieldset(attrs attrs.Props, children ...Node) *Element {
	return newElement("fieldset", attrs, children...)
}

// Legend creates a <legend> element.
func Legend(attrs attrs.Props, children ...Node) *Element {
	return newElement("legend", attrs, children...)
}

// Datalist creates a <datalist> element.
func Datalist(attrs attrs.Props, children ...Node) *Element {
	return newElement("datalist", attrs, children...)
}

// Meter creates a <meter> element.
func Meter(attrs attrs.Props, children ...Node) *Element {
	return newElement("meter", attrs, children...)
}

// Output creates an <output> element.
func Output(attrs attrs.Props, children ...Node) *Element {
	return newElement("output", attrs, children...)
}

// Progress creates a <progress> element.
func Progress(attrs attrs.Props, children ...Node) *Element {
	return newElement("progress", attrs, children...)
}

// --- Semantic Interactive Elements ---

// Dialog creates a <dialog> element.
func Dialog(attrs attrs.Props, children ...Node) *Element {
	return newElement("dialog", attrs, children...)
}

// Menu creates a <menu> element.
func Menu(attrs attrs.Props, children ...Node) *Element {
	return newElement("menu", attrs, children...)
}

// --- Semantic Script Supporting Elements ---

// NoScript creates a <noscript> element.
func NoScript(attrs attrs.Props, children ...Node) *Element {
	return newElement("noscript", attrs, children...)
}

// --- Semantic Text Content Elements ---

// Address creates an <address> element.
func Address(attrs attrs.Props, children ...Node) *Element {
	return newElement("address", attrs, children...)
}

// Details creates a <details> element.
func Details(attrs attrs.Props, children ...Node) *Element {
	return newElement("details", attrs, children...)
}

// FigCaption creates a <figcaption> element.
func FigCaption(attrs attrs.Props, children ...Node) *Element {
	return newElement("figcaption", attrs, children...)
}

// Figure creates a <figure> element.
func Figure(attrs attrs.Props, children ...Node) *Element {
	return newElement("figure", attrs, children...)
}

// Mark creates a <mark> element.
func Mark(attrs attrs.Props, children ...Node) *Element {
	return newElement("mark", attrs, children...)
}

// Summary creates a <summary> element.
func Summary(attrs attrs.Props, children ...Node) *Element {
	return newElement("summary", attrs, children...)
}

// Time creates a <time> element.
func Time(attrs attrs.Props, children ...Node) *Element {
	return newElement("time", attrs, children...)
}

// ========== Tables ==========

// Table creates a <table> element.
func Table(attrs attrs.Props, children ...Node) *Element {
	return newElement("table", attrs, children...)
}

// THead creates a <thead> element.
func THead(attrs attrs.Props, children ...Node) *Element {
	return newElement("thead", attrs, children...)
}

// TBody creates a <tbody> element.
func TBody(attrs attrs.Props, children ...Node) *Element {
	return newElement("tbody", attrs, children...)
}

// TFoot creates a <tfoot> element.
func TFoot(attrs attrs.Props, children ...Node) *Element {
	return newElement("tfoot", attrs, children...)
}

// Tr creates a <tr> element.
func Tr(attrs attrs.Props, children ...Node) *Element {
	return newElement("tr", attrs, children...)
}

// Th creates a <th> element.
func Th(attrs attrs.Props, children ...Node) *Element {
	return newElement("th", attrs, children...)
}

// Td creates a <td> element.
func Td(attrs attrs.Props, children ...Node) *Element {
	return newElement("td", attrs, children...)
}

// ========== Embedded Content ==========

// IFrames creates an <iframe> element.
func IFrame(attrs attrs.Props, children ...Node) *Element {
	return newElement("iframe", attrs, children...)
}

// Audio creates an <audio> element.
func Audio(attrs attrs.Props, children ...Node) *Element {
	return newElement("audio", attrs, children...)
}

// Video creates a <video> element.
func Video(attrs attrs.Props, children ...Node) *Element {
	return newElement("video", attrs, children...)
}

// Source creates a <source> element.
func Source(attrs attrs.Props, children ...Node) *Element {
	return newElement("source", attrs, children...)
}

// ========== Image Map Elements ==========

// Map creates a <map> element.
func Map(attrs attrs.Props, children ...Node) *Element {
	return newElement("map", attrs, children...)
}

// Area creates an <area> element.
func Area(attrs attrs.Props) *Element {
	return newElement("area", attrs)
}

// ========== Other ==========

// None creates a NoneNode, representing a no-operation in rendering.
func None() NoneNode {
	return NoneNode{}
}

// Raw takes html content and returns a RawNode.
func Raw(html string) RawNode {
	return RawNode(html)
}
