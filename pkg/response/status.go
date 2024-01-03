package response

// Status corresponds to a status of an order.
type Status string

const (
	// StatusPending is the status of a pending order. hintCode describes the status of the order.
	StatusPending = Status("pending")
	// StatusComplete is the status of a complete order. CompletionData holds User information.
	StatusComplete = Status("complete")
	// StatusFailed is the status of a failed order. hintCode describes the error.
	StatusFailed = Status("failed")
)
