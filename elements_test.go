package elem

import (
	"testing"

	"github.com/chasefleming/elem-go/attrs"
	"github.com/stretchr/testify/assert"
)

// ========== Document Structure ==========

func TestBody(t *testing.T) {
	expected := `<body class="page-body"><p>Welcome to Elem!</p></body>`
	el := Body(attrs.Props{attrs.Class: "page-body"}, P(nil, Text("Welcome to Elem!")))
	assert.Equal(t, expected, el.Render())
}

func TestHtml(t *testing.T) {
	expected := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Elem Page</title></head><body><p>Welcome to Elem!</p></body></html>`
	el := Html(attrs.Props{attrs.Lang: "en"},
		Head(nil,
			Meta(attrs.Props{attrs.Charset: "UTF-8"}),
			Title(nil, Text("Elem Page")),
		),
		Body(nil, P(nil, Text("Welcome to Elem!"))),
	)
	assert.Equal(t, expected, el.Render())
}

func TestHtmlWithOptions(t *testing.T) {
	expected := `<html lang="en"><head><meta charset="UTF-8"><title>Elem Page</title></head><body><p>Welcome to Elem!</p></body></html>`
	el := Html(attrs.Props{attrs.Lang: "en"},
		Head(nil,
			Meta(attrs.Props{attrs.Charset: "UTF-8"}),
			Title(nil, Text("Elem Page")),
		),
		Body(nil, P(nil, Text("Welcome to Elem!"))),
	)
	assert.Equal(t, expected, el.RenderWithOptions(RenderOptions{DisableHtmlPreamble: true}))
}

// ========== Text Formatting and Structure ==========

func TestA(t *testing.T) {
	expected := `<a href="https://example.com">Visit Example</a>`
	el := A(attrs.Props{attrs.Href: "https://example.com"}, Text("Visit Example"))
	assert.Equal(t, expected, el.Render())
}

func TestBlockquote(t *testing.T) {
	expected := `<blockquote>Quote text</blockquote>`
	el := Blockquote(nil, Text("Quote text"))
	assert.Equal(t, expected, el.Render())
}

func TestBr(t *testing.T) {
	expected := `<br>`
	el := Br(nil)
	assert.Equal(t, expected, el.Render())
}

func TestCode(t *testing.T) {
	expected := `<code>Code snippet</code>`
	el := Code(nil, Text("Code snippet"))
	assert.Equal(t, expected, el.Render())
}

func TestDiv(t *testing.T) {
	expected := `<div class="container">Hello, Elem!</div>`
	el := Div(attrs.Props{attrs.Class: "container"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestEm(t *testing.T) {
	expected := `<em>Italic text</em>`
	el := Em(nil, Text("Italic text"))
	assert.Equal(t, expected, el.Render())
}

func TestH1(t *testing.T) {
	expected := `<h1 class="title">Hello, Elem!</h1>`
	el := H1(attrs.Props{attrs.Class: "title"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestH2(t *testing.T) {
	expected := `<h2 class="subtitle">Hello, Elem!</h2>`
	el := H2(attrs.Props{attrs.Class: "subtitle"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestH3(t *testing.T) {
	expected := `<h3>Hello, Elem!</h3>`
	el := H3(nil, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestH4(t *testing.T) {
	expected := `<h4>Hello, Elem!</h4>`
	el := H4(nil, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestH5(t *testing.T) {
	expected := `<h5>Hello, Elem!</h5>`
	el := H5(nil, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestH6(t *testing.T) {
	expected := `<h6>Hello, Elem!</h6>`
	el := H6(nil, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestHr(t *testing.T) {
	expected := `<hr>`
	el := Hr(nil)
	assert.Equal(t, expected, el.Render())
}

func TestI(t *testing.T) {
	expected1 := `<i>Idiomatic Text</i>`
	expected2 := `<i class="fa-regular fa-face-smile"></i>`
	el := I(nil, Text("Idiomatic Text"))
	assert.Equal(t, expected1, el.Render())
	el = I(attrs.Props{attrs.Class: "fa-regular fa-face-smile"})
	assert.Equal(t, expected2, el.Render())
}

func TestP(t *testing.T) {
	expected := `<p>Hello, Elem!</p>`
	el := P(nil, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestPre(t *testing.T) {
	expected := `<pre>Preformatted text</pre>`
	el := Pre(nil, Text("Preformatted text"))
	assert.Equal(t, expected, el.Render())
}

func TestSpan(t *testing.T) {
	expected := `<span class="highlight">Hello, Elem!</span>`
	el := Span(attrs.Props{attrs.Class: "highlight"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestStrong(t *testing.T) {
	expected := `<strong>Bold text</strong>`
	el := Strong(nil, Text("Bold text"))
	assert.Equal(t, expected, el.Render())
}

func TestSub(t *testing.T) {
	expected := `<sub>2</sub>`
	el := Sub(nil, Text("2"))
	assert.Equal(t, expected, el.Render())
}

func TestSup(t *testing.T) {
	expected := `<sup>2</sup>`
	el := Sup(nil, Text("2"))
	assert.Equal(t, expected, el.Render())
}

func TestB(t *testing.T) {
	expected := `<b>Important text</b>`
	el := B(nil, Text("Important text"))
	assert.Equal(t, expected, el.Render())
}

func TestU(t *testing.T) {
	expected := `<u>Unarticulated text</u>`
	el := U(nil, Text("Unarticulated text"))
	assert.Equal(t, expected, el.Render())
}

// ========= Cdata ==========
func TestCdata(t *testing.T) {
	expected := `<![CDATA[Some CDATA content]]>`
	actual := CdataNode("Some CDATA content").Render()
	assert.Equal(t, expected, actual)
}

func TestCdataEscaping (t *testing.T) {
	expected := `<![CDATA[Some CDATA content with ]]&gt; in it]]>`
	actual := CdataNode("Some CDATA content with ]]> in it").Render()
	assert.Equal(t, expected, actual)

	expected = `<![CDATA[These < and > should not be escaped]]>`
	actual = CdataNode("These < and > should not be escaped").Render()
	assert.Equal(t, expected, actual)
}

// ========== Comments ==========
func TestComment(t *testing.T) {
	expected := `<!-- this is a comment -->`
	actual := Comment("this is a comment").Render()
	assert.Equal(t, expected, actual)
}

func TestCommentInElement(t *testing.T) {
	expected := `<div>not a comment<!-- this is a comment --></div>`
	actual := Div(nil, Text("not a comment"), Comment("this is a comment")).Render()
	assert.Equal(t, expected, actual)
}

func TestCommentEscaping(t *testing.T) {
	expected := `<!-- &lt;!-- this is a comment with some disallowed character sequences --&gt; --!&gt; -->`
	actual := Comment("<!-- this is a comment with some disallowed character sequences --> --!>").Render()
	assert.Equal(t, expected, actual)

	expected = `<!-- &gt; comment cannot start with this string -->`
	actual = Comment("> comment cannot start with this string").Render()
	assert.Equal(t, expected, actual)

	expected = `<!-- -&gt; comment cannot start with this string either -->`
	actual = Comment("-> comment cannot start with this string either").Render()
	assert.Equal(t, expected, actual)

	expected = `<!-- comment cannot end with this string &lt;!- -->`
	actual = Comment("comment cannot end with this string <!-").Render()
	assert.Equal(t, expected, actual)
}

// ========== Lists ==========

func TestLi(t *testing.T) {
	expected := `<li>Item 1</li>`
	el := Li(nil, Text("Item 1"))
	assert.Equal(t, expected, el.Render())
}

func TestUl(t *testing.T) {
	expected := `<ul><li>Item 1</li><li>Item 2</li></ul>`
	el := Ul(nil, Li(nil, Text("Item 1")), Li(nil, Text("Item 2")))
	assert.Equal(t, expected, el.Render())
}

func TestOl(t *testing.T) {
	expected := `<ol><li>Item 1</li><li>Item 2</li></ol>`
	el := Ol(nil, Li(nil, Text("Item 1")), Li(nil, Text("Item 2")))
	assert.Equal(t, expected, el.Render())
}

func TestDl(t *testing.T) {
	expected := `<dl><dt>Term 1</dt><dd>Description 1</dd><dt>Term 2</dt><dd>Description 2</dd></dl>`
	el := Dl(nil, Dt(nil, Text("Term 1")), Dd(nil, Text("Description 1")), Dt(nil, Text("Term 2")), Dd(nil, Text("Description 2")))
	assert.Equal(t, expected, el.Render())
}

func TestDt(t *testing.T) {
	expected := `<dt>Term 1</dt>`
	el := Dt(nil, Text("Term 1"))
	assert.Equal(t, expected, el.Render())
}

func TestDd(t *testing.T) {
	expected := `<dd>Description 1</dd>`
	el := Dd(nil, Text("Description 1"))
	assert.Equal(t, expected, el.Render())
}

// ========== Forms ==========

func TestButton(t *testing.T) {
	expected := `<button class="btn">Click Me</button>`
	el := Button(attrs.Props{attrs.Class: "btn"}, Text("Click Me"))
	assert.Equal(t, expected, el.Render())
}

func TestForm(t *testing.T) {
	expected := `<form action="/submit" method="post"><input name="username" type="text"></form>`
	el := Form(attrs.Props{attrs.Action: "/submit", attrs.Method: "post"}, Input(attrs.Props{attrs.Type: "text", attrs.Name: "username"}))
	assert.Equal(t, expected, el.Render())
}

func TestInput(t *testing.T) {
	expected := `<input name="username" placeholder="Enter your username" type="text">`
	el := Input(attrs.Props{attrs.Type: "text", attrs.Name: "username", attrs.Placeholder: "Enter your username"})
	assert.Equal(t, expected, el.Render())
}

func TestLabel(t *testing.T) {
	expected := `<label for="username">Username</label>`
	el := Label(attrs.Props{attrs.For: "username"}, Text("Username"))
	assert.Equal(t, expected, el.Render())
}

func TestSelectAndOption(t *testing.T) {
	expected := `<select name="color"><option value="red">Red</option><option value="blue">Blue</option></select>`
	el := Select(attrs.Props{attrs.Name: "color"}, Option(attrs.Props{attrs.Value: "red"}, Text("Red")), Option(attrs.Props{attrs.Value: "blue"}, Text("Blue")))
	assert.Equal(t, expected, el.Render())
}

func TestSelectAndOptgroup(t *testing.T) {
	expected := `<select name="cars"><optgroup label="Swedish Cars"><option value="volvo">Volvo</option><option value="saab">Saab</option></optgroup><optgroup label="German Cars"><option value="mercedes">Mercedes</option><option value="audi">Audi</option></optgroup></select>`
	el := Select(attrs.Props{attrs.Name: "cars"}, Optgroup(attrs.Props{attrs.Label: "Swedish Cars"}, Option(attrs.Props{attrs.Value: "volvo"}, Text("Volvo")), Option(attrs.Props{attrs.Value: "saab"}, Text("Saab"))), Optgroup(attrs.Props{attrs.Label: "German Cars"}, Option(attrs.Props{attrs.Value: "mercedes"}, Text("Mercedes")), Option(attrs.Props{attrs.Value: "audi"}, Text("Audi"))))
	assert.Equal(t, expected, el.Render())
}

func TestTextarea(t *testing.T) {
	expected := `<textarea name="comment" rows="5">Leave a comment...</textarea>`
	el := Textarea(attrs.Props{attrs.Name: "comment", attrs.Rows: "5"}, Text("Leave a comment..."))
	assert.Equal(t, expected, el.Render())
}

// ========== Boolean attributes ==========
func TestCheckedTrue(t *testing.T) {
	expected := `<input checked name="allow" type="checkbox">`
	el := Input(attrs.Props{attrs.Type: "checkbox", attrs.Name: "allow", attrs.Checked: "true"})
	assert.Equal(t, expected, el.Render())
}

func TestCheckedFalse(t *testing.T) {
	expected := `<input name="allow" type="checkbox">`
	el := Input(attrs.Props{attrs.Type: "checkbox", attrs.Name: "allow", attrs.Checked: "false"})
	assert.Equal(t, expected, el.Render())
}

func TestCheckedEmpty(t *testing.T) {
	expected := `<input name="allow" type="checkbox">`
	el := Input(attrs.Props{attrs.Type: "checkbox", attrs.Name: "allow", attrs.Checked: ""})
	assert.Equal(t, expected, el.Render())
}

// ========== Hyperlinks and Multimedia ==========

func TestImg(t *testing.T) {
	expected := `<img alt="An image" src="image.jpg">`
	el := Img(attrs.Props{attrs.Src: "image.jpg", attrs.Alt: "An image"})
	assert.Equal(t, expected, el.Render())
}

// ========== Head Elements ==========

func TestBase(t *testing.T) {
	expected := `<base href="https://example.com">`
	el := Base(attrs.Props{attrs.Href: "https://example.com"})
	assert.Equal(t, expected, el.Render())
}

func TestLink(t *testing.T) {
	expected := `<link href="https://example.com/styles.css" rel="stylesheet">`
	el := Link(attrs.Props{attrs.Rel: "stylesheet", attrs.Href: "https://example.com/styles.css"})
	assert.Equal(t, expected, el.Render())
}

func TestMeta(t *testing.T) {
	expected := `<meta charset="UTF-8">`
	el := Meta(attrs.Props{attrs.Charset: "UTF-8"})
	assert.Equal(t, expected, el.Render())
}

func TestScript(t *testing.T) {
	expected := `<script src="https://example.com/script.js"></script>`
	el := Script(attrs.Props{attrs.Src: "https://example.com/script.js"})
	assert.Equal(t, expected, el.Render())
}

func TestStyle(t *testing.T) {
	expected := `<style type="text/css">.test-class {color: #333;}</style>`
	cssContent := `.test-class {color: #333;}`
	el := Style(attrs.Props{attrs.Type: "text/css"}, CSS(cssContent))
	assert.Equal(t, expected, el.Render())
}

