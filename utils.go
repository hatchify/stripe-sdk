package stripe

import (
	sgo "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Iter is an iterator for subscriptions.
type Iter struct {
	*sgo.Iter
}

func (s *Stripe) paramsToBody(params interface{}) (out []byte) {
	body := &form.Values{}
	form.AppendTo(body, params)

	var bodyString string

	if !body.Empty() {
		bodyString = body.Encode()
	}

	return []byte(bodyString)
}
