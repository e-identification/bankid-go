package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// BankIdIssueDate wraps time.Time and accepts multiple date/time encodings:
// - RFC3339 and RFC3339Nano timestamps (e.g., "2025-08-09T12:34:56Z")
// - ISO 8601 date only with 'Z' (e.g., "2025-08-09Z")
// - ISO 8601 date only with UTC offset (e.g., "2025-08-09+00:00")
// - Bare date (e.g., "2025-08-09"), interpreted as midnight UTC.
type BankIdIssueDate struct {
	time time.Time
}

// Time returns the underlying time value in UTC.
// If the receiver is nil or represents a zero value, it returns time.Time{}.
func (d *BankIdIssueDate) Time() time.Time {
	if d == nil {
		return time.Time{}
	}

	return d.time
}

// UnmarshalJSON implements json.Unmarshaler for BankIdIssueDate.
// It accepts the following formats and normalizes the result to UTC:
//   - RFC3339/RFC3339Nano timestamps (e.g., "2025-08-09T12:34:56Z").
//   - ISO 8601 date-only with 'Z' (e.g., "2025-08-09Z").
//   - ISO 8601 date-only with a UTC offset (e.g., "2025-08-09+00:00").
//   - Bare date (e.g., "2025-08-09"), which is interpreted as midnight UTC.
//
// If the input is null, it sets the zero time. For date-only inputs, the time
// is set to 00:00:00 UTC. For timestamp inputs, the parsed time is converted to UTC.
func (d *BankIdIssueDate) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		d.time = time.Time{}
		return nil
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	s = strings.TrimSpace(s)

	// Try multiple layouts
	layouts := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02Z07:00", // ISO 8601 date with zone (e.g., 2025-08-09+00:00)
		"2006-01-02Z",      // ISO 8601 date with Z (e.g., 2025-08-09Z)
		"2006-01-02",       // Date only
	}

	var (
		parsed time.Time
		err    error
	)

	for _, layout := range layouts {
		parsed, err = time.Parse(layout, s)
		if err == nil {
			// For date-only inputs (with or without zone), normalize to midnight UTC
			switch layout {
			case "2006-01-02", "2006-01-02Z", "2006-01-02Z07:00":
				parsed = time.Date(parsed.Year(), parsed.Month(), parsed.Day(), 0, 0, 0, 0, time.UTC)
			default:
				parsed = parsed.UTC()
			}

			d.time = parsed

			return nil
		}
	}

	// If nothing matched, return the last error from time.Parse
	return err
}

// MarshalJSON implements json.Marshaler for BankIdIssueDate.
// It returns "null" for a zero value, otherwise formats the time in RFC3339 (UTC).
func (d BankIdIssueDate) MarshalJSON() ([]byte, error) {
	if d.time.IsZero() {
		return []byte("null"), nil
	}

	return json.Marshal(d.time.UTC().Format(time.RFC3339))
}

// CompletionData holds the final state of an order.
type CompletionData struct {
	// Information related to the User
	User User `json:"user"`
	// Information related to the Device
	Device Device `json:"device"`
	// BankIDIssueDate is when the user's BankID was issued.
	// Accepts "YYYY-MM-DDZ", "YYYY-MM-DD+00:00", bare date "YYYY-MM-DD",
	// and full RFC3339/RFC3339Nano timestamps.
	BankIDIssueDate BankIdIssueDate `json:"bankIdIssueDate"`
	// Information about extra verifications that were part of the transaction.
	StepUp bool `json:"stepUp"`
	// The content of the signature is described in BankID Signature Profile specification. String. Base64-encoded
	Signature string `json:"signature"`
	// The OCSP response. String. Base64-encoded. The OCSP response is signed by a certificate that has the same issuer
	// as the certificate being verified. The OSCP response has an extension for Nonce
	OcspResponse string `json:"ocspResponse"`
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
	CompletionData CompletionData `json:"completionData"`
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
