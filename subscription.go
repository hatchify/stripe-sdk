package stripe

// SubscriptionStatus is the list of allowed values for the subscription's status.
type SubscriptionStatus string

// List of values that SubscriptionStatus can take.
const (
	SubscriptionStatusActive            SubscriptionStatus = "active"
	SubscriptionStatusAll               SubscriptionStatus = "all"
	SubscriptionStatusCanceled          SubscriptionStatus = "canceled"
	SubscriptionStatusIncomplete        SubscriptionStatus = "incomplete"
	SubscriptionStatusIncompleteExpired SubscriptionStatus = "incomplete_expired"
	SubscriptionStatusPastDue           SubscriptionStatus = "past_due"
	SubscriptionStatusTrialing          SubscriptionStatus = "trialing"
	SubscriptionStatusUnpaid            SubscriptionStatus = "unpaid"
)

// SubscriptionCollectionMethod is the type of collection method for this subscription's invoices.
type SubscriptionCollectionMethod string

// List of values that SubscriptionCollectionMethod can take.
const (
	SubscriptionCollectionMethodChargeAutomatically SubscriptionCollectionMethod = "charge_automatically"
	SubscriptionCollectionMethodSendInvoice         SubscriptionCollectionMethod = "send_invoice"
)

// SubscriptionPauseCollectionBehavior is the payment collection behavior a paused subscription.
type SubscriptionPauseCollectionBehavior string

// List of values that SubscriptionPauseCollectionBehavior can take.
const (
	SubscriptionPauseCollectionBehaviorKeepAsDraft       SubscriptionPauseCollectionBehavior = "keep_as_draft"
	SubscriptionPauseCollectionBehaviorMarkUncollectible SubscriptionPauseCollectionBehavior = "mark_uncollectible"
	SubscriptionPauseCollectionBehaviorVoid              SubscriptionPauseCollectionBehavior = "void"
)

// SubscriptionPaymentBehavior lets you control the behavior of subscription creation in case of
// a failed payment.
type SubscriptionPaymentBehavior string

// List of values that SubscriptionPaymentBehavior can take.
const (
	SubscriptionPaymentBehaviorAllowIncomplete     SubscriptionPaymentBehavior = "allow_incomplete"
	SubscriptionPaymentBehaviorErrorIfIncomplete   SubscriptionPaymentBehavior = "error_if_incomplete"
	SubscriptionPaymentBehaviorPendingIfIncomplete SubscriptionPaymentBehavior = "pending_if_incomplete"
)

// SubscriptionProrationBehavior determines how to handle prorations when billing cycles change.
type SubscriptionProrationBehavior string

// List of values that SubscriptionProrationBehavior can take.
const (
	SubscriptionProrationBehaviorAlwaysInvoice    SubscriptionProrationBehavior = "always_invoice"
	SubscriptionProrationBehaviorCreateProrations SubscriptionProrationBehavior = "create_prorations"
	SubscriptionProrationBehaviorNone             SubscriptionProrationBehavior = "none"
)

// SubscriptionPendingInvoiceItemIntervalInterval controls the interval at which pending invoice
// items should be invoiced.
type SubscriptionPendingInvoiceItemIntervalInterval string

// List of values that SubscriptionPendingInvoiceItemIntervalInterval can take.
const (
	SubscriptionPendingInvoiceItemIntervalIntervalDay   SubscriptionPendingInvoiceItemIntervalInterval = "day"
	SubscriptionPendingInvoiceItemIntervalIntervalMonth SubscriptionPendingInvoiceItemIntervalInterval = "month"
	SubscriptionPendingInvoiceItemIntervalIntervalWeek  SubscriptionPendingInvoiceItemIntervalInterval = "week"
	SubscriptionPendingInvoiceItemIntervalIntervalYear  SubscriptionPendingInvoiceItemIntervalInterval = "year"
)

