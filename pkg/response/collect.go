package response

import (
	"fmt"
	"time"
)

// CompletionData holds the final state of an order.
type CompletionData struct {
	// Information related to the User
	User User `json:"user"`
	// Information related to the Device
	Device Device `json:"device"`
	// The date the BankID was issued to the user
	BankIDIssueDate time.Time `json:"bankIdIssueDate"`
	// Information about extra verifications that were part of the transaction.
	StepUp bool `json:"stepUp"`
	// The content of the signature is described in BankID Signature Profile specification. String. Base64-encoded
	Signature string `json:"signature"`
	// The OCSP response. String. Base64-encoded. The OCSP response is signed by a certificate that has the same issuer
	// as the certificate being verified. The OSCP response has an extension for Nonce
	OcspResponse string `json:"ocspResponse"`

	// Indicates the risk level of the order based on data available in the order.
	//
	// The possible values have the following meaning:
	//
	// 	low: No or low risk identified in the available order data.
	// 	moderate: Might require further action, investigation or follow-up by you based on the order data.
	// 	high: The order should be blocked or cancelled by you and needs further action, investigation or follow-up. This value will only be returned if you have requested to have the risk assement to be provided, but not supplied a risk condition.
	//
	// Note: This is only returned if requested in the order, and it may be absent if the risk could not be calculated.
	// If you have sent the correct endUserIp and additional data, a risk indication with the value "high" means there are signs of the channel binding being compromised, or other highly concerning circumstances.
	Risk string `json:"risk"`
}

// User holds information related to the user.
type User struct {
	// The personal number
	PersonalNumber string `json:"personalNumber"`
	// The given name and surname of the User
	Name string `json:"name"`
	// The given name of the User
	GivenName string `json:"givenName"`
	// The surname of the User
	Surname string `json:"surname"`
}

// Device holds information related to the device.
type Device struct {
	// The IP address of the User agent as the BankID server discovers it.
	IPAddress string `json:"ipAddress"`
	// Unique hardware identifier for the user’s device.
	UHI string `json:"uhi"`
}

// Cert holds information related to the certificate.
type Cert struct {
	// Start of validity of the users BankID.
	NotBefore string `json:"notBefore"`
	// End of validity of the Users BankID.
	NotAfter string `json:"notAfter"`
}

// CollectResponse contains the fields specific for the collect api response.
type CollectResponse struct {
	OrderRef       string         `json:"orderRef"`
	Status         Status         `json:"status"`
	HintCode       string         `json:"hintCode"`
	CompletionData CompletionData `json:"CompletionData"`
}

func (c *CollectResponse) String() string {
	return fmt.Sprintf("%#v", c)
}

// IsPending return true if the order is being processed. hintCode describes the status of the order.
func (c *CollectResponse) IsPending() bool {
	return c.Status == StatusPending
}

// IsFailed return true if something went wrong with the order. hintCode describes the error.
func (c *CollectResponse) IsFailed() bool {
	return c.Status == StatusFailed
}

// IsComplete return true if the order is complete. CompletionData holds User information.
func (c *CollectResponse) IsComplete() bool {
	return c.Status == StatusComplete
}

// OnDecode is called on decode.
func (c *CollectResponse) OnDecode() {
	// no op
}