// ========== Semantic Elements ==========

// --- Semantic Sectioning Elements ---

func TestArticle(t *testing.T) {
	expected := `<article><h2>Article Title</h2><p>Article content.</p></article>`
	el := Article(nil, H2(nil, Text("Article Title")), P(nil, Text("Article content.")))
	assert.Equal(t, expected, el.Render())
}

func TestAside(t *testing.T) {
	expected := `<aside><p>Sidebar content.</p></aside>`
	el := Aside(nil, P(nil, Text("Sidebar content.")))
	assert.Equal(t, expected, el.Render())
}

func TestFooter(t *testing.T) {
	expected := `<footer><p>Footer content.</p></footer>`
	el := Footer(nil, P(nil, Text("Footer content.")))
	assert.Equal(t, expected, el.Render())
}

func TestHeader(t *testing.T) {
	expected := `<header class="site-header"><h1>Welcome to Elem!</h1></header>`
	el := Header(attrs.Props{attrs.Class: "site-header"}, H1(nil, Text("Welcome to Elem!")))
	assert.Equal(t, expected, el.Render())
}

func TestMainElem(t *testing.T) {
	expected := `<main><p>Main content goes here.</p></main>`
	el := Main(nil, P(nil, Text("Main content goes here.")))
	assert.Equal(t, expected, el.Render())
}

