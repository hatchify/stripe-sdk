# Stripe SDK [![GoDoc](https://godoc.org/github.com/hatchify/stripe-sdk?status.svg)](https://godoc.org/github.com/hatchify/stripe-sdk) ![Status](https://img.shields.io/badge/status-beta-yellow.svg)

Stripe SDK is an SDK wrapper for the Stripe API

## Usage 
### New
```go
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
```
