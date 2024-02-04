package htmx

const (
	// HTMX Reference: https://htmx.org/reference/
	// Request Modifiers
	HXGet    = "hx-get"
	HXPost   = "hx-post"
	HXPut    = "hx-put"
	HXDelete = "hx-delete"
	HXPatch  = "hx-patch"

	// Request Headers, Content-Type, additional data and request control
	HXHeaders  = "hx-headers"
	HXContent  = "hx-content"
	HXInclude  = "hx-include"
	HXRequest  = "hx-request"
	HXSync     = "hx-sync"
	HXValidate = "hx-validate"

	// Request Parameters
	HXParams = "hx-params"
	HXValues = "hx-values"

	// Request Timeout and Retries
	HXTimeout      = "hx-timeout"
	HXRetry        = "hx-retry"
	HXRetryTimeout = "hx-retry-timeout"

	// Response Processing
	HXSwap      = "hx-swap"
	HXTarget    = "hx-target"
	HXSwapOOB   = "hx-swap-oob"
	HXSelect    = "hx-select"
	HXSelectOOB = "hx-select-oob"
	HXExt       = "hx-ext"
	HXVals      = "hx-vals"

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
	HXReplaceURL  = "hx-replace-url"
	HXHistory     = "hx-history"
	HXHistoryElt  = "hx-history-elt"
	HXHistoryAttr = "hx-history-attr"

	// Error Handling
	HXBoost = "hx-boost"
	HXError = "hx-error"

	// Caching
	HXCache = "hx-cache"

	// HTMX Configuration
	HXDisable     = "hx-disable"
	HXDisabledElt = "hx-disabled-elt"
	HXDisinherit  = "hx-disinherit"
	HXEncoding    = "hx-encoding"
	HXExt         = "hx-ext"
	HXPreserve    = "hx-preserve"
	HXPrompt      = "hx-prompt"

	// HTMX Events, Reference: https://htmx.org/reference/#events
	// Reference for hx-on attribute: https://htmx.org/attributes/hx-on/
	HXOn                      = "hx-on"
	HXOnAbort                 = "hx-on--abort"
	HXOnAfterOnLoad           = "hx-on--after-on-load"
	HXOnAfterProcessNode      = "hx-on--after-process-node"
	HXOnAfterRequest          = "hx-on--after-request"
	HXOnAfterSettle           = "hx-on--after-settle"
	HXOnAfterSwap             = "hx-on--after-swap"
	HXOnBeforeCleanupElement  = "hx-on--before-cleanup-element"
	HXOnBeforeOnLoad          = "hx-on--before-on-load"
	HXOnBeforeProcessNode     = "hx-on--before-process-node"
	HXOnBeforeRequest         = "hx-on--before-request"
	HXOnBeforeSwap            = "hx-on--before-swap"
	HXOnBeforeSend            = "hx-on--before-send"
	HXOnConfigRequest         = "hx-on--config-request"
	HXOnConfirm               = "hx-on--confirm"
	HXOnHistoryCacheError     = "hx-on--history-cache-error"
	HXOnHistoryCacheMiss      = "hx-on--history-cache-miss"
	HXOnHistoryCacheMissError = "hx-on--history-cache-miss-error"
	HXOnHistoryCacheMissLoad  = "hx-on--history-cache-miss-load"
	HXOnHistoryRestore        = "hx-on--history-restore"
	HXOnBeforeHistorySave     = "hx-on--before-history-save"
	HXOnLoad                  = "hx-on--load"
	HXOnNoSSESourceError      = "hx-on--no-sse-source-error"
	HXOnOnLoadError           = "hx-on--on-load-error"
	HXOnOOBAfterSwap          = "hx-on--oob-after-swap"
	HXOnOOBErrorNoTarget      = "hx-on--oob-error-no-target"
	HXOnPrompt                = "hx-on--prompt"
	HXOnPushedIntoHistory     = "hx-on--pushed-into-history"
	HXOnResponseError         = "hx-on--response-error"
	HXOnSendError             = "hx-on--send-error"
	HXOnSSEError              = "hx-on--sse-error"
	HXOnSSEOpen               = "hx-on--sse-open"
	HXOnSwapError             = "hx-on--swap-error"
	HXOnTargetError           = "hx-on--target-error"
	HXOnTimeout               = "hx-on--timeout"
	HXOnValidationValidate    = "hx-on--validation-validate"
	HXOnValidationFailed      = "hx-on--validation-failed"
	HXOnValidationHalted      = "hx-on--validation-halted"
	HXOnXHRAbort              = "hx-on--xhr-abort"
	HXOnXHRLoadend            = "hx-on--xhr-loadend"
	HXOnXHRLoadstart          = "hx-on--xhr-loadstart"
	HXOnXHRProgress           = "hx-on--xhr-progress"
)
