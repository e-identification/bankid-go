package payload

import "github.com/NicklasWallgren/bankid/v2/pkg/internal/http"

// SignPayload holds the required and optional fields for the sign payload.
type SignPayload struct {
	http.Payload `json:"-"`
	// The User IP address as seen by RP. String, IPv4 and IPv6 is allowed.
	EndUserIP string `validate:"ip" json:"endUserIp"`
	// Requirements on how the auth or sign order must be performed.
	Requirement *Requirement `json:"requirement,omitempty"`
	// The text to be displayed and signed. The text can be formatted using CR, LF and CRLF for new lines.
	// 1--40 000 characters after base 64 encoding.
	UserVisibleData UserDataString `validate:"required,base64Length=40000" json:"userVisibleData"`
	// Data not displayed for the User.
	// 1-200 000 characters after base 64-encoding.
	UserNonVisibleData UserDataString `validate:"omitempty,base64Length=200000" json:"userNonVisibleData,omitempty"`
	// This parameter indicates that userVisibleData holds formatting characters.
	UserVisibleDataFormat string `validate:"omitempty,eq=simpleMarkdownV1" json:"userVisibleDataFormat,omitempty"`
}
