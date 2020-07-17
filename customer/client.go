// Package customer provides the /customers APIs
package customer

import (
	"github.com/hatchify/stripe-sdk/form"
)

// Client is used to invoke /customers APIs.
type Client struct {
	B   Stripe.Backend
	Key string
}

// New creates a new customer.
func New(params *Stripe.CustomerParams) (*Stripe.Customer, error) {
	return getC().New(params)
}

// New creates a new customer.
func (c Client) New(params *Stripe.CustomerParams) (*Stripe.Customer, error) {
	cust := &Stripe.Customer{}
	err := c.B.Call("POST", "/v1/customers", c.Key, params, cust)
	return cust, err
}

// Get returns the details of a customer.
func Get(id string, params *Stripe.CustomerParams) (*Stripe.Customer, error) {
	return getC().Get(id, params)
}

// Get returns the details of a customer.
func (c Client) Get(id string, params *Stripe.CustomerParams) (*Stripe.Customer, error) {
	path := Stripe.FormatURLPath("/v1/customers/%s", id)
	cust := &Stripe.Customer{}
	err := c.B.Call("GET", path, c.Key, params, cust)
	return cust, err
}

// Update updates a customer's properties.
func Update(id string, params *Stripe.CustomerParams) (*Stripe.Customer, error) {
	return getC().Update(id, params)
}

// Update updates a customer's properties.
func (c Client) Update(id string, params *Stripe.CustomerParams) (*Stripe.Customer, error) {
	path := Stripe.FormatURLPath("/v1/customers/%s", id)
	cust := &Stripe.Customer{}
	err := c.B.Call("POST", path, c.Key, params, cust)
	return cust, err
}

// Del removes a customer.
func Del(id string, params *Stripe.CustomerParams) (*Stripe.Customer, error) {
	return getC().Del(id, params)
}

// Del removes a customer.
func (c Client) Del(id string, params *Stripe.CustomerParams) (*Stripe.Customer, error) {
	path := Stripe.FormatURLPath("/v1/customers/%s", id)
	cust := &Stripe.Customer{}
	err := c.B.Call("DELETE", path, c.Key, params, cust)
	return cust, err
}

// List returns a list of customers.
func List(params *Stripe.CustomerListParams) *Iter {
	return getC().List(params)
}

// List returns a list of customers.
func (c Client) List(listParams *Stripe.CustomerListParams) *Iter {
	return &Iter{Stripe.GetIter(listParams, func(p *Stripe.Params, b *form.Values) ([]interface{}, Stripe.ListMeta, error) {
		list := &Stripe.CustomerList{}
		err := c.B.CallRaw("GET", "/v1/customers", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for customers.
type Iter struct {
	*Stripe.Iter
}

// Customer returns the customer which the iterator is currently pointing to.
func (i *Iter) Customer() *Stripe.Customer {
	return i.Current().(*Stripe.Customer)
}

func getC() Client {
	return Client{Stripe.GetBackend(Stripe.APIBackend), Stripe.Key}
}
