package kopokopo

import (
	"errors"
	"strings"
)

// CreateWebhookReq struct
type CreateWebhookReq struct {
	EventType string `json:"event_type,omitempty"` //The type of event you are subscribing to
	URL       string `json:"url,omitempty"`        // The http end point to send the webhook.
	Scope     string `json:"scope,omitempty"`      // The scope of the webhook subscription.
	ScopeRef  string `json:"scope_reference,omitempty"`
}

// Validate returns nil if the struct is valid
func (cwr CreateWebhookReq) Validate() error {
	if cwr.EventType != "buygoods_transaction_received" &&
		cwr.EventType != "buygoods_transaction_reversed" &&
		cwr.EventType != "b2b_transaction_received" &&
		cwr.EventType != "m2m_transaction_received" &&
		cwr.EventType != "settlement_transfer_completed" &&
		cwr.EventType != "customer_created" {
		return errors.New("invalid event type")
	}
	if !strings.HasPrefix(cwr.URL, "https") {
		return errors.New("URL is not secured with TLS")
	}
	if cwr.Scope != "company" && cwr.Scope != "till" {
		return errors.New("invalid scope")
	}
	if cwr.Scope == "till" && cwr.ScopeRef == "" {
		return errors.New("scope reference is required")
	}
	return nil
}

// Destination struct
type Destination struct {
	Type     string              `json:"type,omitempty"`
	Resource DestinationResource `json:"resource,omitempty"`
}

// DestinationResource struct
type DestinationResource struct {
	Reference           string `json:"reference,omitempty"`         // The destination reference
	AccountName         string `json:"account_name,omitempty"`      // The name as indicated on the bank account
	AccountNumber       string `json:"account_number,omitempty"`    // The bank account number
	BankBranchReference string `json:"bank_branch_ref,omitempty"`   // An identifier identifying the destination bank branch
	SettlementMethod    string `json:"settlement_method,omitempty"` // EFT or RTS
	FirstName           string `json:"first_name,omitempty"`        // String	First name of the recipient
	LastName            string `json:"last_name,omitempty"`         // Last name of recipient
	Email               string `json:"email,omitempty"`             // Email of recipient
	PhoneNumber         string `json:"phone_number,omitempty"`      // Phone number
	Network             string `json:"network,omitempty"`           // The mobile network to which the phone number belongs
}

// Disbursements struct
type Disbursements struct {
	Status                 string `json:"status,omitempty"`                  // The status of the disbursement
	Amount                 string `json:"amount,omitempty"`                  // The amount of the disbursement
	OriginationTime        string `json:"origination_time,omitempty"`        // The Timestamp of when the transaction took place
	TransactionalReference string `json:"transactional_reference,omitempty"` // The reference from the transaction. i.e mpesa reference It is null for eft transactions
}

// Resource struct
type Resource struct {
	ID                     string          `json:"id,omitempty"`        // The api reference of the transaction
	Amount                 float64         `json:"amount,omitempty"`    // The amount of the transaction
	Status                 string          `json:"status,omitempty"`    // The status of the transaction
	System                 string          `json:"system,omitempty"`    // The mobile money system
	Currency               string          `json:"currency,omitempty"`  // Currency
	Reference              string          `json:"reference,omitempty"` // The mpesa reference
	TransactionalReference string          `json:"transactional_reference,omitempty"`
	TillNumber             string          `json:"till_number,omitempty"`  // The till number to which the payment was made
	SendingTill            string          `json:"sending_till,omitempty"` // The till number of the sender
	AccountName            string          `json:"account_name,omitempty"`
	AccountNumber          string          `json:"account_number,omitempty"`
	BankBranchReference    string          `json:"bank_branch_ref,omitempty"`
	SettlementMethod       string          `json:"settlement_method,omitempty"`
	SendingMerchant        string          `json:"sending_merchant,omitempty"`    // Name of merchant
	SenderPhoneNumber      string          `json:"sender_phone_number,omitempty"` // The phone number that sent the payment
	OriginationTime        string          `json:"origination_time,omitempty"`    // The transaction timestamp
	SenderLastName         string          `json:"sender_last_name,omitempty"`    // Last name of payer
	SenderFirstName        string          `json:"sender_first_name,omitempty"`   // First name of payer
	SenderMiddleName       string          `json:"sender_middle_name,omitempty"`  // Middle name of payer
	Destination            Destination     `json:"destination,omitempty"`         // The destination of the settlement transfer
	Disbursements          []Disbursements `json:"disbursements,omitempty"`       // These are the disbursements in that particular transfer batch
}

