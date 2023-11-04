package elem

import (
	"github.com/chasefleming/elem-go/attrs"
)

// ========== Document Structure ==========

func Body(attrs attrs.Props, children ...Node) *Element {
	return NewElement("body", attrs, children...)
}

func Head(attrs attrs.Props, children ...Node) *Element {
	return NewElement("head", attrs, children...)
}

func Html(attrs attrs.Props, children ...Node) *Element {
	return NewElement("html", attrs, children...)
}

func Title(attrs attrs.Props, children ...Node) *Element {
	return NewElement("title", attrs, children...)
}

// ========== Text Formatting and Structure ==========

func A(attrs attrs.Props, children ...Node) *Element {
	return NewElement("a", attrs, children...)
}

func Br(attrs attrs.Props) *Element {
	return NewElement("br", attrs)
}

func Blockquote(attrs attrs.Props, children ...Node) *Element {
	return NewElement("blockquote", attrs, children...)
}

func Code(attrs attrs.Props, children ...Node) *Element {
	return NewElement("code", attrs, children...)
}

func Div(attrs attrs.Props, children ...Node) *Element {
	return NewElement("div", attrs, children...)
}

func Em(attrs attrs.Props, children ...Node) *Element {
	return NewElement("em", attrs, children...)
}

func H1(attrs attrs.Props, children ...Node) *Element {
	return NewElement("h1", attrs, children...)
}

func H2(attrs attrs.Props, children ...Node) *Element {
	return NewElement("h2", attrs, children...)
}

func H3(attrs attrs.Props, children ...Node) *Element {
	return NewElement("h3", attrs, children...)
}

func Hr(attrs attrs.Props) *Element {
	return NewElement("hr", attrs)
}

func I(attrs attrs.Props, children ...Node) *Element {
	return NewElement("i", attrs, children...)
}

func P(attrs attrs.Props, children ...Node) *Element {
	return NewElement("p", attrs, children...)
}

func Pre(attrs attrs.Props, children ...Node) *Element {
	return NewElement("pre", attrs, children...)
}

func Span(attrs attrs.Props, children ...Node) *Element {
	return NewElement("span", attrs, children...)
}

func Strong(attrs attrs.Props, children ...Node) *Element {
	return NewElement("strong", attrs, children...)
}

func Text(content string) TextNode {
	return TextNode(content)
}

// ========== Lists ==========

func Li(attrs attrs.Props, children ...Node) *Element {
	return NewElement("li", attrs, children...)
}

func Ul(attrs attrs.Props, children ...Node) *Element {
	return NewElement("ul", attrs, children...)
}

func Ol(attrs attrs.Props, children ...Node) *Element {
	return NewElement("ol", attrs, children...)
}

func Dl(attrs attrs.Props, children ...Node) *Element {
	return NewElement("dl", attrs, children...)
}

func Dt(attrs attrs.Props, children ...Node) *Element {
	return NewElement("dt", attrs, children...)
}

func Dd(attrs attrs.Props, children ...Node) *Element {
	return NewElement("dd", attrs, children...)
}

// ========== Forms ==========

func Button(attrs attrs.Props, children ...Node) *Element {
	return NewElement("button", attrs, children...)
}

func Form(attrs attrs.Props, children ...Node) *Element {
	return NewElement("form", attrs, children...)
}

func Input(attrs attrs.Props) *Element {
	return NewElement("input", attrs)
}

func Label(attrs attrs.Props, children ...Node) *Element {
	return NewElement("label", attrs, children...)
}

func Option(attrs attrs.Props, content TextNode) *Element {
	return NewElement("option", attrs, content)
}

func Select(attrs attrs.Props, children ...Node) *Element {
	return NewElement("select", attrs, children...)
}

func Textarea(attrs attrs.Props, content TextNode) *Element {
	return NewElement("textarea", attrs, content)
}

// ========== Hyperlinks and Multimedia ==========

func Img(attrs attrs.Props) *Element {
	return NewElement("img", attrs)
}

// ========== Meta Elements ==========

