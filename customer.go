package stripe

type Customer struct {
	ID                  string        `json:"id"`
	Object              string        `json:"object"`
	Address             interface{}   `json:"address"`
	Balance             int           `json:"balance"`
	Created             int           `json:"created"`
	Currency            string        `json:"currency"`
	DefaultSource       interface{}   `json:"default_source"`
	Delinquent          bool          `json:"delinquent"`
	Description         interface{}   `json:"description"`
	Discount            interface{}   `json:"discount"`
	Email               interface{}   `json:"email"`
	InvoicePrefix       string        `json:"invoice_prefix"`
	Livemode            bool          `json:"livemode"`
	Name                interface{}   `json:"name"`
	NextInvoiceSequence int           `json:"next_invoice_sequence"`
	Phone               interface{}   `json:"phone"`
	PreferredLocales    []interface{} `json:"preferred_locales"`
	Shipping            interface{}   `json:"shipping"`
	Subscriptions       Subscriptions `json:"subscriptions"`
	TaxExempt           string        `json:"tax_exempt"`
}

type Subscriptions struct {
	Object  string        `json:"object"`
	Data    []interface{} `json:"data"`
	HasMore bool          `json:"has_more"`
	URL     string        `json:"url"`
}