func TestNav(t *testing.T) {
	expected := `<nav><a href="/home">Home</a><a href="/about">About</a></nav>`
	el := Nav(nil, A(attrs.Props{attrs.Href: "/home"}, Text("Home")), A(attrs.Props{attrs.Href: "/about"}, Text("About")))
	assert.Equal(t, expected, el.Render())
}

func TestSection(t *testing.T) {
	expected := `<section><h3>Section Title</h3><p>Section content.</p></section>`
	el := Section(nil, H3(nil, Text("Section Title")), P(nil, Text("Section content.")))
	assert.Equal(t, expected, el.Render())
}

func TestHgroup(t *testing.T) {
	expected := `<hgroup><h1>Frankenstein</h1><p>Or: The Modern Prometheus</p></hgroup>`
	el := Hgroup(nil, H1(nil, Text("Frankenstein")), P(nil, Text("Or: The Modern Prometheus")))
	assert.Equal(t, expected, el.Render())
}

// --- Semantic Form Elements ---

func TestFieldset(t *testing.T) {
	expected := `<fieldset class="custom-fieldset"><legend>Personal Information</legend><input name="name" type="text"></fieldset>`
	el := Fieldset(attrs.Props{attrs.Class: "custom-fieldset"}, Legend(nil, Text("Personal Information")), Input(attrs.Props{attrs.Type: "text", attrs.Name: "name"}))
	assert.Equal(t, expected, el.Render())
}

