package htmx

// HTMX attribute reference: https://htmx.org/reference/
// This package targets htmx 2.x. Constants that only worked in htmx 1.x, or
// that never corresponded to a real htmx attribute, are kept for backwards
// compatibility in the deprecated block at the bottom of this file.
const (
	// Core Attributes
	HXGet       = "hx-get"
	HXPost      = "hx-post"
	HXPushURL   = "hx-push-url"
	HXSelect    = "hx-select"
	HXSelectOOB = "hx-select-oob"
	HXSwap      = "hx-swap"
	HXSwapOOB   = "hx-swap-oob"
	HXTarget    = "hx-target"
	HXTrigger   = "hx-trigger"
	HXVals      = "hx-vals"

	// Additional Attributes
	HXBoost       = "hx-boost"
	HXConfirm     = "hx-confirm"
	HXDelete      = "hx-delete"
	HXDisable     = "hx-disable"
	HXDisabledElt = "hx-disabled-elt"
	HXDisinherit  = "hx-disinherit"
	HXEncoding    = "hx-encoding"
	HXExt         = "hx-ext"
	HXHeaders     = "hx-headers"
	HXHistory     = "hx-history"
	HXHistoryElt  = "hx-history-elt"
	HXInclude     = "hx-include"
	HXIndicator   = "hx-indicator"
	HXInherit     = "hx-inherit"
	HXParams      = "hx-params"
	HXPatch       = "hx-patch"
	HXPreserve    = "hx-preserve"
	HXPrompt      = "hx-prompt"
	HXPut         = "hx-put"
	HXReplaceURL  = "hx-replace-url"
	HXRequest     = "hx-request"
	HXSync        = "hx-sync"
	HXValidate    = "hx-validate"

	// Server Sent Events extension attributes
	// Reference: https://github.com/bigskysoftware/htmx-extensions/blob/main/src/sse/README.md
	SSEConnect = "sse-connect"
	SSESwap    = "sse-swap"
	SSEClose   = "sse-close"

	// WebSockets extension attributes
	// Reference: https://github.com/bigskysoftware/htmx-extensions/blob/main/src/ws/README.md
	WSConnect = "ws-connect"
	WSSend    = "ws-send"

	// HTMX Events, Reference: https://htmx.org/reference/#events
	// Reference for hx-on attributes: https://htmx.org/attributes/hx-on/
	HXOnAbort                     = "hx-on--abort"
	HXOnAfterOnLoad               = "hx-on--after-on-load"
	HXOnAfterProcessNode          = "hx-on--after-process-node"
	HXOnAfterRequest              = "hx-on--after-request"
	HXOnAfterSettle               = "hx-on--after-settle"
	HXOnAfterSwap                 = "hx-on--after-swap"
	HXOnBeforeCleanupElement      = "hx-on--before-cleanup-element"
	HXOnBeforeOnLoad              = "hx-on--before-on-load"
	HXOnBeforeProcessNode         = "hx-on--before-process-node"
	HXOnBeforeRequest             = "hx-on--before-request"
	HXOnBeforeSwap                = "hx-on--before-swap"
	HXOnBeforeSend                = "hx-on--before-send"
	HXOnBeforeTransition          = "hx-on--before-transition"
	HXOnConfigRequest             = "hx-on--config-request"
	HXOnConfirm                   = "hx-on--confirm"
	HXOnHistoryCacheError         = "hx-on--history-cache-error"
	HXOnHistoryCacheHit           = "hx-on--history-cache-hit"
	HXOnHistoryCacheMiss          = "hx-on--history-cache-miss"
	HXOnHistoryCacheMissLoadError = "hx-on--history-cache-miss-load-error"
	HXOnHistoryCacheMissLoad      = "hx-on--history-cache-miss-load"
	HXOnHistoryRestore            = "hx-on--history-restore"
	HXOnBeforeHistorySave         = "hx-on--before-history-save"
	HXOnLoad                      = "hx-on--load"
	HXOnNoSSESourceError          = "hx-on--no-sse-source-error"
	HXOnOnLoadError               = "hx-on--on-load-error"
	HXOnOOBAfterSwap              = "hx-on--oob-after-swap"
	HXOnOOBBeforeSwap             = "hx-on--oob-before-swap"
	HXOnOOBErrorNoTarget          = "hx-on--oob-error-no-target"
	HXOnPrompt                    = "hx-on--prompt"
	HXOnPushedIntoHistory         = "hx-on--pushed-into-history"
	HXOnReplacedInHistory         = "hx-on--replaced-in-history"
	HXOnResponseError             = "hx-on--response-error"
	HXOnSendAbort                 = "hx-on--send-abort"
	HXOnSendError                 = "hx-on--send-error"
	HXOnSSEError                  = "hx-on--sse-error"
	HXOnSSEOpen                   = "hx-on--sse-open"
	HXOnSwapError                 = "hx-on--swap-error"
	HXOnTargetError               = "hx-on--target-error"
	HXOnTimeout                   = "hx-on--timeout"
	HXOnValidationValidate        = "hx-on--validation-validate"
	HXOnValidationFailed          = "hx-on--validation-failed"
	HXOnValidationHalted          = "hx-on--validation-halted"
	HXOnXHRAbort                  = "hx-on--xhr-abort"
	HXOnXHRLoadend                = "hx-on--xhr-loadend"
	HXOnXHRLoadstart              = "hx-on--xhr-loadstart"
	HXOnXHRProgress               = "hx-on--xhr-progress"
)

