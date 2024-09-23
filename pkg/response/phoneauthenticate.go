package response

import (
	"fmt"
)

// PhoneAuthenticateResponse contains the fields specific for the phone authentication api response.
type PhoneAuthenticateResponse struct {
	// Used to collect the status of the order.
	OrderRef string `json:"orderRef"`
}

func (a *PhoneAuthenticateResponse) String() string {
	return fmt.Sprintf("%#v", a)
}

// OnDecode is called on decode.
func (a *PhoneAuthenticateResponse) OnDecode() {
	// no op
}
