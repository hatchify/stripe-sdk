package stripe

import (
	"fmt"
	"testing"

	sgo "github.com/stripe/stripe-go/v71"
)

func TestStripe_CustomerFlowNewGetUpdateDelete(t *testing.T) {
	var err error
	s := setup(t)

	var name = "TestingFromGo"
	var params = &sgo.CustomerParams{Name: &name}

	var cust *sgo.Customer

	if cust, err = s.CustomerNew(params); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("val: %+v\n", cust.ID)

	//Let's get our creation
	var custGot *sgo.Customer
	if custGot, err = s.CustomerGet(cust.ID, params); err != nil {
		t.Fatal(err)
	}

	if cust.ID != custGot.ID {
		t.Fatal("mismatched ids")
	}

	//Let's update our monstrosity
	var custUpdated *sgo.Customer
	var updatedName = "TestingFromGo2"
	var updatedParams = &sgo.CustomerParams{Name: &updatedName}
	if custUpdated, err = s.CustomerUpdate(cust.ID, updatedParams); err != nil {
		t.Fatal(err)
	}

	if cust.Name == custUpdated.Name {
		//fmt.Printf("'%+v' -> '%+v'", cust.Name, custUpdated.Name)
		t.Fatal("name didn't update2")
	}

	if _, err = s.CustomerDel(cust.ID, params); err != nil {
		t.Fatal("didn't delete")
	}

}

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

	if _, err = s.CustomerGet("cus_HiykXILsFQrdhe111", params); err == nil {
		t.Fatal(err)
	}
}

//Testing for failure case
func TestStripe_CustomerUpdate(t *testing.T) {

	var err error
	s := setup(t)

	var params = &sgo.CustomerParams{}

	if _, err = s.CustomerUpdate("cus_HiykXILsFQrdhe", params); err == nil {
		t.Fatal(err)
	}
}

//Testing for failure case
func TestStripe_CustomerDel(t *testing.T) {

	var err error
	s := setup(t)

	var params = &sgo.CustomerParams{}

	if _, err = s.CustomerDel("cus_HiykXILsFQrdhe", params); err == nil {
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
