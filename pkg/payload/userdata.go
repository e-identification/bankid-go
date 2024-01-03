package payload

import "encoding/base64"

// UserDataString holds the user data.
type UserDataString string

// UnmarshalJSON unmarshal a JSON into a UserDataString.
func (u *UserDataString) UnmarshalJSON(bytes []byte) error {
	result, err := base64.StdEncoding.DecodeString(string(bytes))
	if err != nil {
		return err // nolint:wrapcheck
	}

	userDataString := UserDataString(result) // nolint:staticcheck
	u = &userDataString                      // nolint:ineffassign,wastedassign

	return nil
}

// MarshalJSON marshals the type into JSON.
func (u UserDataString) MarshalJSON() ([]byte, error) {
	return []byte("\"" + base64.StdEncoding.EncodeToString([]byte(u)) + "\""), nil
}