type Subscription struct {
	ID                            string            `json:"id"`
	Object                        string            `json:"object"`
	ApplicationFeePercent         interface{}       `json:"application_fee_percent"`
	BillingCycleAnchor            int               `json:"billing_cycle_anchor"`
	BillingThresholds             interface{}       `json:"billing_thresholds"`
	CancelAt                      interface{}       `json:"cancel_at"`
	CancelAtPeriodEnd             bool              `json:"cancel_at_period_end"`
	CanceledAt                    int               `json:"canceled_at"`
	CollectionMethod              string            `json:"collection_method"`
	Created                       int               `json:"created"`
	CurrentPeriodEnd              int               `json:"current_period_end"`
	CurrentPeriodStart            int               `json:"current_period_start"`
	Customer                      string            `json:"customer"`
	DaysUntilDue                  interface{}       `json:"days_until_due"`
	DefaultPaymentMethod          interface{}       `json:"default_payment_method"`
	DefaultSource                 interface{}       `json:"default_source"`
	DefaultTaxRates               []interface{}     `json:"default_tax_rates"`
	Discount                      interface{}       `json:"discount"`
	EndedAt                       int               `json:"ended_at"`
	Items                         SubscriptionItems `json:"items"`
	LatestInvoice                 interface{}       `json:"latest_invoice"`
	Livemode                      bool              `json:"livemode"`
	NextPendingInvoiceItemInvoice interface{}       `json:"next_pending_invoice_item_invoice"`
	PauseCollection               interface{}       `json:"pause_collection"`
	PendingInvoiceItemInterval    interface{}       `json:"pending_invoice_item_interval"`
	PendingSetupIntent            interface{}       `json:"pending_setup_intent"`
	PendingUpdate                 interface{}       `json:"pending_update"`
	Plan                          SubscriptionPlan  `json:"plan"`
	Quantity                      int               `json:"quantity"`
	Schedule                      interface{}       `json:"schedule"`
	StartDate                     int               `json:"start_date"`
	Status                        string            `json:"status"`
	TaxPercent                    interface{}       `json:"tax_percent"`
	TransferData                  interface{}       `json:"transfer_data"`
	TrialEnd                      int               `json:"trial_end"`
	TrialStart                    int               `json:"trial_start"`
}

type SubscriptionPlan struct {
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
	TrialPeriodDays int         `json:"trial_period_days"`
	UsageType       string      `json:"usage_type"`
}

type SubscriptionRecurring struct {
	AggregateUsage interface{} `json:"aggregate_usage"`
	Interval       string      `json:"interval"`
	IntervalCount  int         `json:"interval_count"`
	UsageType      string      `json:"usage_type"`
}

type SubscriptionPrice struct {
	ID                string                `json:"id"`
	Object            string                `json:"object"`
	Active            bool                  `json:"active"`
	BillingScheme     string                `json:"billing_scheme"`
	Created           int                   `json:"created"`
	Currency          string                `json:"currency"`
	Livemode          bool                  `json:"livemode"`
	LookupKey         interface{}           `json:"lookup_key"`
	Nickname          interface{}           `json:"nickname"`
	Product           string                `json:"product"`
	Recurring         SubscriptionRecurring `json:"recurring"`
	TiersMode         interface{}           `json:"tiers_mode"`
	TransformQuantity interface{}           `json:"transform_quantity"`
	Type              string                `json:"type"`
	UnitAmount        int                   `json:"unit_amount"`
	UnitAmountDecimal string                `json:"unit_amount_decimal"`
}

type SubscriptionData struct {
	ID                string            `json:"id"`
	Object            string            `json:"object"`
	BillingThresholds interface{}       `json:"billing_thresholds"`
	Created           int               `json:"created"`
	Plan              SubscriptionPlan  `json:"plan"`
	Price             SubscriptionPrice `json:"price"`
	Quantity          int               `json:"quantity"`
	Subscription      string            `json:"subscription"`
	TaxRates          []interface{}     `json:"tax_rates"`
}

type SubscriptionItems struct {
	Object  string             `json:"object"`
	Data    []SubscriptionData `json:"data"`
	HasMore bool               `json:"has_more"`
	URL     string             `json:"url"`
}

type SubscriptionsList struct {
	Object  string         `json:"object"`
	Data    []Subscription `json:"data"`
	HasMore bool           `json:"has_more"`
	URL     string         `json:"url"`
}
