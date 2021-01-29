package bankid

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

const (
	// StatusPending is the status of a pending order. hintCode describes the status of the order.
	StatusPending = "pending"
	// StatusComplete is the status of a complete order. completionData holds user information.
	StatusComplete = "complete"
	// StatusFailed is the status of a failed order. hintCode describes the error.
	StatusFailed = "failed"
	// HintCodeOutstandingTransaction is the hint for a order that is pending. The client has not yet received the order.
	// The hintCode will later change to noClient, started or userSign.
	HintCodeOutstandingTransaction = "outstandingTransaction"
	// HintCodeNoClient is the hint for a order that is pending. The client has not yet received the order.
	HintCodeNoClient = "noClient"
	// HintCodeStarted is the hint for a order that is pending. A Client has started with the 'autostarttoken' but a usable ID has not yet been found in the started client.
	// When the client start the may be a short delay until all ID:s are registered.
	// The user may not have any usable ID:s at all, or has not yet inserted their smart card.
	HintCodeStarted = "started"
	// HintCodeUserSign is the hint for a order that is pending. A client has received the order.
	HintCodeUserSign = "userSign"
	// HintCodeExpiredTransaction is the hint for a order that has expired.
	HintCodeExpiredTransaction = "expiredTransaction"
	// HintCodeCertificateError is a hint for when the provided certificate is invalid.
	HintCodeCertificateError = "certificateErr"
	// HintCodeUserCancel is a hint for when a user has cancelled a order.
	HintCodeUserCancel = "userCancel"
	// HintCodeCancelled is a hint for a order that has been cancelled.
	HintCodeCancelled = "cancelled"
	// HintCodeStartFailed is a hint for when a order could not be initialized.
	HintCodeStartFailed = "startFailed"
	// ErrorAlreadyInProgress is the hint for a order where an auth or sign request with personal number was sent, but an order for the user is already in progress. The order is aborted. No order is created.
	// Details are found in details.
	ErrorAlreadyInProgress = "alreadyInProgress"
)

// Response is the interface implemented by types that holds the response context fields.
type Response interface {
	Decode(subject io.ReadCloser, bankID *BankID) (Response, error)
}

// AuthenticateResponse contains the fields specific for the authentication response.
type AuthenticateResponse struct {
	// Used as reference to this order when the client is started automatically.
	AutoStartToken string `json:"autoStartToken"`
	// Used to collect the status of the order.
	OrderRef string `json:"orderRef"`
	// Used to compute the animated QR code.
	QrStartToken string `json:"qrStartToken"`
	// Used to compute the animated QR code.
	QrStartSecret string `json:"qrStartSecret"`
	bankID        *BankID
}

func (a *AuthenticateResponse) String() string {
	return fmt.Sprintf("%#v", a)
}

// Decode reads the JSON-encoded value and stories it in a authenticate response struct.
func (a *AuthenticateResponse) Decode(subject io.ReadCloser, bankID *BankID) (Response, error) {
	err := decode(subject, &a)
	if err != nil {
		return nil, err
	}

	a.bankID = bankID

	return a, nil
}

// Collect - Collects the result of a sign or auth order suing the orderRef as reference.
//
// RP should keep calling collect every two seconds as long as status indicates pending.
// RP must abort if status indicates failed. The user identity is returned when complete.
func (a AuthenticateResponse) Collect(context context.Context) (*CollectResponse, error) {
	return a.bankID.Collect(context, &CollectPayload{OrderRef: a.OrderRef})
}

// Cancel - Cancels an ongoing sign or auth order.
//
// This is typically used if the user cancels the order in your service or app.
func (a AuthenticateResponse) Cancel(context context.Context) (*CancelResponse, error) {
	return a.bankID.Cancel(context, &CancelPayload{OrderRef: a.OrderRef})
}

// SignResponse contains the fields specific for the sign response.
type SignResponse struct {
	AuthenticateResponse
}

// Decode reads the JSON-encoded value and stories it in a sign response struct.
func (s *SignResponse) Decode(subject io.ReadCloser, bankID *BankID) (Response, error) {
	err := decode(subject, &s)
	if err != nil {
		return nil, err
	}

	s.bankID = bankID

	return s, nil
}

// CollectResponse contains the fields specific for the collect response.
type CollectResponse struct {
	OrderRef       string         `json:"orderRef"`
	Status         string         `json:"status"`
	HintCode       string         `json:"hintCode"`
	CompletionData completionData `json:"completionData"`
}

func (c CollectResponse) String() string {
	return fmt.Sprintf("%#v", c)
}

// IsPending return true if the order is being processed. hintCode describes the status of the order.
func (c CollectResponse) IsPending() bool {
	return c.Status == StatusPending
}

// IsFailed return true if something went wrong with the order. hintCode describes the error.
func (c CollectResponse) IsFailed() bool {
	return c.Status == StatusFailed
}

// IsComplete return true if the order is complete. completionData holds user information.
func (c CollectResponse) IsComplete() bool {
	return c.Status == StatusComplete
}

// Decode reads the JSON-encoded value and stories it in a collect response struct.
func (c *CollectResponse) Decode(subject io.ReadCloser, bankID *BankID) (Response, error) {
	err := decode(subject, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// CancelResponse contrains fields for the cancel response.
type CancelResponse struct{}

// Decode reads the JSON-encoded value and stories it in a cancel response struct.
func (c *CancelResponse) Decode(subject io.ReadCloser, bankID *BankID) (Response, error) {
	err := decode(subject, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// ErrorResponse contrains fields fpr the error response.
type ErrorResponse struct {
	ErrorCode string `json:"errorCode"`
	Details   string `json:"details"`
	bankID    *BankID
}

func (e *ErrorResponse) String() string {
	return fmt.Sprintf("%#v", e)
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%s. %s", e.ErrorCode, e.Details)
}

// Decode reads the JSON-encoded value and stories it in a error response struct.
func (e *ErrorResponse) Decode(subject io.ReadCloser, bankID *BankID) (Response, error) {
	err := decode(subject, &e)
	if err != nil {
		return nil, err
	}

	e.bankID = bankID

	return e, nil
}

type completionData struct {
	// Information related to the user
	User user `json:"user"`
	// Information related to the device
	Device device `json:"device"`
	// Information related to the users certificate (BankID)
	Cert cert `  json:"cert"`
	// The content of the signature is described in BankID Signature Profile specification. String. Base64-encoded
	Signature string `json:"signature"`
	// The OCSP response. String. Base64-encoded. The OCSP response is signed by a certificate that has the same issuer
	// as the certificate being verified. The OSCP response has an extension for Nonce
	OcspResponse string `json:"ocspResponse"`
}

type user struct {
	// The personal number
	PersonalNumber string `json:"personalNumber"`
	// The given name and surname of the user
	Name string `json:"name"`
	// The given name of the user
	GivenName string `json:"givenName"`
	// The surname of the user
	Surname string `json:"surname"`
}

type device struct {
	// The IP address of the user agent as the BankID server discovers it.
	IPAddress string `json:"ipAddress"`
}

type cert struct {
	// Start of validity of the users BankID.
	NotBefore string `json:"notBefore"`
	// End of validity of the Users BankID.
	NotAfter string `json:"notAfter"`
}

func decode(subject io.ReadCloser, target interface{}) error {
	decoder := json.NewDecoder(subject)
	decoder.DisallowUnknownFields()

	return decoder.Decode(&target)
}
