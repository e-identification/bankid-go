package payload

import "github.com/e-identification/bankid/pkg/internal/http"

// AuthenticationPayload holds the required and optional fields of the authentication request.
type AuthenticationPayload struct {
	http.Payload `json:"-"`
	// The User IP address as seen by RP. String, IPv4 and IPv6 is allowed.
	EndUserIP string `validate:"ip" json:"endUserIp"`
	// Requirements on how the auth or sign order must be performed.
	Requirement *Requirement `json:"requirement,omitempty"`
	// The text to be displayed during authentication. The value must be base 64-encoded.
	// The text can be formatted using CR, LF and CRLF for new lines.
	// 1-1 500 characters after base 64 encoding.
	UserVisibleData UserDataString `validate:"omitempty,base64Length=1500" json:"userVisibleData,omitempty"`
	// Data is not displayed to the user. The value must be base 64-encoded.
	// 1-1 500 characters after base 64-encoding.
	UserNonVisibleData UserDataString `validate:"omitempty,base64Length=1500" json:"userNonVisibleData,omitempty"`
	// This parameter indicates that userVisibleData holds formatting characters.
	UserVisibleDataFormat string `validate:"omitempty,eq=simpleMarkdownV1" json:"userVisibleDataFormat,omitempty"`
}
