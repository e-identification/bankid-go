package payload

// PhoneRequirement holds the required and optional fields of the Requirement payload.
type PhoneRequirement struct {
	CardReader          string `validate:"omitempty,len=10" json:"cardReader,omitempty"`
	CertificatePolicies string `validate:"omitempty,len=10" json:"certificatePolicies,omitempty"`
	// If true, users are required to sign the transaction with their PIN code.
	PinCode bool `json:"pinCode,omitempty"`
}
