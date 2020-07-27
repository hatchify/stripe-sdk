package stripe

import (
	"encoding/json"
	"io"
	"net/http"
	"path"

	"github.com/hatchify/requester"
	stripe "github.com/stripe/stripe-go/v71"
)

const (
	// Hostname of API
	Hostname = "https://api.stripe.com"
	// APIVersion Current API version
	APIVersion = "v1"
)

const (
	// RouteGetCharges is the route for getting all the charges
	RouteGetCharges = "charges"
	// RouteCustomers is the route for getting all customers
	RouteCustomers = "customers"
	// RouteSubscriptions is the route for getting all the subscriptions
	RouteSubscriptions = "subscriptions"
)

// New will return a new instance of the Stripe API SDK
func New(apiKey string) (s *Stripe, err error) {
	var u Stripe

	u.req = requester.New(&http.Client{}, stripe.APIURL)

	// Set API key
	u.apiKey = apiKey

	// Assign pointer reference
	s = &u
	return
}

func (s *Stripe) SetRequester(newReq requester.Interface) {
	s.req = newReq
}

// Stripe manages requests to the Stripe API
type Stripe struct {
	req requester.Interface

	// Login credentials
	apiKey string
}

func (s *Stripe) request(method, endpoint string, opts requester.Opts, body []byte, resp interface{}) (err error) {
	var res *http.Response

	// We authenticate with BasicAuth
	var setBasicAuth requester.Modifier = func(request *http.Request, client *http.Client) (err error) {
		request.SetBasicAuth(s.apiKey, "")
		return nil
	}

	// These content-type headers are needed for when we post things
	var setHeaders requester.Headers = requester.NewHeaders(requester.Header{
		Key: "Content-Type",
		Val: "application/json",
	})

	opts = append(opts, setBasicAuth, setHeaders)

	if res, err = s.req.Request(method, s.getURL(endpoint), body, opts); err != nil {
		return
	}
	// Defer closing the HTTP response body
	defer res.Body.Close()

	// Process HTTP response from Stripe API
	return s.processResponse(res, resp)
}

func (s *Stripe) getURL(endpoint string) (url string) {
	// Set the url path by concatenating the api version and the provided endpoint
	return path.Join(APIVersion, endpoint)
}

func (s *Stripe) processResponse(res *http.Response, value interface{}) (err error) {
	// Check to see if error was successful
	if res.StatusCode >= 400 {
		// Error status code encountered, process as error
		return s.processError(res.Body)
	}

	// Initialize new JSON decoder and attempt to decode as provided value
	if value != nil {
		err = json.NewDecoder(res.Body).Decode(&value)
	}
	return
}

func (s *Stripe) processError(body io.Reader) (err error) {
	var errResp errorResponse
	// Initialize new JSON decoder and attempt to decode as an error response
	if err = json.NewDecoder(body).Decode(&errResp); err != nil {
		// Error encountered while decoding, return
		return
	}

	// Set returning error as the error response's Error value
	err = errResp.Error
	return
}
