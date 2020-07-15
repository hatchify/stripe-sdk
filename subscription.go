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
