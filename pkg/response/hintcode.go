package response

// HintCode corresponds to a hint code returned by the collect endpoint.
type HintCode string

const (
	// HintCodeOutstandingTransaction is the hint for an order that is pending. The client has not yet received the order.
	// The hintCode will later change to noClient, started or userSign.
	HintCodeOutstandingTransaction = HintCode("outstandingTransaction")
	// HintCodeNoClient is the hint for an order that is pending. The client has not yet received the order.
	HintCodeNoClient = HintCode("noClient")
	// HintCodeStarted is the hint for an order that is pending. A Client has started with the 'autostarttoken'
	// but a usable ID has not yet been found in the started client.
	// When the client start the may be a short delay until all ID:s are registered.
	// The User may not have any usable ID:s at all, or has not yet inserted their smart card.
	HintCodeStarted = HintCode("started")
	// HintCodeUserSign is the hint for an order that is pending. A client has received the order.
	HintCodeUserSign = HintCode("userSign")
	// HintCodeExpiredTransaction is the hint for an order that has expired.
	HintCodeExpiredTransaction = HintCode("expiredTransaction")
	// HintCodeCertificateError is a hint for when the provided certificate is invalid.
	HintCodeCertificateError = HintCode("certificateErr")
	// HintCodeUserCancel is a hint for when a User has cancelled an order.
	HintCodeUserCancel = HintCode("userCancel")
	// HintCodeCancelled is a hint for an order that has been cancelled.
	HintCodeCancelled = HintCode("cancelled")
	// HintCodeStartFailed is a hint for when an order could not be initialized.
	HintCodeStartFailed = HintCode("startFailed")
)
