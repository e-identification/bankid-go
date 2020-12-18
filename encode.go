package bankid

import "encoding/json"

type encoder interface {
	encode(payload Payload) ([]byte, error)
}

type jsonEncoder struct{}

func newJSONEncoder() encoder {
	return &jsonEncoder{}
}

func (e jsonEncoder) encode(payload Payload) ([]byte, error) {
	return json.Marshal(payload)
}
