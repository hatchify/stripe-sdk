package stripe

import (
	"fmt"
	"testing"

	sgo "github.com/stripe/stripe-go/v71"
)

func TestStripe_CustomerNew(t *testing.T) {

	var err error
	s := setup(t)

	var name = "TestingFromGo"
	var params = &sgo.CustomerParams{Description: &name}

	if _, err = s.CustomerNew(params); err != nil {
		t.Fatal(err)
	}

}

func TestStripe_CustomerGet(t *testing.T) {

	var err error
	s := setup(t)

	var params = &sgo.CustomerParams{}

	if _, err = s.CustomerGet("cus_HiykXILsFQrdhe", params); err != nil {
		t.Fatal(err)
	}
}

func TestStripe_CustomerUpdate(t *testing.T) {

	var err error
	s := setup(t)

	var params = &sgo.CustomerParams{}

	if _, err = s.CustomerUpdate("cus_HiykXILsFQrdhe", params); err != nil {
		t.Fatal(err)
	}
}

func TestStripe_CustomerDel(t *testing.T) {

	var err error
	s := setup(t)

	var params = &sgo.CustomerParams{}

	if _, err = s.CustomerDel("cus_HiykXILsFQrdhe", params); err != nil {
		t.Fatal(err)
	}
}

func TestStripe_CustomerList(t *testing.T) {
	s := setup(t)

	var params = &sgo.CustomerListParams{}

	i := s.CustomerList(params)

	i.Next()

	fmt.Println(i.Customer())
}
