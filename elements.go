package elem

// ========== Document Structure ==========

func Body(props Attrs, children ...interface{}) *Element {
	return NewElement("body", props, children...)
}

func Head(props Attrs, children ...interface{}) *Element {
	return NewElement("head", props, children...)
}

func Html(props Attrs, children ...interface{}) *Element {
	return NewElement("html", props, children...)
}

func Title(props Attrs, children ...interface{}) *Element {
	return NewElement("title", props, children...)
}

// ========== Text Formatting and Structure ==========

func A(props Attrs, children ...interface{}) *Element {
	return NewElement("a", props, children...)
}

func Br(props Attrs) *Element {
	return NewElement("br", props)
}

func Blockquote(props Attrs, children ...interface{}) *Element {
	return NewElement("blockquote", props, children...)
}

func Code(props Attrs, children ...interface{}) *Element {
	return NewElement("code", props, children...)
}

func Div(props Attrs, children ...interface{}) *Element {
	return NewElement("div", props, children...)
}

func Em(props Attrs, children ...interface{}) *Element {
	return NewElement("em", props, children...)
}

func H1(props Attrs, children ...interface{}) *Element {
	return NewElement("h1", props, children...)
}

func H2(props Attrs, children ...interface{}) *Element {
	return NewElement("h2", props, children...)
}

func H3(props Attrs, children ...interface{}) *Element {
	return NewElement("h3", props, children...)
}

func Hr(props Attrs) *Element {
	return NewElement("hr", props)
}

func P(props Attrs, children ...interface{}) *Element {
	return NewElement("p", props, children...)
}

func Pre(props Attrs, children ...interface{}) *Element {
	return NewElement("pre", props, children...)
}

func Span(props Attrs, children ...interface{}) *Element {
	return NewElement("span", props, children...)
}

func Strong(props Attrs, children ...interface{}) *Element {
	return NewElement("strong", props, children...)
}

func Text(content string) string {
	return content
}

// ========== Lists ==========

func Li(props Attrs, children ...interface{}) *Element {
	return NewElement("li", props, children...)
}

func Ul(props Attrs, children ...interface{}) *Element {
	return NewElement("ul", props, children...)
}

// ========== Forms ==========

func Button(props Attrs, children ...interface{}) *Element {
	return NewElement("button", props, children...)
}

func Form(attrs Attrs, children ...interface{}) *Element {
	return NewElement("form", attrs, children...)
}

func Input(attrs Attrs) *Element {
	return NewElement("input", attrs)
}

func Label(attrs Attrs, children ...interface{}) *Element {
	return NewElement("label", attrs, children...)
}

func Option(attrs Attrs, content string) *Element {
	return NewElement("option", attrs, content)
}

func Select(attrs Attrs, children ...interface{}) *Element {
	return NewElement("select", attrs, children...)
}

func Textarea(attrs Attrs, content string) *Element {
	return NewElement("textarea", attrs, content)
}

// ========== Hyperlinks and Multimedia ==========

func Img(props Attrs) *Element {
	return NewElement("img", props)
}

// ========== Meta Elements ==========

func Meta(props Attrs) *Element {
	return NewElement("meta", props)
}

func Script(props Attrs, children ...interface{}) *Element {
	return NewElement("script", props, children...)
}
