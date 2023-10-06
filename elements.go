package elem

func Body(props Props, children ...interface{}) *Element {
	return NewElement("body", props, children...)
}

func Button(props Props, children ...interface{}) *Element {
	return NewElement("button", props, children...)
}

func Div(props Props, children ...interface{}) *Element {
	return NewElement("div", props, children...)
}

func H1(props Props, children ...interface{}) *Element {
	return NewElement("h1", props, children...)
}

func H2(props Props, children ...interface{}) *Element {
	return NewElement("h2", props, children...)
}

func H3(props Props, children ...interface{}) *Element {
	return NewElement("h3", props, children...)
}

func Head(props Props, children ...interface{}) *Element {
	return NewElement("head", props, children...)
}

func Html(props Props, children ...interface{}) *Element {
	return NewElement("html", props, children...)
}

func P(props Props, children ...interface{}) *Element {
	return NewElement("p", props, children...)
}

func A(props Props, children ...interface{}) *Element {
	return NewElement("a", props, children...)
}

func Span(props Props, children ...interface{}) *Element {
	return NewElement("span", props, children...)
}

func Ul(props Props, children ...interface{}) *Element {
	return NewElement("ul", props, children...)
}

func Li(props Props, children ...interface{}) *Element {
	return NewElement("li", props, children...)
}

func Img(props Props) *Element {
	return NewElement("img", props)
}

func Meta(props Props) *Element {
	return NewElement("meta", props)
}

func Script(props Props) *Element {
	return NewElement("script", props)
}

func Text(content string) string {
	return content
}

func Title(props Props, children ...interface{}) *Element {
	return NewElement("title", props, children...)
}