// Event struct
type Event struct {
	Type     string   `json:"type,omitempty"`     // The type of transaction
	Resource Resource `json:"resource,omitempty"` // The resource corresponding to the event.
}

// Links struct
type Links struct {
	Self     string `json:"self,omitempty"`
	Resource string `json:"resource,omitempty"`
	Callback string `json:"callback_url,omitempty"`
}

// BuyGoodsTrans struct
type BuyGoodsTrans struct {
	Topic     string `json:"topic,omitempty"`      // The ID of the Webhook Event
	ID        string `json:"id,omitempty"`         // The topic of the webhook.
	CreatedAt string `json:"created_at,omitempty"` // The timestamp of when the webhook event was created.
	Event     Event  `json:"event,omitempty"`
	Links     Links  `json:"_links,omitempty"` // A JSON object containing links to the Webhook Event and the corresponding Buygoods Transaction resource
}

// Validate returns nil if the struct is valid
func (bgt BuyGoodsTrans) Validate() error {
	return nil
}

// CustomerResource struct
type CustomerResource struct {
	LastName    string `json:"last_name,omitempty"`    // Last name of payer
	FirstName   string `json:"first_name,omitempty"`   // First name of payer
	MiddleName  string `json:"middle_name,omitempty"`  // Middle name of payer
	PhoneNumber string `json:"phone_number,omitempty"` // The phone number that sent the payment
}

// CustomerEvent struct
type CustomerEvent struct {
	Type     string           `json:"type,omitempty"`     // The type of record (Mobile Money User)
	Resource CustomerResource `json:"resource,omitempty"` // The resource corresponding to the event.
}

// CustomerReq struct
type CustomerReq struct {
	Topic     string        `json:"topic,omitempty"`      // The ID of the Webhook Event
	ID        string        `json:"id,omitempty"`         // The topic of the webhook.
	CreatedAt string        `json:"created_at,omitempty"` // The timestamp of when the webhook event was created.
	Event     CustomerEvent `json:"event,omitempty"`
	Links     Links         `json:"_links,omitempty"` // A JSON object containing links to the Webhook Event and the corresponding Buygoods Transaction resource
}

// Subscriber struct
type Subscriber struct {
	LastName    string `json:"last_name,omitempty"`    // Last name of the subscriber
	FirstName   string `json:"first_name,omitempty"`   // First name of the subscriber
	MiddleName  string `json:"middle_name,omitempty"`  // Middle name of the subscriber
	PhoneNumber string `json:"phone_number,omitempty"` // The phone number of the subscriber from which the payment will be made
	Email       string `json:"email,omitempty"`        // E-mail address of the subscriber - optional
}

// Amount struct
type Amount struct {
	Value    string `json:"value,omitempty"`    // The amount of the transaction
	Currency string `json:"currency,omitempty"` // Currency
}

// ReceiveMpesaReq struct
type ReceiveMpesaReq struct {
	PaymentChannel string                 `json:"payment_channel,omitempty"` // The payment channel to be used eg. M-PESA
	TillNumber     string                 `json:"till_number,omitempty"`     // The online payments till number from the Kopo Kopo dashboard to which the payment will be made
	Subscriber     Subscriber             `json:"subscriber,omitempty"`      // A Subscriber JSON object see below
	Amount         Amount                 `json:"amount,omitempty"`          // An Amount JSON object containing currency and amount
	Metadata       map[string]interface{} `json:"metadata,omitempty"`        // An optional JSON object containing a maximum of 5 key value pairs
	Links          Links                  `json:"_links,omitempty"`          // A JOSN object containing the call back URL where the result of the Incoming Payment will be posted
}

// Validate returns nil if the struct is valid
func (rmr ReceiveMpesaReq) Validate() error {
	return nil
}

// IncomingPaymentEvent struct
type IncomingPaymentEvent struct {
	Type       string    `json:"type,omitempty"` // The ID of the Webhook Event
	ID         string    `json:"id,omitempty"`   // The topic of the webhook.
	Attributes Attribute `json:"attributes,omitempty"`
}

// Attribute struct
type Attribute struct {
	InitiationTime string                 `json:"initiation_time,omitempty"` // The timestamp of when the webhook event was created.
	Status         string                 `json:"status,omitempty"`          // A status string denoting the status of the Incoming Payment
	Resource       Resource               `json:"resource,omitempty"`        // A JSON Object encapsulating the event of the request
	Metadata       map[string]interface{} `json:"metadata,omitempty"`        // An optional JSON object containing a maximum of 5 key value pairs
	Links          Links                  `json:"_links,omitempty"`          // A JSON object containing links to the Webhook Event and the corresponding Buygoods Transaction resource

}

