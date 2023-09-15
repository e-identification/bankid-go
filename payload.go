package bankid

// Payload is the interface implemented by types that holds the fields to be delivered to the API.
type Payload interface{}

// AuthenticationPayload holds the required and optional fields of the authentication request.
type AuthenticationPayload struct {
	Payload `json:"-"`
	// The user IP address as seen by RP. String, IPv4 and IPv6 is allowed.
	EndUserIP string `validate:"ip" json:"endUserIp"`
	// Requirements on how the auth or sign order must be performed.
	Requirement *Requirement `json:"requirement,omitempty"`
	// The text to be displayed during authentication. The value must be base 64-encoded.
	// The text can be formatted using CR, LF and CRLF for new lines.
	// 1-1 500 characters after base 64 encoding.
	UserVisibleData string `validate:"omitempty,base64Length=1500" json:"userVisibleData,omitempty"`
	// Data is not displayed to the user. The value must be base 64-encoded.
	// 1-1 500 characters after base 64-encoding.
	UserNonVisibleData string `validate:"omitempty,base64Length=1500" json:"userNonVisibleData,omitempty"`
	// This parameter indicates that userVisibleData holds formatting characters.
	UserVisibleDataFormat string `validate:"omitempty,eq=simpleMarkdownV1" json:"userVisibleDataFormat,omitempty"`
}

// CancelPayload holds the required fields of the collect Payload.
type CancelPayload struct {
	Payload `json:"-"`
	// The orderRef from the response from authentication or sign.
	OrderRef string `json:"orderRef"`
}

// CollectPayload holds the required fields of the collect Payload.
type CollectPayload struct {
	Payload `json:"-"`
	// The orderRef from the response from authentication or sign.
	OrderRef string `json:"orderRef"`
}

// SignPayload holds the required and optional fields for the sign Payload.
type SignPayload struct {
	Payload `json:"-"`
	// The user IP address as seen by RP. String, IPv4 and IPv6 is allowed.
	EndUserIP string `validate:"ip" json:"endUserIp"`
	// Requirements on how the auth or sign order must be performed.
	Requirement *Requirement `json:"requirement,omitempty"`
	// The text to be displayed and signed. The text can be formatted using CR, LF and CRLF for new lines.
	// 1-40 000 characters after base 64 encoding.
	UserVisibleData string `validate:"required,base64Length=40000" json:"userVisibleData"`
	// Data not displayed for the user.
	// 1-200 000 characters after base 64-encoding.
	UserNonVisibleData string `validate:"omitempty,base64Length=200000" json:"userNonVisibleData,omitempty"`
	// This parameter indicates that userVisibleData holds formatting characters.
	UserVisibleDataFormat string `validate:"omitempty,eq=simpleMarkdownV1" json:"userVisibleDataFormat,omitempty"`
}

// Requirement holds the required and optional fields of the Requirement DTO.
type Requirement struct {
	CardReader          string `validate:"omitempty,len=10" json:"cardReader,omitempty"`
	CertificatePolicies string `validate:"omitempty,len=10" json:"certificatePolicies,omitempty"`
	// If true. the client needs to provide MRTD (Machine readable travel document) information to
	// complete the order.
	Mrtd bool `json:"mrtd,omitempty"`
	// A personal identification number to be used to complete the transaction.
	PersonalNumber string `validate:"omitempty,numeric,len=12" json:"personalNumber,omitempty"`
	// If true, users are required to sign the transaction with their PIN code.
	PinCode bool `json:"pinCode,omitempty"`
}
