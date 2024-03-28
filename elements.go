package elem

import (
	"github.com/chasefleming/elem-go/attrs"
)

// ========== Document Structure ==========

// Body creates a <body> element.
func Body(children ...Node) *Element {
	return newElement("body", children...)
}

// Head creates a <head> element.
func Head(children ...Node) *Element {
	return newElement("head", children...)
}

// Html creates an <html> element.
func Html(children ...Node) *Element {
	return newElement("html", children...)
}

// Title creates a <title> element.
func Title(children ...Node) *Element {
	return newElement("title", children...)
}

// ========== Text Formatting and Structure ==========

// A creates an <a> element.
func A(children ...Node) *Element {
	return newElement("a", children...)
}

// Br creates a <br> element.
func Br(attrs attrs.Props) *Element {
	return newElement("br", attrs)
}

// Blockquote creates a <blockquote> element.
func Blockquote(children ...Node) *Element {
	return newElement("blockquote", children...)
}

// Code creates a <code> element.
func Code(children ...Node) *Element {
	return newElement("code", children...)
}

// Div creates a <div> element.
func Div(children ...Node) *Element {
	return newElement("div", children...)
}

// Em creates an <em> element.
func Em(children ...Node) *Element {
	return newElement("em", children...)
}

// H1 creates an <h1> element.
func H1(children ...Node) *Element {
	return newElement("h1", children...)
}

// H2 creates an <h2> element.
func H2(children ...Node) *Element {
	return newElement("h2", children...)
}

// H3 creates an <h3> element.
func H3(children ...Node) *Element {
	return newElement("h3", children...)
}

// H4 creates an <h4> element.
func H4(children ...Node) *Element {
	return newElement("h4", children...)
}

// H5 creates an <h5> element.
func H5(children ...Node) *Element {
	return newElement("h5", children...)
}

// H6 creates an <h6> element.
func H6(children ...Node) *Element {
	return newElement("h6", children...)
}

// Hgroup creates an <hgroup> element.
func Hgroup(children ...Node) *Element {
	return newElement("hgroup", children...)
}

// Hr creates an <hr> element.
func Hr(attrs attrs.Props) *Element {
	return newElement("hr", attrs)
}

// I creates an <i> element.
func I(children ...Node) *Element {
	return newElement("i", children...)
}

// P creates a <p> element.
func P(children ...Node) *Element {
	return newElement("p", children...)
}

// Pre creates a <pre> element.
func Pre(children ...Node) *Element {
	return newElement("pre", children...)
}

// Span creates a <span> element.
func Span(children ...Node) *Element {
	return newElement("span", children...)
}

// Strong creates a <strong> element.
func Strong(children ...Node) *Element {
	return newElement("strong", children...)
}

// Sub creates a <sub> element.
func Sub(children ...Node) *Element {
	return newElement("sub", children...)
}

// Sup creates a <sub> element.
func Sup(children ...Node) *Element {
	return newElement("sup", children...)
}

// B creates a <b> element.
func B(children ...Node) *Element {
	return newElement("b", children...)
}

