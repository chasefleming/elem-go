package htmx

const (
	// Request Modifiers
	HXGet    = "hx-get"
	HXPost   = "hx-post"
	HXPut    = "hx-put"
	HXDelete = "hx-delete"
	HXPatch  = "hx-patch"

	// Request Headers and Content-Type
	HXHeaders = "hx-headers"
	HXContent = "hx-content"

	// Request Parameters
	HXParams = "hx-params"
	HXValues = "hx-values"

	// Request Timeout and Retries
	HXTimeout      = "hx-timeout"
	HXRetry        = "hx-retry"
	HXRetryTimeout = "hx-retry-timeout"

	// Response Processing
	HXSwap    = "hx-swap"
	HXTarget  = "hx-target"
	HXSwapOOB = "hx-swap-oob"
	HXSelect  = "hx-select"
	HXExt     = "hx-ext"
	HXVals    = "hx-vals"

	// Events
	HXTrigger           = "hx-trigger"
	HXConfirm           = "hx-confirm"
	HXOn                = "hx-on"
	HXTriggeringElement = "hx-triggering-element"
	HXTriggeringEvent   = "hx-triggering-event"

	// Indicators
	HXIndicator = "hx-indicator"

	// History
	HXPushURL     = "hx-push-url"
	HXHistoryElt  = "hx-history-elt"
	HXHistoryAttr = "hx-history-attr"

	// Error Handling
	HXBoost = "hx-boost"
	HXError = "hx-error"

	// Caching
	HXCache = "hx-cache"
)