func TestLegend(t *testing.T) {
	expected := `<legend class="custom-legend">Legend Title</legend>`
	el := Legend(attrs.Props{attrs.Class: "custom-legend"}, Text("Legend Title"))
	assert.Equal(t, expected, el.Render())
}

func TestDatalist(t *testing.T) {
	expected := `<datalist id="exampleList"><option value="Option1">Option 1</option><option value="Option2">Option 2</option></datalist>`
	el := Datalist(attrs.Props{attrs.ID: "exampleList"}, Option(attrs.Props{attrs.Value: "Option1"}, Text("Option 1")), Option(attrs.Props{attrs.Value: "Option2"}, Text("Option 2")))
	assert.Equal(t, expected, el.Render())
}

func TestMeter(t *testing.T) {
	expected := `<meter max="100" min="0" value="50">50%</meter>`
	el := Meter(attrs.Props{attrs.Min: "0", attrs.Max: "100", attrs.Value: "50"}, Text("50%"))
	assert.Equal(t, expected, el.Render())
}

func TestOutput(t *testing.T) {
	expected := `<output for="inputId" name="result">Output</output>`
	el := Output(attrs.Props{attrs.For: "inputId", attrs.Name: "result"}, Text("Output"))
	assert.Equal(t, expected, el.Render())
}

