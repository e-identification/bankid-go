package bankid

import "encoding/json"

type Encoder interface {
	encode(payload Payload) ([]byte, error)
}

type jsonEncoder struct{}

func newJSONEncoder() Encoder {
	return &jsonEncoder{}
}

func (e jsonEncoder) encode(payload Payload) ([]byte, error) {
	return json.Marshal(payload)
}
