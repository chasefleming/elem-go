package attrs

const (
	// Universal Attributes
	Alt             = "alt"
	Class           = "class"
	Contenteditable = "contenteditable"
	Dir             = "dir" // Direction, e.g., "ltr" or "rtl"
	ID              = "id"
	Lang            = "lang"
	Style           = "style"
	Tabindex        = "tabindex"
	Title           = "title"

	// Link/Script Attributes
	Async       = "async"
	CrossOrigin = "crossorigin"
	Defer       = "defer"
	Href        = "href"
	Integrity   = "integrity"
	Rel         = "rel"
	Src         = "src"
	Target      = "target"

	// Meta Attributes
	Charset   = "charset"
	Content   = "content"
	HTTPequiv = "http-equiv" // e.g., for refresh or setting content type

	// Image/Embed Attributes
	Height = "height"
	Width  = "width"

	// Form/Input Attributes
	Accept         = "accept"
	Action         = "action"
	Autocapitalize = "autocapitalize"
	Autocomplete   = "autocomplete"
	Autofocus      = "autofocus"
	Cols           = "cols"
	Checked        = "checked"
	Disabled       = "disabled"
	For            = "for" // Used with label elements
	Max            = "max"
	MaxLength      = "maxlength"
	Method         = "method" // e.g., "GET", "POST"
	Min            = "min"
	Multiple       = "multiple"
	Name           = "name"
	Placeholder    = "placeholder"
	Readonly       = "readonly"
	Required       = "required"
	Rows           = "rows"
	Type           = "type"
	Value          = "value"

	// Miscellaneous Attributes
	DataPrefix = "data-" // Used for custom data attributes e.g., "data-custom"
	Download   = "download"
	Draggable  = "draggable"
	Role       = "role" // Used for ARIA roles
	Spellcheck = "spellcheck"

	// Table Attributes
	RowSpan	       = "rowspan"
	ColSpan	       = "colspan"
	Scope          = "scope"
	Headers        = "headers"
)

// list of supported boolean attributes
// see https://html.spec.whatwg.org/multipage/indices.html#attributes-3 for reference
var BooleanAttrs = map[string]struct{}{
	Async:     {},
	Autofocus: {},
	Checked:   {},
	Defer:     {},
	Disabled:  {},
	Multiple:  {},
	Readonly:  {},
	Required:  {},
}
