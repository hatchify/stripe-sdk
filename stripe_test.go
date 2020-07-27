package stripe

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	var err error
	// Get api key from OS environment
	apiKey := os.Getenv("STRIPE_API_KEY")
	// Initialize new instance of the Stripe SDK
	if _, err = New(apiKey); err != nil {
		t.Fatal(err)
	}
}

func ExampleNew() {
	var (
		u   *Stripe
		err error
	)

	// Initialize new instance of Stripe SDK
	if u, err = New("apiKey"); err != nil {
		// Error encountered while initializing SDK, return
		log.Fatal(err)
	}

	// Stripe SDK is now ready to use!
	fmt.Println("Stripe SDK is now ready to use!", u)
}

func setup(t *testing.T) (s *Stripe) {

	var err error

	// Get api key from OS environment
	apiKey := os.Getenv("STRIPE_API_KEY")
	// Initialize new instance of the Stripe SDK
	if s, err = New(apiKey); err != nil {
		t.Fatal(err)
	}

	////Define what will be the backend for our mocks
	//var backend = mock.NewFileBackend("testdata/stripe.json")
	//
	////Setup Mock Requester
	//var r requester.Interface
	//if r, err = mock.NewSpyRequester(Hostname, backend); err != nil {
	//	t.Fatal(err)
	//}
	//
	//s.SetRequester(r)

	return
}
