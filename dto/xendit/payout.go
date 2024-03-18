package xendit

import (
	"math/big"
	"time"
)

type XenditPayoutRequestReq struct {
	ReferenceID       string `json:"reference_id"`
	ChannelCode       string `json:"channel_code"`
	ChannelProperties struct {
		AccountHolderName string `json:"account_holder_name"`
		AccountNumber     string `json:"account_number"`
	} `json:"channel_properties"`
	Amount              big.Rat `json:"amount"`
	Description         *string `json:"description"`
	Currency            string  `json:"currency"`
	ReceiptNotification *struct {
		EmailTo *[]string `json:"email_to"`
		EmailCC *[]string `json:"email_cc"`
	} `json:"receipt_notification"`
}

type XenditPayoutRequestReqHeader struct {
	IdempotencyKey string `json:"Idempotency-key"`
}

type XenditPayoutObject struct {
	ID                   string    `json:"id"`
	Amount               big.Rat   `json:"amount"`
	ChannelCode          string    `json:"channel_code"`
	Currency             string    `json:"currency"`
	Description          *string   `json:"description"`
	ReferenceID          string    `json:"reference_id"`
	Status               string    `json:"status"`
	Created              time.Time `json:"created"`
	Updated              time.Time `json:"updated"`
	EstimatedArrivalTime time.Time `json:"estimated_arrival_time"`
	BusinessID           string    `json:"business_id"`
	ChannelProperties    struct {
		AccountNumber     string `json:"account_number"`
		AccountHolderName string `json:"account_holder_name"`
	} `json:"channel_properties"`
	ReceiptNotification *struct {
		EmailTo  *[]string `json:"email_to"`
		EmailCC  *[]string `json:"email_cc"`
		EmailBCC *[]string `json:"email_bcc"`
	} `json:"receipt_notification"`
}

type XenditPayoutRequestRes struct {
	XenditPayoutObject
}

type XenditPayoutCallbackData struct {
	ID                  string    `json:"id"`
	Amount              big.Rat   `json:"amount"`
	Status              string    `json:"status"`
	Created             time.Time `json:"created"`
	Updated             time.Time `json:"updated"`
	Currency            string    `json:"currency"`
	Description         *string   `json:"description"`
	ChannelCode         string    `json:"channel_code"`
	ReferenceID         string    `json:"reference_id"`
	AccountNumber       string    `json:"account_number"`
	IdempotencyKey      string    `json:"idempotency_key"`
	ChannelCategory     string    `json:"channel_category"`
	AccountHolderName   string    `json:"account_holder_name"`
	ConnectorReference  string    `json:"connector_reference"`
	ReceiptNotification *struct {
		EmailCC *[]string `json:"email_cc"`
	} `json:"receipt_notification"`
	EstimatedArrivalTime time.Time `json:"estimated_arrival_time"`
}

type XenditPayoutCallbackReq struct {
	Event      string                   `json:"event"`
	BusinessId string                   `json:"business_id"`
	Created    time.Time                `json:"created"`
	Data       XenditPayoutCallbackData `json:"data"`
}

type XenditPayoutCallbackHeader struct {
	XCallbackToken string `json:"x-callback-token"`
	WebhookId      string `json:"webhook-id"`
}
