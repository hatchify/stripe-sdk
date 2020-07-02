package stripe

import "fmt"

// Error represents an Stripe API error response payload
type Error struct {
	// The type of error returned. One of:
	//	- api_connection_error
	//	- api_error
	//	- authentication_error
	//	- card_error
	//	- idempotency_error
	//	- invalid_request_error
	//	- rate_limit_error
	Type string `json:"type,omitempty"`
	// For some errors that could be handled programmatically, a short string indicating the error code reported.
	Code string `json:"code,omitempty"`
	// For card errors resulting from a card issuer decline, a short string indicating the card issuer’s
	// reason for the decline if they provide one.
	DeclineCode string `json:"decline_code,omitempty"`
	// A human-readable message providing more details about the error.
	// For card errors, these messages can be shown to your users.
	Message string `json:"message,omitempty"`
	// If the error is parameter-specific, the parameter related to the error.
	// For example, you can use this to display a message near the correct form field.
	Param string `json:"param,omitempty"`
	// The PaymentIntent object for errors returned on a request involving a PaymentIntent.
	PaymentIntent string `json:"payment_intent,omitempty"`
}

// Error will return the error representation of the Error
// Note: This causes Error to match the error interface
func (e *Error) Error() string {
	// Return a formatted version of the error message and code
	return fmt.Sprintf("%s [%s]", e.Message, e.Code)
}

// errorResponse is a response wrapper to match the Stripe API payload
type errorResponse struct {
	Error *Error `json:"error"`
}
