package response

// CancelResponse contains fields for the cancel api response.
type CancelResponse struct{}

// OnDecode is called on decode.
func (c *CancelResponse) OnDecode() {
	// no op
}
