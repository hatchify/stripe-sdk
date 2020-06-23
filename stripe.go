package stripe

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
)

const (
	// Hostname of API
	Hostname = "https://api.stripe.com"
	// APIVersion Current API version
	APIVersion = "1.3"
)

const (
	// RouteGetAccount is the route for getting current account
	RouteGetAccount = "account"
)

// New will return a new instance of the Stripe API SDK
func New(apiKey string) (up *Stripe, err error) {
	var u Stripe
	if u.host, err = url.Parse(Hostname); err != nil {
		return
	}

	// Set API key
	u.apiKey = apiKey
	// Assign pointer reference
	up = &u
	return
}

// Stripe manages requests to the Stripe API
type Stripe struct {
	hc   http.Client
	host *url.URL

	// Login credentials
	apiKey string
}

func (u *Stripe) request(method, endpoint string, body io.Reader, resp interface{}) (err error) {
	var req *http.Request
	// Create a new request
	if req, err = u.newHTTPRequest(method, u.getURL(endpoint), body); err != nil {
		// Error encountered while creating new HTTP request, return
		return
	}

	var res *http.Response
	// Perform request using SDK's underlying HTTP client
	if res, err = u.hc.Do(req); err != nil {
		// Error encountered while performing request, return
		return
	}
	// Defer closing the HTTP response body
	defer res.Body.Close()

	// Process HTTP response from Stripe API
	return u.processResponse(res, resp)
}

func (u *Stripe) newHTTPRequest(method, url string, body io.Reader) (req *http.Request, err error) {
	// Create a new request using provided method, url, and body
	if req, err = http.NewRequest(method, url, body); err != nil {
		// Error encoutered while creating new HTTP request, return
		return
	}

	// Set API authentication using the username/password provided at SDK initialization
	req.SetBasicAuth(u.apiKey, "")
	return
}

func (u *Stripe) getURL(endpoint string) (url string) {
	// Create copy of host url.URL by derefencing source pointer
	reqURL := *u.host
	// Set the url path by concatinating the api version and the provided endpoint
	reqURL.Path = path.Join(APIVersion, endpoint)
	// Return the string representation of the built url
	return reqURL.String()
}

func (u *Stripe) processResponse(res *http.Response, value interface{}) (err error) {
	// Check to see if error was successful
	if res.StatusCode >= 400 {
		// Error status code encountered, process as error
		return u.processError(res.Body)
	}

	// Initialize new JSON decoder and attempt to decode as provided value
	err = json.NewDecoder(res.Body).Decode(&value)
	return
}

func (u *Stripe) processError(body io.Reader) (err error) {
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