// Deprecated constants, kept so existing code keeps compiling. They will be
// removed in a future release.
const (
	// Deprecated: hx-content is not an htmx attribute and has no effect.
	HXContent = "hx-content"

	// Deprecated: hx-values is not an htmx attribute and has no effect. Use HXVals instead.
	HXValues = "hx-values"

	// Deprecated: hx-timeout is not an htmx attribute and has no effect.
	// Set a timeout via HXRequest, e.g. hx-request='{"timeout": 1000}'.
	HXTimeout = "hx-timeout"

	// Deprecated: hx-retry is not an htmx attribute and has no effect.
	HXRetry = "hx-retry"

	// Deprecated: hx-retry-timeout is not an htmx attribute and has no effect.
	HXRetryTimeout = "hx-retry-timeout"

	// Deprecated: hx-triggering-element is not an htmx attribute and has no effect.
	HXTriggeringElement = "hx-triggering-element"

	// Deprecated: hx-triggering-event is not an htmx attribute and has no effect.
	HXTriggeringEvent = "hx-triggering-event"

	// Deprecated: hx-history-attr is not an htmx attribute and has no effect. Use HXHistory or HXHistoryElt instead.
	HXHistoryAttr = "hx-history-attr"

	// Deprecated: hx-error is not an htmx attribute and has no effect.
	HXError = "hx-error"

	// Deprecated: hx-cache is not an htmx attribute and has no effect.
	HXCache = "hx-cache"

	// Deprecated: hx-sse was removed in htmx 2.0. Use the SSE extension
	// attributes SSEConnect, SSESwap, and SSEClose instead.
	HXSSE = "hx-sse"

	// Deprecated: hx-ws was removed in htmx 2.0. Use the WebSockets extension
	// attributes WSConnect and WSSend instead.
	HXWS = "hx-ws"

	// Deprecated: the bare hx-on attribute was removed in htmx 2.0. Use the
	// per-event HXOn* constants (e.g. HXOnBeforeRequest) instead.
	HXOn = "hx-on"

	// Deprecated: this constant was misnamed; the htmx event is
	// htmx:historyCacheMissLoadError. Use HXOnHistoryCacheMissLoadError instead.
	HXOnHistoryCacheMissError = "hx-on--history-cache-miss-load-error"
)
