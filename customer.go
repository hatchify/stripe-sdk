// Package customer provides the /customers APIs
package stripe

import (
	"fmt"

	"github.com/hatchify/requester"

	sgo "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

func (s *Stripe) paramsToOpts(params sgo.ParamsContainer) (out requester.Opts) {

	h := requester.NewHeaders()

	//for k, v := range params.Headers {
	//	for _, line := range v {
	//		fmt.Println(k, line)
	//
	//		h.Add(requester.Header{
	//			Key: k,
	//			Val: line,
	//		})
	//	}
	//}

	return requester.Opts{h}
}

func (s *Stripe) listParamsToOpts(params *sgo.CustomerListParams) requester.Opts {
	return requester.Opts{}
}

func (s *Stripe) paramsToBody(params *sgo.CustomerParams) (out []byte) {
	body := &form.Values{}
	form.AppendTo(body, params)

	var bodyString string

	if !body.Empty() {
		bodyString = body.Encode()
	}

	return []byte(bodyString)
}

// New creates a new customer.
func (s *Stripe) CustomerNew(params *sgo.CustomerParams) (cust *sgo.Customer, err error) {

	cust = &sgo.Customer{}

	opts := s.paramsToOpts(params)
	body := s.paramsToBody(params)

	err = s.request("POST", "customers", opts, body, cust)

	return
}

// Get returns the details of a customer.
func (s *Stripe) CustomerGet(id string, params *sgo.CustomerParams) (cust *sgo.Customer, err error) {
	path := fmt.Sprintf("customers/%s", id)

	cust = &sgo.Customer{}

	opts := s.paramsToOpts(params)

	err = s.request("GET", path, opts, nil, cust)

	return
}

// Update updates a customer's properties.
func (s *Stripe) CustomerUpdate(id string, params *sgo.CustomerParams) (cust *sgo.Customer, err error) {
	path := fmt.Sprintf("customers/%s", id)

	cust = &sgo.Customer{}

	opts := s.paramsToOpts(params)
	body := s.paramsToBody(params)

	err = s.request("POST", path, opts, body, cust)

	return
}

// Del removes a customer.
func (s *Stripe) CustomerDel(id string, params *sgo.CustomerParams) (cust *sgo.Customer, err error) {
	path := fmt.Sprintf("customers/%s", id)

	cust = &sgo.Customer{}

	opts := s.paramsToOpts(params)

	err = s.request("DELETE", path, opts, nil, cust)

	return
}

// List returns a list of customers.
func (s *Stripe) CustomerList(listParams *sgo.CustomerListParams) *Iter {
	return &Iter{sgo.GetIter(listParams, func(p *sgo.Params, b *form.Values) ([]interface{}, sgo.ListMeta, error) {
		list := &sgo.CustomerList{}

		opts := s.listParamsToOpts(listParams)
		err := s.request("GET", "customers", opts, nil, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for customers.
type Iter struct {
	*sgo.Iter
}

// Customer returns the customer which the iterator is currently pointing to.
func (i *Iter) Customer() *sgo.Customer {
	return i.Current().(*sgo.Customer)
}