func TestProgress(t *testing.T) {
	expected := `<progress max="100" value="60"></progress>`
	el := Progress(attrs.Props{attrs.Max: "100", attrs.Value: "60"})
	assert.Equal(t, expected, el.Render())
}

// --- Semantic Interactive Elements ---

func TestDialog(t *testing.T) {
	expected := `<dialog open><p>This is an open dialog window</p></dialog>`
	el := Dialog(attrs.Props{attrs.Open: "true"}, P(nil, Text("This is an open dialog window")))
	assert.Equal(t, expected, el.Render())
}

func TestMenu(t *testing.T) {
	expected := `<menu><li>Item One</li><li>Item Two</li></menu>`
	el := Menu(nil, Li(nil, Text("Item One")), Li(nil, Text("Item Two")))
	assert.Equal(t, expected, el.Render())
}

// --- Semantic Script Supporting Elements ---

func TestNoScript(t *testing.T) {
	expected := `<noscript><p>JavaScript is required for this application.</p></noscript>`
	el := NoScript(nil, P(nil, Text("JavaScript is required for this application.")))
	assert.Equal(t, expected, el.Render())
}

// --- Semantic Text Content Elements ---

func TestAbbr(t *testing.T) {
	expected := `<abbr title="Web Hypertext Application Technology Working Group">WHATWG</abbr>`
	el := Abbr(attrs.Props{attrs.Title: "Web Hypertext Application Technology Working Group"}, Text("WHATWG"))
	assert.Equal(t, expected, el.Render())
}

func TestAddress(t *testing.T) {
	expected := `<address>123 Example St.</address>`
	el := Address(nil, Text("123 Example St."))
	assert.Equal(t, expected, el.Render())
}

func TestCite(t *testing.T) {
	expected := `<p>My favorite book is <cite>The Reality Dysfunction</cite> by Peter F. Hamilton.</p>`
	el := P(nil, Text("My favorite book is "), Cite(nil, Text("The Reality Dysfunction")), Text(" by Peter F. Hamilton."))
	assert.Equal(t, expected, el.Render())
}

func TestDetails(t *testing.T) {
	expected := `<details><summary>More Info</summary><p>Details content here.</p></details>`
	el := Details(nil, Summary(nil, Text("More Info")), P(nil, Text("Details content here.")))
	assert.Equal(t, expected, el.Render())
}

