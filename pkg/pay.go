package kopokopo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var payEndpoint = "api/v1"

// AddPayRecipients Adding PAY recipients
func (sdk kSDK) AddPayRecipients(token string, recipient AddPAYRecipient) (string, error) {
	if err := recipient.Validate(); err != nil {
		return "", err
	}
	data, err := json.Marshal(recipient)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/%s/pay_recipients", sdk.baseURL, payEndpoint)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	locationURL, err := sdk.getHeaderParams(req, token)
	if err != nil {
		return "", err
	}
	id := strings.TrimPrefix(locationURL, fmt.Sprintf("%s/%s/pay_recipients/", sdk.baseURL, payEndpoint))
	return id, nil
}

// CreatePayment Create a Payment
func (sdk kSDK) CreatePayment(token string, payment CreatePaymentReq) (string, error) {
	if err := payment.Validate(); err != nil {
		return "", err
	}
	data, err := json.Marshal(payment)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/%s/payments", sdk.baseURL, payEndpoint)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	locationURL, err := sdk.getHeaderParams(req, token)
	if err != nil {
		return "", err
	}
	id := strings.TrimPrefix(locationURL, fmt.Sprintf("%s/%s/payments/", sdk.baseURL, payEndpoint))
	return id, nil
}

// func (sdk kSDK) ProcessPayment(grantType string) (string, error) {
// 	panic("Not implemented")
// }

// QueryPayment Query Payment status
// func (sdk kSDK) QueryPayment(token, id string) (string, error) {
// 	if id == "" {
// 		return IncomingPaymentEvent{}, EmptyID
// 	}

// 	url := fmt.Sprintf("%s/%s/", sdk.baseURL, mpesaEndpoint)
// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		return IncomingPaymentEvent{}, err
// 	}
// 	resp, err := sdk.getBodyParams(req, token)
// 	if err != nil {
// 		return IncomingPaymentEvent{}, err
// 	}
// 	var ipe IncomingPaymentEvent
// 	if err := json.Unmarshal(resp, &ipe); err != nil {
// 		return ipe, err
// 	}
// 	return ipe, nil
// }
