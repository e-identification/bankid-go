package response

import (
	"fmt"
	"time"
)

// AuthenticateResponse contains the fields specific for the authentication api response.
type AuthenticateResponse struct {
	// Used as reference to this order when the client is started automatically.
	AutoStartToken string `json:"autoStartToken"`
	// Used to collect the status of the order.
	OrderRef string `json:"orderRef"`
	// Used to compute the animated QR code.
	QrStartToken string `json:"qrStartToken"`
	// Used to compute the animated QR code.
	QrStartSecret string `json:"qrStartSecret"`
	// The time when response was returned.
	TimeOfResponse time.Time
}

func (a *AuthenticateResponse) String() string {
	return fmt.Sprintf("%#v", a)
}

// OnDecode is called on decode.
func (a *AuthenticateResponse) OnDecode() {
	a.TimeOfResponse = time.Now()
}