func TestDetailsWithOpenFalse(t *testing.T) {
	expected := `<details><summary>More Info</summary><p>Details content here.</p></details>`
	el := Details(attrs.Props{attrs.Open: "false"}, Summary(nil, Text("More Info")), P(nil, Text("Details content here.")))
	assert.Equal(t, expected, el.Render())
}

func TestDetailsWithOpenTrue(t *testing.T) {
	expected := `<details open><summary>More Info</summary><p>Details content here.</p></details>`
	el := Details(attrs.Props{attrs.Open: "true"}, Summary(nil, Text("More Info")), P(nil, Text("Details content here.")))
	assert.Equal(t, expected, el.Render())
}

func TestData(t *testing.T) {
	expected := `<data value="8">Eight</data>`
	el := Data(attrs.Props{attrs.Value: "8"}, Text("Eight"))
	assert.Equal(t, expected, el.Render())
}

func TestFigCaption(t *testing.T) {
	expected := `<figcaption>Description of the figure.</figcaption>`
	el := FigCaption(nil, Text("Description of the figure."))
	assert.Equal(t, expected, el.Render())
}

func TestFigure(t *testing.T) {
	expected := `<figure><img alt="An image" src="image.jpg"><figcaption>An image</figcaption></figure>`
	el := Figure(nil, Img(attrs.Props{attrs.Src: "image.jpg", attrs.Alt: "An image"}), FigCaption(nil, Text("An image")))
	assert.Equal(t, expected, el.Render())
}

func TestKbd(t *testing.T) {
	expected := `<p>To make George eat an apple, select <kbd>File | Eat Apple...</kbd></p>`
	el := P(nil, Text("To make George eat an apple, select "), Kbd(nil, Text("File | Eat Apple...")))
	assert.Equal(t, expected, el.Render())
}

func TestMark(t *testing.T) {
	expected := `<p>You must <mark>highlight</mark> this word.</p>`
	el := P(nil, Text("You must "), Mark(nil, Text("highlight")), Text(" this word."))
	assert.Equal(t, expected, el.Render())
}

func TestQ(t *testing.T) {
	expected := `<p>The W3C's mission is <q cite="https://www.w3.org/Consortium/">To lead the World Wide Web to its full potential</q>.</p>`
	el := P(nil, Text("The W3C's mission is "), Q(attrs.Props{attrs.Cite: "https://www.w3.org/Consortium/"}, Text("To lead the World Wide Web to its full potential")), Text("."))
	assert.Equal(t, expected, el.Render())
}

func TestSamp(t *testing.T) {
	expected := `<p>The computer said <samp>Too much cheese in tray two</samp> but I didn't know what that meant.</p>`
	el := P(nil, Text("The computer said "), Samp(nil, Text("Too much cheese in tray two")), Text(" but I didn't know what that meant."))
	assert.Equal(t, expected, el.Render())
}

func TestSmall(t *testing.T) {
	expected := `<p>Single room <small>breakfast included, VAT not included</small></p>`
	el := P(nil, Text("Single room "), Small(nil, Text("breakfast included, VAT not included")))
	assert.Equal(t, expected, el.Render())
}

func TestSummary(t *testing.T) {
	expected := `<details><summary>Summary Title</summary></details>`
	el := Details(nil, Summary(nil, Text("Summary Title")))
	assert.Equal(t, expected, el.Render())
}

func TestTime(t *testing.T) {
	expected := `<time datetime="2023-01-01T00:00:00Z">New Year's Day</time>`
	el := Time(attrs.Props{attrs.Datetime: "2023-01-01T00:00:00Z"}, Text("New Year's Day"))
	assert.Equal(t, expected, el.Render())
}

func TestVar(t *testing.T) {
	expected := `<p>After a few moment's thought, she wrote <var>E</var>.</p>`
	el := P(nil, Text("After a few moment's thought, she wrote "), Var(nil, Text("E")), Text("."))
	assert.Equal(t, expected, el.Render())
}

