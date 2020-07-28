package stripe

import (
	"testing"
	"time"

	sgo "github.com/stripe/stripe-go/v71"
)

func TestStripe_SubscriptionNew(t *testing.T) {

	var err error
	s := setup(t)

	var name = "TestingFromGo"
	var params = &sgo.CustomerParams{Description: &name}
	var cust *sgo.Customer

	if cust, err = s.CustomerNew(params); err != nil {
		t.Fatal(err)
	}

	custID := &cust.ID

	var subParams = &sgo.SubscriptionParams{
		Customer: custID,
		AddInvoiceItems: []*sgo.SubscriptionAddInvoiceItemParams{
			{
				Price:    sgo.String("price_123"),
				Quantity: sgo.Int64(2),
			},
			{
				PriceData: &sgo.InvoiceItemPriceDataParams{
					Currency:   sgo.String(string(sgo.CurrencyUSD)),
					UnitAmount: sgo.Int64(1000),
					Product:    sgo.String("prod_123"),
				},
				Quantity: sgo.Int64(4),
			},
		},
		BillingCycleAnchor: sgo.Int64(time.Now().AddDate(0, 0, 12).Unix()),
		CollectionMethod:   sgo.String(string(sgo.SubscriptionCollectionMethodChargeAutomatically)),
		DaysUntilDue:       sgo.Int64(30),
		Items: []*sgo.SubscriptionItemsParams{
			{
				Price:    sgo.String("price_ABC"),
				Quantity: sgo.Int64(10),
			},
			{
				PriceData: &sgo.SubscriptionItemPriceDataParams{
					Currency: sgo.String(string(sgo.CurrencyUSD)),
					Product:  sgo.String("prod_ABC"),
					Recurring: &sgo.SubscriptionItemPriceDataRecurringParams{
						Interval:      sgo.String(string(sgo.PriceRecurringIntervalMonth)),
						IntervalCount: sgo.Int64(6),
					},
					UnitAmount: sgo.Int64(1000),
				},
			},
		},
	}

	if _, err = s.SubscriptionNew(subParams); err != nil {
		t.Fatal(err)
	}
}

////TODO: Implement the rest
//func TestStripe_SubscriptionFlowNewGetUpdateDelete(t *testing.T) {
//	var err error
//	s := setup(t)
//
//	var name = "TestingFromGo"
//	var params = &sgo.CustomerParams{Name: &name}
//
//	var cust *sgo.Customer
//
//	if cust, err = s.CustomerNew(params); err != nil {
//		t.Fatal(err)
//	}
//
//	fmt.Printf("val: %+v\n", cust.ID)
//
//	//Let's get our creation
//	var custGot *sgo.Customer
//	if custGot, err = s.CustomerGet(cust.ID, params); err != nil {
//		t.Fatal(err)
//	}
//
//	if cust.ID != custGot.ID {
//		t.Fatal("mismatched ids")
//	}
//
//	//Let's update our monstrosity
//	var custUpdated *sgo.Customer
//	var updatedName = "TestingFromGo2"
//	var updatedParams = &sgo.CustomerParams{Name: &updatedName}
//	if custUpdated, err = s.CustomerUpdate(cust.ID, updatedParams); err != nil {
//		t.Fatal(err)
//	}
//
//	if cust.Name == custUpdated.Name {
//		//fmt.Printf("'%+v' -> '%+v'", cust.Name, custUpdated.Name)
//		t.Fatal("name didn't update2")
//	}
//
//	if _, err = s.CustomerDel(cust.ID, params); err != nil {
//		t.Fatal("didn't delete")
//	}
//
//}
//func TestStripe_SubscriptionGet(t *testing.T) {
//
//	var err error
//	s := setup(t)
//
//	var params = &sgo.CustomerParams{}
//
//	if _, err = s.CustomerGet("cus_HiykXILsFQrdhe111", params); err == nil {
//		t.Fatal(err)
//	}
//}
//
////Testing for failure case
//func TestStripe_SubscriptionUpdate(t *testing.T) {
//
//	var err error
//	s := setup(t)
//
//	var params = &sgo.CustomerParams{}
//
//	if _, err = s.CustomerUpdate("cus_HiykXILsFQrdhe", params); err == nil {
//		t.Fatal(err)
//	}
//}
//
////Testing for failure case
//func TestStripe_SubscriptionDel(t *testing.T) {
//
//	var err error
//	s := setup(t)
//
//	var params = &sgo.CustomerParams{}
//
//	if _, err = s.CustomerDel("cus_HiykXILsFQrdhe", params); err == nil {
//		t.Fatal(err)
//	}
//}
//
//func TestStripe_SubscriptionList(t *testing.T) {
//	s := setup(t)
//
//	var params = &sgo.CustomerListParams{}
//
//	i := s.CustomerList(params)
//
//	i.Next()
//
//	fmt.Println(i.Customer())
//}
