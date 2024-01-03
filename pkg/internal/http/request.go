package http

// Payload is the interface implemented by types that holds the fields to be delivered to the API.
type Payload any

// Request holds the field related to http request.
type Request struct {
	URI           string
	Payload       Payload
	Response      Response
	ErrorResponse error
}