func TestRuby(t *testing.T) {
	expected := `<ruby>漢</ruby>`
	el := Ruby(nil, Text("漢"))
	assert.Equal(t, expected, el.Render())
}

func TestRt(t *testing.T) {
	expected := `<ruby> 漢 <rp>(</rp><rt>kan</rt><rp>)</rp> 字 <rp>(</rp><rt>ji</rt><rp>)</rp> </ruby>`
	el := Ruby(nil, Text(" 漢 "), Rp(nil, Text("(")), Rt(nil, Text("kan")), Rp(nil, Text(")")), Text(" 字 "), Rp(nil, Text("(")), Rt(nil, Text("ji")), Rp(nil, Text(")")), Text(" "))
	assert.Equal(t, expected, el.Render())
}

func TestRp(t *testing.T) {
	expected := `<ruby> 漢 <rp>(</rp> 字 <rp>)</rp> </ruby>`
	el := Ruby(nil, Text(" 漢 "), Rp(nil, Text("(")), Text(" 字 "), Rp(nil, Text(")")), Text(" "))
	assert.Equal(t, expected, el.Render())
}

// ========== Tables ==========

func TestTr(t *testing.T) {
	expected := `<tr>Row content.</tr>`
	el := Tr(nil, Text("Row content."))
	assert.Equal(t, expected, el.Render())
}

func TestTd(t *testing.T) {
	expected := `<tr><td><h1>Cell one.</h1></td><td>Cell two.</td></tr>`
	el := Tr(nil, Td(nil, H1(nil, Text("Cell one."))), Td(nil, Text("Cell two.")))
	assert.Equal(t, expected, el.Render())
}

func TestTh(t *testing.T) {
	expected := `<tr><th>First name</th><th>Last name</th><th>Age</th></tr>`
	el := Tr(nil, Th(nil, Text("First name")), Th(nil, Text("Last name")), Th(nil, Text("Age")))
	assert.Equal(t, expected, el.Render())
}

func TestTHead(t *testing.T) {
	expected := `<thead><tr><td>Text</td><td><a href="/link">Link</a></td></tr></thead>`
	el := THead(nil, Tr(nil, Td(nil, Text("Text")), Td(nil, A(attrs.Props{attrs.Href: "/link"}, Text("Link")))))
	assert.Equal(t, expected, el.Render())
}

func TestTBody(t *testing.T) {
	expected := `<tbody><tr><td>Table body</td></tr></tbody>`
	el := TBody(nil, Tr(nil, Td(nil, Text("Table body"))))
	assert.Equal(t, expected, el.Render())
}

func TestTFoot(t *testing.T) {
	expected := `<tfoot><tr><td><a href="/footer">Table footer</a></td></tr></tfoot>`
	el := TFoot(nil, Tr(nil, Td(nil, A(attrs.Props{attrs.Href: "/footer"}, Text("Table footer")))))
	assert.Equal(t, expected, el.Render())
}

func TestTable(t *testing.T) {
	expected := `<table><tr><th>Table header</th></tr><tr><td>Table content</td></tr></table>`
	el := Table(nil, Tr(nil, Th(nil, Text("Table header"))), Tr(nil, Td(nil, Text("Table content"))))
	assert.Equal(t, expected, el.Render())
}

// ========== Embedded Content ==========

func TestEmbedLink(t *testing.T) {
	expected := `<iframe src="https://www.youtube.com/embed/446E-r0rXHI"></iframe>`
	el := IFrame(attrs.Props{attrs.Src: "https://www.youtube.com/embed/446E-r0rXHI"})
	assert.Equal(t, expected, el.Render())
}

func TestAllowFullScreen(t *testing.T) {
	expected := `<iframe allowfullscreen src="https://www.youtube.com/embed/446E-r0rXHI"></iframe>`
	el := IFrame(attrs.Props{attrs.Src: "https://www.youtube.com/embed/446E-r0rXHI", attrs.AllowFullScreen: "true"})
	assert.Equal(t, expected, el.Render())
}

