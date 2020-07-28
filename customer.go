// Package customer provides the /customers APIs
package stripe

import (
	"fmt"

	sgo "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// New creates a new customer.
func (s *Stripe) CustomerNew(params *sgo.CustomerParams) (cust *sgo.Customer, err error) {

	cust = &sgo.Customer{}

	body := s.paramsToBody(params)

	err = s.request("POST", RouteCustomers, nil, body, cust)

	return
}

// Get returns the details of a customer.
func (s *Stripe) CustomerGet(id string, params *sgo.CustomerParams) (cust *sgo.Customer, err error) {
	path := fmt.Sprintf("%s/%s", RouteCustomers, id)

	cust = &sgo.Customer{}

	err = s.request("GET", path, nil, nil, cust)

	return
}

// Update updates a customer's properties.
func (s *Stripe) CustomerUpdate(id string, params *sgo.CustomerParams) (cust *sgo.Customer, err error) {
	path := fmt.Sprintf("%s/%s", RouteCustomers, id)

	cust = &sgo.Customer{}

	body := s.paramsToBody(params)

	err = s.request("POST", path, nil, body, cust)

	return
}

// Del removes a customer.
func (s *Stripe) CustomerDel(id string, params *sgo.CustomerParams) (cust *sgo.Customer, err error) {
	path := fmt.Sprintf("%s/%s", RouteCustomers, id)

	cust = &sgo.Customer{}

	err = s.request("DELETE", path, nil, nil, cust)

	return
}

// List returns a list of customers.
func (s *Stripe) CustomerList(listParams *sgo.CustomerListParams) *Iter {
	return &Iter{sgo.GetIter(listParams, func(p *sgo.Params, b *form.Values) ([]interface{}, sgo.ListMeta, error) {
		list := &sgo.CustomerList{}

		err := s.request("GET", RouteCustomers, nil, nil, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Customer returns the customer which the iterator is currently pointing to.
func (i *Iter) Customer() *sgo.Customer {
	return i.Current().(*sgo.Customer)
}
