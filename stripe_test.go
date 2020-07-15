package stripe

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hatchify/requester"
	"github.com/hatchify/requester/mock"
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

	//Define what will be the backend for our mocks
	var backend = mock.NewFileBackend("testdata/stripe.json")

	//Setup Mock Requester
	var r requester.Interface
	if r, err = mock.NewSpyRequester(Hostname, backend); err != nil {
		t.Fatal(err)
	}

	s.SetRequester(r)

	return
}

func TestStripe_GetCharges(t *testing.T) {

	var err error
	s := setup(t)

	var out *interface{}
	if out, err = s.GetCharges(); err != nil {
		t.Fatal(err)
	}

	//debug
	fmt.Printf("%+v\n", *out)
}

func TestStripe_GetCustomers(t *testing.T) {

	var err error
	s := setup(t)

	var out *Customer
	if out, err = s.GetCustomers(); err != nil {
		t.Fatal(err)
	}

	//debug
	fmt.Printf("%+v\n", *out)
}

func TestStripe_CreateCustomer(t *testing.T) {

	var err error
	s := setup(t)

	var out *Customer
	if out, err = s.CreateCustomer(); err != nil {
		t.Fatal(err)
	}

	//debug
	fmt.Printf("%+v\n", *out)
}

func TestStripe_GetSubscriptionList(t *testing.T) {

	var err error
	s := setup(t)

	var out *SubscriptionsList
	if out, err = s.GetSubscriptionList(); err != nil {
		t.Fatal(err)
	}

	//debug
	fmt.Printf("%+v\n", *out)
}
