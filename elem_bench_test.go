package elem

import (
	"fmt"
	"testing"

	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/styles"
)

func BenchmarkRenderSimpleElement(b *testing.B) {
	el := Div(attrs.Props{attrs.ID: "container", attrs.Class: "main"}, Text("Hello, World!"))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		el.Render()
	}
}

func BenchmarkRenderManyAttributes(b *testing.B) {
	el := Input(attrs.Props{
		attrs.Type:         "text",
		attrs.ID:           "username",
		attrs.Name:         "username",
		attrs.Class:        "form-control input-lg",
		attrs.Placeholder:  "Enter your username",
		attrs.Value:        "chase",
		attrs.Required:     "true",
		attrs.Autofocus:    "true",
		attrs.Maxlength:    "64",
		attrs.Autocomplete: "off",
	})
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		el.Render()
	}
}

func BenchmarkRenderLargeList(b *testing.B) {
	items := make([]string, 1000)
	for i := range items {
		items[i] = fmt.Sprintf("Item %d", i)
	}
	list := Ul(attrs.Props{attrs.Class: "list"},
		TransformEach(items, func(item string) Node {
			return Li(attrs.Props{attrs.Class: "list-item"}, Text(item))
		})...,
	)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Render()
	}
}

func BenchmarkRenderDeepNesting(b *testing.B) {
	el := Span(nil, Text("leaf"))
	for i := 0; i < 100; i++ {
		el = Div(attrs.Props{attrs.Class: "level"}, el)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		el.Render()
	}
}

func BenchmarkRenderTextEscaping(b *testing.B) {
	el := P(nil, Text(`5 < 6 && 7 > 4, "quoted" & 'single' <script>alert("xss")</script>`))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		el.Render()
	}
}

// BenchmarkBuildAndRenderPage measures constructing the element tree and
// rendering it, mirroring the typical per-request usage pattern.
func BenchmarkBuildAndRenderPage(b *testing.B) {
	rows := make([]int, 50)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		page := Html(nil,
			Head(nil,
				Title(nil, Text("Dashboard")),
				Meta(attrs.Props{attrs.Charset: "utf-8"}),
			),
			Body(attrs.Props{attrs.Style: styles.Props{
				styles.Margin:     "0",
				styles.FontFamily: "sans-serif",
			}.ToInline()},
				Header(attrs.Props{attrs.Class: "header"},
					H1(nil, Text("Dashboard")),
				),
				Table(attrs.Props{attrs.Class: "data"},
					TransformEach(rows, func(row int) Node {
						return Tr(nil,
							Td(nil, Text(fmt.Sprintf("Row %d", row))),
							Td(nil, Text("value")),
							Td(nil, Text("status")),
						)
					})...,
				),
			),
		)
		page.Render()
	}
}
