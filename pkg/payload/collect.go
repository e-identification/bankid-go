package payload

import "github.com/e-identification/bankid-go/pkg/internal/http"

// CollectPayload holds the required fields of the collect Payload.
type CollectPayload struct {
	http.Payload `json:"-"`
	// The orderRef from the response from authentication or sign.
	OrderRef string `json:"orderRef"`
}
