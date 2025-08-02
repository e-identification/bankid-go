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
	// HintCodeUserMrtd is the hint for an order that is pending.
	// A client has launched and received the order but additional steps
	// for providing MRTD information is required to proceed with the order.
	HintCodeUserMrtd = HintCode("userMrtd")
	// Order is waiting for the user to confirm that they have received this
	// order while in a call with your organization.
	HintCodeUserCallConfirm = HintCode("userCallConfirm")
	// HintCodeUserSign is the hint for an order that is pending. A client has received the order.
	HintCodeUserSign = HintCode("userSign")
	// HintCodeProcessing is the hint for when the order is pending. The client has received
	// and signed the order, but signature processing is underway.
	HintCodeProcessing = HintCode("processing")
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
	// HintCodeUserDeclinedCall is a hint for when the order was cancelled because
	// the user indicated in the app that they are not in a call with your organization.
	HintCodeUserDeclinedCall = HintCode("userDeclinedCall")
	// HintCodeNotSupportedByUserApp is a hint for an order that was picked up by a
	// client that does not support the requested feature.
	// The BankID client used by the user needs to be updated to a later version that
	// supports the features that are required to complete the order.
	HintCodeNotSupportedByUserApp = HintCode("notSupportedByUserApp")
	// HintCodeTransactionRiskBlocked is a hint for when the risk for the
	// order was too high and the order was blocked.
	HintCodeTransactionRiskBlocked = HintCode("transactionRiskBlocked")
)
