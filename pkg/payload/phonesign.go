package payload

import "github.com/NicklasWallgren/bankid/v2/pkg/internal/http"

// PhoneSignPayload holds the required and optional fields for the phone sign payload.
type PhoneSignPayload struct {
	http.Payload `json:"-"`
	// A personal identification number to be used to complete the transaction.
	PersonalNumber string `validate:"omitempty,numeric,len=12" json:"personalNumber,omitempty"`
	// Indicate if the user or the RP initiated the phone call.
	CallInitiator string `validate:"oneof=user RP" json:"callInitiator"`
	// Requirements on how the auth or sign order must be performed.
	Requirement *PhoneRequirement `json:"requirement,omitempty"`
	// The text to be displayed and signed. The text can be formatted using CR, LF and CRLF for new lines.
	// 1--40 000 characters after base 64 encoding.
	UserVisibleData UserDataString `validate:"required,base64Length=40000" json:"userVisibleData"`
	// Data not displayed for the User.
	// 1-200 000 characters after base 64-encoding.
	UserNonVisibleData UserDataString `validate:"omitempty,base64Length=200000" json:"userNonVisibleData,omitempty"`
	// This parameter indicates that userVisibleData holds formatting characters.
	UserVisibleDataFormat string `validate:"omitempty,eq=simpleMarkdownV1" json:"userVisibleDataFormat,omitempty"`
}