func Link(attrs attrs.Props) *Element {
	return NewElement("link", attrs)
}

func Meta(attrs attrs.Props) *Element {
	return NewElement("meta", attrs)
}

func Script(attrs attrs.Props, children ...Node) *Element {
	return NewElement("script", attrs, children...)
}

// ========== Semantic Elements ==========

// --- Semantic Sectioning Elements ---

func Article(attrs attrs.Props, children ...Node) *Element {
	return NewElement("article", attrs, children...)
}

func Aside(attrs attrs.Props, children ...Node) *Element {
	return NewElement("aside", attrs, children...)
}

func Footer(attrs attrs.Props, children ...Node) *Element {
	return NewElement("footer", attrs, children...)
}

func Header(attrs attrs.Props, children ...Node) *Element {
	return NewElement("header", attrs, children...)
}

func Main(attrs attrs.Props, children ...Node) *Element {
	return NewElement("main", attrs, children...)
}

func Nav(attrs attrs.Props, children ...Node) *Element {
	return NewElement("nav", attrs, children...)
}

func Section(attrs attrs.Props, children ...Node) *Element {
	return NewElement("section", attrs, children...)
}

// ========== Semantic Form Elements ==========

func Fieldset(attrs attrs.Props, children ...Node) *Element {
	return NewElement("fieldset", attrs, children...)
}

func Legend(attrs attrs.Props, children ...Node) *Element {
	return NewElement("legend", attrs, children...)
}

func Datalist(attrs attrs.Props, children ...Node) *Element {
	return NewElement("datalist", attrs, children...)
}

func Meter(attrs attrs.Props, children ...Node) *Element {
	return NewElement("meter", attrs, children...)
}

func Output(attrs attrs.Props, children ...Node) *Element {
	return NewElement("output", attrs, children...)
}

func Progress(attrs attrs.Props, children ...Node) *Element {
	return NewElement("progress", attrs, children...)
}

// --- Semantic Interactive Elements ---

func Dialog(attrs attrs.Props, children ...Node) *Element {
	return NewElement("dialog", attrs, children...)
}

func Menu(attrs attrs.Props, children ...Node) *Element {
	return NewElement("menu", attrs, children...)
}

// --- Semantic Script Supporting Elements ---

func NoScript(attrs attrs.Props, children ...Node) *Element {
	return NewElement("noscript", attrs, children...)
}

// --- Semantic Text Content Elements ---

func Address(attrs attrs.Props, children ...Node) *Element {
	return NewElement("address", attrs, children...)
}

func Details(attrs attrs.Props, children ...Node) *Element {
	return NewElement("details", attrs, children...)
}

func FigCaption(attrs attrs.Props, children ...Node) *Element {
	return NewElement("figcaption", attrs, children...)
}

func Figure(attrs attrs.Props, children ...Node) *Element {
	return NewElement("figure", attrs, children...)
}

func Mark(attrs attrs.Props, children ...Node) *Element {
	return NewElement("mark", attrs, children...)
}

func Summary(attrs attrs.Props, children ...Node) *Element {
	return NewElement("summary", attrs, children...)
}

func Time(attrs attrs.Props, children ...Node) *Element {
	return NewElement("time", attrs, children...)
}

// ========== Tables ==========

func Table(attrs attrs.Props, children ...Node) *Element {
	return NewElement("table", attrs, children...)
}

func THead(attrs attrs.Props, children ...Node) *Element {
	return NewElement("thead", attrs, children...)
}

func TBody(attrs attrs.Props, children ...Node) *Element {
	return NewElement("tbody", attrs, children...)
}

func TFoot(attrs attrs.Props, children ...Node) *Element {
	return NewElement("tfoot", attrs, children...)
}

func Tr(attrs attrs.Props, children ...Node) *Element {
	return NewElement("tr", attrs, children...)
}

func Th(attrs attrs.Props, children ...Node) *Element {
	return NewElement("th", attrs, children...)
}

func Td(attrs attrs.Props, children ...Node) *Element {
	return NewElement("td", attrs, children...)
}
