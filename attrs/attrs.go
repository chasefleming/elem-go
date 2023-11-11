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
	Loading         = "loading"

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
	IsMap  = "ismap"

	// Semantic Text Attributes
	DateTime = "datetime"

	// Form/Input Attributes
	Accept         = "accept"
	Action         = "action"
	Autocapitalize = "autocapitalize"
	Autocomplete   = "autocomplete"
	Autofocus      = "autofocus"
	Cols           = "cols"
	Checked        = "checked"
	Disabled       = "disabled"
	For            = "for"
	Form           = "form"
	List           = "list"
	Low            = "low"
	High           = "high"
	Max            = "max"
	MaxLength      = "maxlength"
	Method         = "method" // e.g., "GET", "POST"
	Min            = "min"
	Multiple       = "multiple"
	Name           = "name"
	NoValidate     = "novalidate"
	Optimum        = "optimum"
	Placeholder    = "placeholder"
	Readonly       = "readonly"
	Required       = "required"
	Rows           = "rows"
	Selected       = "selected"
	Type           = "type"
	Value          = "value"

	// Interactive Attributes
	Open = "open"

	// Miscellaneous Attributes
	DataPrefix = "data-" // Used for custom data attributes e.g., "data-custom"
	Download   = "download"
	Draggable  = "draggable"
	Role       = "role" // Used for ARIA roles
	Spellcheck = "spellcheck"

	// Table Attributes
	RowSpan = "rowspan"
	ColSpan = "colspan"
	Scope   = "scope"
	Headers = "headers"

	// IFrame Attributes
	Allow           = "allow"
	AllowFullScreen = "allowfullscreen"
	CSP             = "csp"
	ReferrerPolicy  = "referrerpolicy"
	Sandbox         = "sandbox"
	SrcDoc          = "srcdoc"

	// Audio/Video Attributes
	Controls = "controls"
	Loop     = "loop"
	Muted    = "muted"
	Preload  = "preload"
	Autoplay = "autoplay"

	// Video-Specific Attributes
	Poster      = "poster"
	Playsinline = "playsinline"

	// Source Element-Specific Attributes
	Media = "media"
	Sizes = "sizes"
)

type Props map[string]string
