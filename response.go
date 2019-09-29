package bankid

import (
	"encoding/json"
	"fmt"
	"io"
)

const (
	// The order is being processed. hintCode describes the status of the order.
	StatusPending = "pending"
	// The order is complete. completionData holds user information
	StatusComplete = "failed"
	// Something went wrong with the order. hintCode describes the error.
	StatusFailed = "complete"
	// The order is pending. The client has not yet received the order.
	// The hintCode will later change to noClient, started or userSign.
	HintCodeOutstandingTransaction = "outstandingTransaction"
	// The order is pending. The client has not yet received the order.
	HintCodeNoClient = "noClient"
	// The order is pending. A Client has started with the 'autostarttoken' but a usable ID has not yet been found in the started client.
	// When the client start the may be a short delay until all ID:s are registered.
	// The user may not have any usable ID:s at all, or has not yet inserted their smart card.
	HintCodeStarted            = "started"
	HintCodeUserSign           = "userSign"
	HintCodeExpiredTransaction = "expiredTransaction"
	HintCodeCertificateError   = "certificateErr"
	HintCodeUserCancel         = "userCancel"
	HintCodeCancelled          = "cancelled"
	HintCodeStartFailed        = "startFailed"
	// An auth or sign request with personal number was sent, but an order for the user is already in progress. The order is aborted. No order is created.
	// Details are found in details.
	ErrorAlreadyInProgress = "alreadyInProgress"
)

type Envelope struct {
	ErrorCode string `json:"errorCode"`
	Details   string `json:"details"`
	Response
}

type Response interface {
	IsSuccess() bool
	Decode(subject io.ReadCloser, bankId *BankId) (Response, error)
}

// IsSuccess return true if the request was successful, false otherwise
func (e Envelope) IsSuccess() bool {
	return len(e.ErrorCode) <= 0
}

type AuthenticateResponse struct {
	Envelope
	// Used as reference to this order when the client is started automatically.
	AutoStartToken string `json:"autoStartToken"`
	// Used to collect the status of the order.
	OrderRef string `json:"orderRef"`
	bankId   *BankId
}

// Decode reads the JSON-encoded value and stories it in a authenticate response struct
func (a *AuthenticateResponse) Decode(subject io.ReadCloser, bankId *BankId) (Response, error) {
	err := decode(subject, &a)

	if err != nil {
		return nil, err
	}

	a.bankId = bankId

	return a, nil
}

// Collect - Collects the result of a sign or auth order suing the orderRef as reference.
//
// RP should keep calling collect every two seconds as long as status indicates pending.
// RP must abort if status indicates failed. The user identity is returned when complete.
func (a AuthenticateResponse) Collect() (*CollectResponse, error) {
	if !a.IsSuccess() {
		return nil, fmt.Errorf("action not applicable. Possible cause: %s %s", a.ErrorCode, a.Details)
	}

	return a.bankId.Collect(&CollectPayload{OrderRef: a.OrderRef})
}

// Cancel - Cancels an ongoing sign or auth order.
//
// This is typically used if the user cancels the order in your service or app.
func (a AuthenticateResponse) Cancel() (*CancelResponse, error) {
	if !a.IsSuccess() {
		return nil, fmt.Errorf("action not applicable. Possible cause: %s %s", a.ErrorCode, a.Details)
	}

	return a.bankId.Cancel(&CancelPayload{OrderRef: a.OrderRef})
}

type SignResponse struct {
	AuthenticateResponse
}

// Decode reads the JSON-encoded value and stories it in a sign response struct
func (s *SignResponse) Decode(subject io.ReadCloser, bankId *BankId) (Response, error) {
	err := decode(subject, &s)

	if err != nil {
		return nil, err
	}

	s.bankId = bankId

	return s, nil
}

type CollectResponse struct {
	Envelope
	OrderRef       string         `json:"orderRef"`
	Status         string         `json:"status"`
	HintCode       string         `json:"hintCode"`
	CompletionData CompletionData `json:"completionData"`
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

// IsAlreadyInProgress returns true if the order is already in progress.
func (c CollectResponse) IsAlreadyInProgress() bool {
	return c.ErrorCode == ErrorAlreadyInProgress
}

// Decode reads the JSON-encoded value and stories it in a collect response struct
func (c *CollectResponse) Decode(subject io.ReadCloser, bankId *BankId) (Response, error) {
	err := decode(subject, &c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

type CancelResponse struct {
	Envelope
}

// Decode reads the JSON-encoded value and stories it in a cancel response struct
func (c *CancelResponse) Decode(subject io.ReadCloser, bankId *BankId) (Response, error) {
	err := decode(subject, &c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

type CompletionData struct {
	// Information related to the user
	User User `json:"user"`
	// Information related to the device
	Device Device `json:"device"`
	// Information related to the users certificate (BankID)
	Cert Cert `  json:"cert"`
	// The content of the signature is described in BankID Signature Profile specification. String. Base64-encoded
	Signature string `json:"signature"`
	//The OCSP response. String. Base64-encoded. The OCSP response is signed by a certificate that has the same issuer
	//as the certificate being verified. The OSCP response has an extension for Nonce
	OcspResponse string `json:"ocspResponse"`
}

type User struct {
	// The personal number
	PersonalNumber string `json:"personalNumber"`
	// The given name and surname of the user
	Name string `json:"name"`
	// The given name of the user
	GivenName string `json:"givenName"`
	// The surname of the user
	Surname string `json:"surname"`
}

type Device struct {
	// The IP address of the user agent as the BankID server discovers it.
	IpAddress string `json:"ipAddress"`
}

type Cert struct {
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
