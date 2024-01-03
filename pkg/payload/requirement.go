package payload

// Requirement holds the required and optional fields of the Requirement payload.
type Requirement struct {
	CardReader          string `validate:"omitempty,len=10" json:"cardReader,omitempty"`
	CertificatePolicies string `validate:"omitempty,len=10" json:"certificatePolicies,omitempty"`
	// A personal identification number to be used to complete the transaction.
	PersonalNumber string `validate:"omitempty,numeric,len=12" json:"personalNumber,omitempty"`
	// If true. the client needs to provide MRTD (Machine readable travel document) information to
	// complete the order.
	Mrtd bool `json:"mrtd,omitempty"`
	// If true, users are required to sign the transaction with their PIN code.
	PinCode bool `json:"pinCode,omitempty"`
}