// U creates a <u> element.
func U(children ...Node) *Element {
	return newElement("u", children...)
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
func Li(children ...Node) *Element {
	return newElement("li", children...)
}

// Ul creates a <ul> element.
func Ul(children ...Node) *Element {
	return newElement("ul", children...)
}

// Ol creates an <ol> element.
func Ol(children ...Node) *Element {
	return newElement("ol", children...)
}

// Dl creates a <dl> element.
func Dl(children ...Node) *Element {
	return newElement("dl", children...)
}

// Dt creates a <dt> element.
func Dt(children ...Node) *Element {
	return newElement("dt", children...)
}

// Dd creates a <dd> element.
func Dd(children ...Node) *Element {
	return newElement("dd", children...)
}

// ========== Forms ==========

// Button creates a <button> element.
func Button(children ...Node) *Element {
	return newElement("button", children...)
}

// Form creates a <form> element.
func Form(children ...Node) *Element {
	return newElement("form", children...)
}

// Input creates an <input> element.
func Input(attrs attrs.Props) *Element {
	return newElement("input", attrs)
}

// Label creates a <label> element.
func Label(children ...Node) *Element {
	return newElement("label", children...)
}

// Optgroup creates an <optgroup> element to group <option>s within a <select> element.
func Optgroup(children ...Node) *Element {
	return newElement("optgroup", children...)
}

// Option creates an <option> element.
// TODO this will not be 100% correct
func Option(children ...Node) *Element {
	return newElement("option", children...)
}

// Select creates a <select> element.
func Select(children ...Node) *Element {
	return newElement("select", children...)
}

// Textarea creates a <textarea> element.
// This will also not be 100% correct
func Textarea(children ...Node) *Element {
	return newElement("textarea", children...)
}

// ========== Hyperlinks and Multimedia ==========

// Img creates an <img> element.
func Img(attrs attrs.Props) *Element {
	return newElement("img", attrs)
}

// ========== Head Elements ==========

// Base creates a <base> element.
func Base(attrs attrs.Props) *Element {
	return newElement("base", attrs)
}

// Link creates a <link> element.
func Link(attrs attrs.Props) *Element {
	return newElement("link", attrs)
}

// Meta creates a <meta> element.
func Meta(attrs attrs.Props) *Element {
	return newElement("meta", attrs)
}

// Script creates a <script> element.
func Script(children ...Node) *Element {
	return newElement("script", children...)
}

// Style creates a <style> element.
func Style(children ...Node) *Element {
	return newElement("style", children...)
}

// ========== Semantic Elements ==========

// --- Semantic Sectioning Elements ---

// Article creates an <article> element.
func Article(children ...Node) *Element {
	return newElement("article", children...)
}

// Aside creates an <aside> element.
func Aside(children ...Node) *Element {
	return newElement("aside", children...)
}

// Footer creates a <footer> element.
func Footer(children ...Node) *Element {
	return newElement("footer", children...)
}

// Header creates a <header> element.
func Header(children ...Node) *Element {
	return newElement("header", children...)
}

// Main creates a <main> element.
func Main(children ...Node) *Element {
	return newElement("main", children...)
}

// Nav creates a <nav> element.
func Nav(children ...Node) *Element {
	return newElement("nav", children...)
}

// Section creates a <section> element.
func Section(children ...Node) *Element {
	return newElement("section", children...)
}

// Details creates a <details> element.
func Details(children ...Node) *Element {
	return newElement("details", children...)
}

// Summary creates a <summary> element.
func Summary(children ...Node) *Element {
	return newElement("summary", children...)
}

// ========== Semantic Form Elements ==========

// Fieldset creates a <fieldset> element.
func Fieldset(children ...Node) *Element {
	return newElement("fieldset", children...)
}

// Legend creates a <legend> element.
func Legend(children ...Node) *Element {
	return newElement("legend", children...)
}

// Datalist creates a <datalist> element.
func Datalist(children ...Node) *Element {
	return newElement("datalist", children...)
}

// Meter creates a <meter> element.
func Meter(children ...Node) *Element {
	return newElement("meter", children...)
}

// Output creates an <output> element.
func Output(children ...Node) *Element {
	return newElement("output", children...)
}

// Progress creates a <progress> element.
func Progress(children ...Node) *Element {
	return newElement("progress", children...)
}

// --- Semantic Interactive Elements ---

// Dialog creates a <dialog> element.
func Dialog(children ...Node) *Element {
	return newElement("dialog", children...)
}

// Menu creates a <menu> element.
func Menu(children ...Node) *Element {
	return newElement("menu", children...)
}

// --- Semantic Script Supporting Elements ---

// NoScript creates a <noscript> element.
func NoScript(children ...Node) *Element {
	return newElement("noscript", children...)
}

// --- Semantic Text Content Elements ---

// Abbr creates an <abbr> element.
func Abbr(children ...Node) *Element {
	return newElement("abbr", children...)
}

// Address creates an <address> element.
func Address(children ...Node) *Element {
	return newElement("address", children...)
}

// Cite creates a <cite> element.
func Cite(children ...Node) *Element {
	return newElement("cite", children...)
}

// Data creates a <data> element.
func Data(children ...Node) *Element {
	return newElement("data", children...)
}

// FigCaption creates a <figcaption> element.
func FigCaption(children ...Node) *Element {
	return newElement("figcaption", children...)
}

// Figure creates a <figure> element.
func Figure(children ...Node) *Element {
	return newElement("figure", children...)
}

// Kbd creates a <kbd> element.
func Kbd(children ...Node) *Element {
	return newElement("kbd", children...)
}

// Mark creates a <mark> element.
func Mark(children ...Node) *Element {
	return newElement("mark", children...)
}

// Q creates a <q> element.
func Q(children ...Node) *Element {
	return newElement("q", children...)
}

// Samp creates a <samp> element.
func Samp(children ...Node) *Element {
	return newElement("samp", children...)
}

// Small creates a <small> element.
func Small(children ...Node) *Element {
	return newElement("small", children...)
}

// Time creates a <time> element.
func Time(children ...Node) *Element {
	return newElement("time", children...)
}

// Var creates a <var> element.
func Var(children ...Node) *Element {
	return newElement("var", children...)
}

// Ruby creates a <ruby> element.
func Ruby(children ...Node) *Element {
	return newElement("ruby", children...)
}

// Rt creates a <rt> element.
func Rt(children ...Node) *Element {
	return newElement("rt", children...)
}

// Rp creates a <rp> element.
func Rp(children ...Node) *Element {
	return newElement("rp", children...)
}

// ========== Tables ==========

// Table creates a <table> element.
func Table(children ...Node) *Element {
	return newElement("table", children...)
}

// THead creates a <thead> element.
func THead(children ...Node) *Element {
	return newElement("thead", children...)
}

// TBody creates a <tbody> element.
func TBody(children ...Node) *Element {
	return newElement("tbody", children...)
}

// TFoot creates a <tfoot> element.
func TFoot(children ...Node) *Element {
	return newElement("tfoot", children...)
}

// Tr creates a <tr> element.
func Tr(children ...Node) *Element {
	return newElement("tr", children...)
}

// Th creates a <th> element.
func Th(children ...Node) *Element {
	return newElement("th", children...)
}

// Td creates a <td> element.
func Td(children ...Node) *Element {
	return newElement("td", children...)
}

// ========== Embedded Content ==========

// IFrames creates an <iframe> element.
func IFrame(children ...Node) *Element {
	return newElement("iframe", children...)
}

// Audio creates an <audio> element.
func Audio(children ...Node) *Element {
	return newElement("audio", children...)
}

// Video creates a <video> element.
func Video(children ...Node) *Element {
	return newElement("video", children...)
}

// Source creates a <source> element.
func Source(children ...Node) *Element {
	return newElement("source", children...)
}

// ========== Image Map Elements ==========

// Map creates a <map> element.
func Map(children ...Node) *Element {
	return newElement("map", children...)
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

// CSS takes css content and returns a TextNode.
func CSS(content string) TextNode {
	return TextNode(content)
}