func TestAudioWithSourceElementsAndFallbackText(t *testing.T) {
	expected := `<audio controls><source src="horse.ogg" type="audio/ogg"><source src="horse.mp3" type="audio/mpeg">Your browser does not support the audio tag.</audio>`
	el := Audio(attrs.Props{attrs.Controls: "true"},
		Source(attrs.Props{attrs.Src: "horse.ogg", attrs.Type: "audio/ogg"}),
		Source(attrs.Props{attrs.Src: "horse.mp3", attrs.Type: "audio/mpeg"}),
		Text("Your browser does not support the audio tag."),
	)
	assert.Equal(t, expected, el.Render())
}

func TestVideoWithSourceElementsAndFallbackText(t *testing.T) {
	expected := `<video controls height="240" width="320"><source src="movie.mp4" type="video/mp4"><source src="movie.ogg" type="video/ogg">Your browser does not support the video tag.</video>`
	el := Video(attrs.Props{attrs.Width: "320", attrs.Height: "240", attrs.Controls: "true"},
		Source(attrs.Props{attrs.Src: "movie.mp4", attrs.Type: "video/mp4"}),
		Source(attrs.Props{attrs.Src: "movie.ogg", attrs.Type: "video/ogg"}),
		Text("Your browser does not support the video tag."),
	)
	assert.Equal(t, expected, el.Render())
}

// ========== Image Map Elements ==========

func TestMapAndArea(t *testing.T) {
	expectedMap := `<map name="map-name"><area alt="Area 1" coords="34,44,270,350" href="#area1" shape="rect"></map>`
	mapEl := Map(attrs.Props{attrs.Name: "map-name"},
		Area(attrs.Props{
			attrs.Href:   "#area1",
			attrs.Alt:    "Area 1",
			attrs.Shape:  "rect",
			attrs.Coords: "34,44,270,350",
		}),
	)
	assert.Equal(t, expectedMap, mapEl.Render())
}

// ========== Other ==========

func TestNone(t *testing.T) {
	el := None()
	expected := ""

	assert.Equal(t, expected, el.Render(), "None should render an empty string")
}

func TestNoneInDiv(t *testing.T) {
	expected := `<div></div>`
	actual := Div(nil, None()).Render()
	assert.Equal(t, expected, actual)
}

func TestText(t *testing.T) {
	expected := `<p>Hello, World!</p>`
	el := P(nil, Text("Hello, World!"))
	assert.Equal(t, expected, el.Render())
}

func TestTextWithEscaping(t *testing.T) {
	expected := `<p>Hello, &lt;em&gt;World!&lt;/em&gt;</p>`
	el := P(nil, Text("Hello, <em>World!</em>"))
	assert.Equal(t, expected, el.Render())
}

func TestTextWithNoQuotesEscaping(t *testing.T) {
	expected := `<p>'Hello,' "World!"</p>`
	el := P(nil, Text(`'Hello,' "World!"`))
	assert.Equal(t, expected, el.Render())
}

func TestRaw(t *testing.T) {
	rawHTML := `<div class="test"><p>Test paragraph</p></div>`
	el := Raw(rawHTML)
	expected := rawHTML

	assert.Equal(t, expected, el.Render())
}

func TestCSS(t *testing.T) {
	cssContent := `.test-class {color: #333;}`
	expected := `.test-class {color: #333;}`
	el := CSS(cssContent)
	assert.Equal(t, expected, el.Render())
}

func TestSingleQuote(t *testing.T) {
	expected := `<div data-values='{"quantity": 5}'></div>`
	el := Div(attrs.Props{
		"data-values": `'{"quantity": 5}'`,
	})
	actual := el.Render()
	assert.Equal(t, expected, actual)
}

func TestFragment(t *testing.T) {
	expected := `<div><p>0</p><p>1</p><p>2</p><p>3</p><p>4</p></div>`
	nodes1 := []Node{
		P(nil,
			Text("1"),
		),
		P(nil,
			Text("2"),
		),
	}
	nodes2 := []Node{
		P(nil,
			Text("3"),
		),
		P(nil,
			Text("4"),
		),
	}
	el := Div(nil,
		P(nil,
			Text("0"),
		),
		Fragment(nodes1...),
		Fragment(nodes2...),
	)
	actual := el.Render()
	assert.Equal(t, expected, actual)
}
