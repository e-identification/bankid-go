package http

// Response is the interface implemented by types that holds the response context fields.
type Response interface {
	OnDecode()
}
