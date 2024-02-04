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
	HXOnAbort                 = "hx-on-htmx-abort"
	HXOnAfterOnLoad           = "hx-on-htmx-after-on-load"
	HXOnAfterProcessNode      = "hx-on-htmx-after-process-node"
	HXOnAfterRequest          = "hx-on-htmx-after-request"
	HXOnAfterSettle           = "hx-on-htmx-after-settle"
	HXOnAfterSwap             = "hx-on-htmx-after-swap"
	HXOnBeforeCleanupElement  = "hx-on-htmx-before-cleanup-element"
	HXOnBeforeOnLoad          = "hx-on-htmx-before-onload"
	HXOnBeforeProcessNode     = "hx-on-htmx-before-process-node"
	HXOnBeforeRequest         = "hx-on-htmx-before-request"
	HXOnBeforeSwap            = "hx-on-htmx-before-swap"
	HXOnBeforeSend            = "hx-on-htmx-before-send"
	HXOnConfigRequest         = "hx-on-htmx-config-request"
	HXOnConfirm               = "hx-on-htmx-confirm"
	HXOnHistoryCacheError     = "hx-on-htmx-history-cache-error"
	HXOnHistoryCacheMiss      = "hx-on-htmx-history-cache-miss"
	HxOnHistoryCacheMissError = "hx-on-htmx-history-cache-miss-error"
	HXOnHistoryCacheMissLoad  = "hx-on-htmx-history-cache-miss-load"
	HXOnHistoryRestore        = "hx-on-htmx-history-restore"
	HxOnBeforeHistorySave     = "hx-on-htmx-before-history-save"
	HXOnLoad                  = "hx-on-htmx-load"
	HXOnNoSSESourceError      = "hx-on-htmx-no-sse-source-error"
	HXOnOnLoadError           = "hx-on-htmx-on-load-error"
	HXOnOOBAfterSwap          = "hx-on-htmx-oob-after-swap"
	HXOnOOBErrorNoTarget      = "hx-on-htmx-oob-error-no-target"
	HXOnPrompt                = "hx-on-htmx-prompt"
	HXOnPushedIntoHistory     = "hx-on-htmx-pushed-into-history"
	HXOnResponseError         = "hx-on-htmx-response-error"
	HXOnSendError             = "hx-on-htmx-send-error"
	HXOnSSEError              = "hx-on-htmx-sse-error"
	HXOnSSEOpen               = "hx-on-htmx-sse-open"
	HXOnSwapError             = "hx-on-htmx-swap-error"
	HXOnTargetError           = "hx-on-htmx-target-error"
	HXOnTimeout               = "hx-on-htmx-timeout"
	HXOnValidationValidate    = "hx-on-htmx-validation-validate"
	HXOnValidationFailed      = "hx-on-htmx-validation-failed"
	HXOnValidationHalted      = "hx-on-htmx-validation-halted"
	HXOnXHRAbort              = "hx-on-htmx-xhr-abort"
	HXOnXHRLoadend            = "hx-on-htmx-xhr-loadend"
	HXOnXHRLoadstart          = "hx-on-htmx-xhr-loadstart"
	HXOnXHRProgress           = "hx-on-htmx-xhr-progress"
)
