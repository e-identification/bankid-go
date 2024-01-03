package response

// SignResponse contains the fields specific for the sign api response.
type SignResponse struct {
	AuthenticateResponse
}

// OnDecode is called on decode.
func (s *SignResponse) OnDecode() {
	// no op
}
