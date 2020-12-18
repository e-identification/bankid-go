package bankid

// Request is the interface implemented by types that holds the request context fields.
type Request interface {
	URI() string
	Payload() Payload
	Response() Response
}

type request struct {
	uri      string
	payload  Payload
	response Response
}

func (r request) URI() string {
	return r.uri
}

func (r request) Payload() Payload {
	return r.payload
}

func (r request) Response() Response {
	return r.response
}

// newAuthenticationRequest returns a new instance of 'Request'.
func newAuthenticationRequest(payload *AuthenticationPayload) Request {
	return &request{uri: "auth", payload: payload, response: &AuthenticateResponse{}}
}

// newSignRequest returns a new instance of 'signRequest'.
func newSignRequest(payload *SignPayload) Request {
	return &request{uri: "sign", payload: payload, response: &SignResponse{}}
}

// newCollectRequest returns a new instance of 'CollectRequest'.
func newCollectRequest(payload *CollectPayload) Request {
	return &request{uri: "collect", payload: payload, response: &CollectResponse{}}
}

// newCancelRequest returns a new instance of 'CancelRequest'.
func newCancelRequest(payload *CancelPayload) Request {
	return &request{uri: "cancel", payload: payload, response: &CancelResponse{}}
}
