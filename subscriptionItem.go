package stripe

// SubscriptionItem is the resource representing a Stripe subscription item.
// For more details see https://stripe.com/docs/api#subscription_items.
type SubscriptionItem struct {
	ID                string                `json:"id"`
	Object            string                `json:"object"`
	BillingThresholds interface{}           `json:"billing_thresholds"`
	Created           int                   `json:"created"`
	Plan              SubscriptionItemPlan  `json:"plan"`
	Price             SubscriptionItemPrice `json:"price"`
	Subscription      string                `json:"subscription"`
	TaxRates          []interface{}         `json:"tax_rates"`
}

type SubscriptionItemPlan struct {
	ID              string      `json:"id"`
	Object          string      `json:"object"`
	Active          bool        `json:"active"`
	AggregateUsage  interface{} `json:"aggregate_usage"`
	Amount          int         `json:"amount"`
	AmountDecimal   string      `json:"amount_decimal"`
	BillingScheme   string      `json:"billing_scheme"`
	Created         int         `json:"created"`
	Currency        string      `json:"currency"`
	Interval        string      `json:"interval"`
	IntervalCount   int         `json:"interval_count"`
	Livemode        bool        `json:"livemode"`
	Nickname        interface{} `json:"nickname"`
	Product         string      `json:"product"`
	Tiers           interface{} `json:"tiers"`
	TiersMode       interface{} `json:"tiers_mode"`
	TransformUsage  interface{} `json:"transform_usage"`
	TrialPeriodDays interface{} `json:"trial_period_days"`
	UsageType       string      `json:"usage_type"`
}

type SubscriptionItemRecurring struct {
	AggregateUsage interface{} `json:"aggregate_usage"`
	Interval       string      `json:"interval"`
	IntervalCount  int         `json:"interval_count"`
	UsageType      string      `json:"usage_type"`
}

type SubscriptionItemPrice struct {
	ID                string                    `json:"id"`
	Object            string                    `json:"object"`
	Active            bool                      `json:"active"`
	BillingScheme     string                    `json:"billing_scheme"`
	Created           int                       `json:"created"`
	Currency          string                    `json:"currency"`
	Livemode          bool                      `json:"livemode"`
	LookupKey         interface{}               `json:"lookup_key"`
	Nickname          interface{}               `json:"nickname"`
	Product           string                    `json:"product"`
	Recurring         SubscriptionItemRecurring `json:"recurring"`
	TiersMode         interface{}               `json:"tiers_mode"`
	TransformQuantity interface{}               `json:"transform_quantity"`
	Type              string                    `json:"type"`
	UnitAmount        int                       `json:"unit_amount"`
	UnitAmountDecimal string                    `json:"unit_amount_decimal"`
}
