package stripe

import (
	"fmt"

	sgo "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// New creates a new subscription.
func (s *Stripe) SubscriptionNew(params *sgo.SubscriptionParams) (sub *sgo.Subscription, err error) {
	sub = &sgo.Subscription{}

	body := s.paramsToBody(params)

	err = s.request("POST", RouteSubscriptions, nil, body, sub)

	return
}

// Get returns the details of a subscription.
func (s *Stripe) SubscriptionGet(id string, params *sgo.SubscriptionParams) (sub *sgo.Subscription, err error) {
	path := fmt.Sprintf("%s/%s", RouteSubscriptions, id)

	sub = &sgo.Subscription{}

	err = s.request("GET", path, nil, nil, sub)

	return
}

// Update updates a subscription's properties.
func (s *Stripe) SubscriptionUpdate(id string, params *sgo.SubscriptionParams) (sub *sgo.Subscription, err error) {
	path := fmt.Sprintf("%s/%s", RouteSubscriptions, id)

	sub = &sgo.Subscription{}

	body := s.paramsToBody(params)

	err = s.request("POST", path, nil, body, sub)

	return
}

// Cancel removes a subscription.
func (s *Stripe) SubscriptionCancel(id string, params *sgo.SubscriptionCancelParams) (sub *sgo.Subscription, err error) {
	path := fmt.Sprintf("%s/%s", RouteSubscriptions, id)

	sub = &sgo.Subscription{}

	err = s.request("DELETE", path, nil, nil, sub)

	return
}

// List returns a list of subscriptions.
func (s *Stripe) List(listParams *sgo.SubscriptionListParams) *Iter {
	return &Iter{sgo.GetIter(listParams, func(p *sgo.Params, b *form.Values) ([]interface{}, sgo.ListMeta, error) {
		list := &sgo.SubscriptionList{}

		err := s.request("GET", RouteSubscriptions, nil, nil, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Subscription returns the subscription which the iterator is currently pointing to.
func (i *Iter) Subscription() *sgo.Subscription {
	return i.Current().(*sgo.Subscription)
}
