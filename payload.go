package bankid

import (
	"encoding/base64"
	"encoding/json"
)

// payloadInterface is the interface implemented by types that holds the fields to be delivered to the API
type payloadInterface interface{}

// payload holds the request fields to be delivered to the API
type payload struct {
	payloadInterface
}

// Requirements holds the required and optional fields of the Requirement DTO
type Requirement struct {
	CardReader             string `validate:"omitempty,len=10" json:"cardReader,omitempty"`
	CertificatePolicies    string `validate:"omitempty,len=10" json:"certificatePolicies,omitempty"`
	IssuerCn               string `validate:"omitempty,len=10" json:"issuerCn,omitempty"`
	AutoStartTokenRequired bool   `json:"autoStartTokenRequired,omitempty"`
	AllowFingerprint       bool   `json:"allowFingerprint,omitempty"`
}

// AuthenticationPayload holds the required and optional fields of the authentication request
type AuthenticationPayload struct {
	*payload
	// The personal number of the user. String 12 digits. Century must be included.
	// If the personal number is excluded, the client must be started with
	// the autoStartToken returned in the response.
	PersonalNumber string `validate:"numeric" json:"personalNumber"`
	// The user IP address as seen by RP. String, IPv4 and IPv6 is allowed.
	EndUserIp string `validate:"ip" json:"endUserIp"`
	// Requirements on how the auth or sign order must be performed.
	Requirement *Requirement `json:"requirement,omitempty"`
}

// SignPayload holds the required and optional fields for the sign payload
type SignPayload struct {
	*payload
	// The personal number of the user. String 12 digits. Century must be included.
	// If the personal number is excluded, the client must be started with
	// the autoStartToken returned in the response.
	PersonalNumber string `validate:"numeric,len=12" json:"personalNumber"`
	// The user IP address as seen by RP. String, IPv4 and IPv6 is allowed.
	EndUserIp string `validate:"ip" json:"endUserIp"`
	// The text to be displayed and signed. The text can be formatted using CR, LF and CRLF for new lines.
	// 1--40 000 characters after base 64 encoding.
	UserVisibleData string `validate:"required,base64Length=40000" json:"userVisibleData"`
	// Data not displayed for the user.
	// 1-200 000 characters after base 64-encoding.
	UserNonVisibleData string `json:"userNonVisibleData,omitempty"`
	// Requirements on how the auth or sign order must be performed.
	Requirement *Requirement `json:"requirement,omitempty"`
}

// MarshalJSON returns a JSON encoded 'SignPayload'
func (s SignPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		PersonalNumber     string       `json:"personalNumber"`
		EndUserIp          string       `json:"endUserIp"`
		UserVisibleData    string       `json:"userVisibleData"`
		UserNonVisibleData string       `json:"userNonVisibleData,omitempty"`
		Requirement        *Requirement `json:"requirement,omitempty"`
	}{
		PersonalNumber:     s.PersonalNumber,
		EndUserIp:          s.EndUserIp,
		UserVisibleData:    base64.StdEncoding.EncodeToString([]byte(s.UserVisibleData)),
		UserNonVisibleData: base64.StdEncoding.EncodeToString([]byte(s.UserNonVisibleData)),
		Requirement:        s.Requirement,
	})
}

// CollectPayload holds the required fields of the collect payload
type CollectPayload struct {
	*payload
	// The orderRef from the response from authentication or sign.
	OrderRef string `json:"orderRef"`
}

// CancelPayload holds the required fields of the collect payload
type CancelPayload struct {
	*payload
	// The orderRef from the response from authentication or sign.
	OrderRef string `json:"orderRef"`
}
