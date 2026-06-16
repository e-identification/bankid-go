package response

// PhoneSignResponse contains the fields specific for the phone sign api response.
type PhoneSignResponse struct {
	PhoneAuthenticateResponse
}

// OnDecode is called on decode.
func (s *PhoneSignResponse) OnDecode() {
	s.PhoneAuthenticateResponse.OnDecode()
}
