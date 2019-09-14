package bankid

const (
	// The authentication URI
	UriAuth = "auth"
	// The sign URI
	UriSign = "sign"
	// The collect URI
	UriCollect = "collect"
	// The cancel URI
	UriCancel = "cancel"
)

type Request interface {
	Uri() string
	Payload() payloadInterface
	Response() Response
}

type request struct {
	uri      string
	payload  payloadInterface
	response Response
}

func (r request) Uri() string {
	return r.uri
}

func (r request) Payload() payloadInterface {
	return r.payload
}

func (r request) Response() Response {
	return r.response
}

// newAuthenticationRequest returns a new instance of 'Request'
func newAuthenticationRequest(payload *AuthenticationPayload) Request {
	return &request{uri: UriAuth, payload: payload, response: &AuthenticateResponse{}}
}

// newSignRequest returns a new instance of 'signRequest'
func newSignRequest(payload *SignPayload) Request {
	return &request{uri: UriSign, payload: payload, response: &SignResponse{}}
}

// newCollectRequest returns a new instance of 'CollectRequest'
func newCollectRequest(payload *CollectPayload) Request {
	return &request{uri: UriCollect, payload: payload, response: &CollectResponse{}}
}

// newCancelRequest returns a new instance of 'CancelRequest'
func newCancelRequest(payload *CancelPayload) Request {
	return &request{uri: UriCancel, payload: payload, response: &CancelResponse{}}
}