// type IncomingPaymentEvent struct {
// 	Type     string   `json:"type,omitempty"`     // The type of record (Mobile Money User)
// 	Resource Resource `json:"resource,omitempty"` // The resource corresponding to the event.
// 	Errors   string   `json:"errors,omitempty"`   // A string containing information on the error than occured
// }

// ProcessIncommingPaymentReq struct
type ProcessIncommingPaymentReq struct {
	Topic          string                 `json:"topic,omitempty"`           // The topic of the request.
	ID             string                 `json:"id,omitempty"`              // The ID of the Incoming Payment
	InitiationTime string                 `json:"initiation_time,omitempty"` // The timestamp of when the webhook event was created.
	Status         string                 `json:"status,omitempty"`          // A status string denoting the status of the Incoming Payment
	Event          IncomingPaymentEvent   `json:"event,omitempty"`           // A JSON Object encapsulating the event of the request
	Metadata       map[string]interface{} `json:"metadata,omitempty"`        // An optional JSON object containing a maximum of 5 key value pairs
	Links          Links                  `json:"_links,omitempty"`          // A JSON object containing links to the Webhook Event and the corresponding Buygoods Transaction resource
}

// type PaymentRecipient struct {
// 	LastName             string `json:"last_name,omitempty"`              // Last name of the recipient
// 	FirstName            string `json:"first_name,omitempty"`             // First name of the recipient
// 	MiddleName           string `json:"middle_name,omitempty"`            // Middle name of the recipient
// 	PhoneNumber          string `json:"phone_number,omitempty"`           // The phone number of the recipient from which the payment will be made
// 	Email                string `json:"email,omitempty"`                  // E-mail address of the recipient - optional
// 	AccountName          string `json:"account_name,omitempty"`           //The name as indicated on the bank account name
// 	BankBranchReference  string `json:"bank_branch_ref,omitempty"`        // An identifier identifying the destination bank branch.
// 	AccountNumber        string `json:"account_number,omitempty"`         // The bank account number
// 	SettlementMethod     string `json:"settlement_method,omitempty"`      // RTS
// 	TillName             string `json:"till_name,omitempty"`              // The name as indicated on the till
// 	TillNumber           string `json:"till_number,omitempty"`            // The till number
// 	PayBillName          string `json:"paybill_name,omitempty"`           // The name referring to the paybill
// 	PayBillNumber        string `json:"paybill_number,omitempty"`         // The paybill business number
// 	PayBillAccountNumber string `json:"paybill_account_number,omitempty"` // The paybill account number
// }

// type AddPAYRecipient struct {
// 	Type             string           `json:"type,omitempty"`              // The type of the recipient eg. mobile wallet or bank account
// 	PaymentRecipient PaymentRecipient `json:"payment_recipient,omitempty"` // 	A JSON object containing details of the recipeint
// }

// type CreatePaymentReq struct {
// 	DestinationType      string                 `json:"destination_type,omitempty"`      // Pay recipient type (bank_account, mobile_wallet, till or paybill
// 	DestinationReference string                 `json:"destination_reference,omitempty"` // Reference for the destination.
// 	Amount               Amount                 `json:"amount,omitempty"`                // A JSON object containing the currency and the amount to be transferred
// 	Description          string                 `json:"description,omitempty"`           // A reason for the payment
// 	Category             string                 `json:"category,omitempty"`              // Categorize the transaction
// 	Tags                 string                 `json:"tags,omitempty"`                  // Define your own tag to label the transaction with
// 	Metadata             map[string]interface{} `json:"metadata,omitempty"`              // A JSON containing upto a maximum of 5 key-value pairs for your own use
// 	Links                Links                  `json:"_links,omitempty"`                // A JSON containing a call back URL where the results of the Payment will be posted. MUST be a secure HTTPS (TLS) endpoint
// }

// type MerchantBankAccountReq struct {
// }

// type MerchantMobileAccountReq struct {
// }

// type BlinkTransferReq struct {
// }

// type TargetedTransferReq struct {
// }

// type SendSMSReq struct {
// 	WebhookID string `json:"webhook_event_reference,omitempty"` // This is the id of the webhook payload you got
// 	Message   string `json:"message,omitempty"`                 // A string containing the message you want to send to the customer
// 	Links     Links  `json:"_links,omitempty"`                  // A JSON object containing the callback_url where the result of the Transaction Notification will be posted
// }
