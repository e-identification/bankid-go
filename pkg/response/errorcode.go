package response

// ErrorCode corresponds to an error code returned by the RP API.
type ErrorCode string

const (
	// ErrorAlreadyInProgress is the status for when an auth or sign request with a personal number was sent,
	// but an order for the user is already in progress. The order is aborted. No order is created.
	ErrorAlreadyInProgress = ErrorCode("alreadyInProgress")
	// ErrorInvalidParameters is the status for when a invalid parameter was sent.
	ErrorInvalidParameters = ErrorCode("invalidParameters")
	// ErrorUnauthorized is the status for when RP does not have access to the service.
	ErrorUnauthorized = ErrorCode("unauthorized")
	// ErrorNotFound is the status for when an erroneous URL path was used.
	ErrorNotFound = ErrorCode("notFound")
	// ErrorMethodNotAllowed is the status for when an invalid method was used. Only http method POST is allowed.
	ErrorMethodNotAllowed = ErrorCode("methodNotAllowed")
	// ErrorRequestTimeout is the status for when it took to long time to transmit a request.
	ErrorRequestTimeout = ErrorCode("requestTimeout")
	// ErrorUnsupportedMediaType is the status for when an unsupported media type was provided.
	ErrorUnsupportedMediaType = ErrorCode("unsupportedMediaType")
	// ErrorInternalError is the status for when an internal technical error has occurred in the BankID system.
	ErrorInternalError = ErrorCode("internalError")
	// ErrorMaintenance is the status for when the BankID RP API is temporarily unavailable.
	ErrorMaintenance = ErrorCode("maintenance")
)
