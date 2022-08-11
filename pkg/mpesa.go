package kopokopo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var mpesaEndpoint = "api/v1/incoming_payments"

// ReceiveMpesaPayment Receive payments from M-PESA users via STK Push.
func (sdk kSDK) ReceiveMpesaPayment(token string, receiveMpesaReq ReceiveMpesaReq) (string, error) {
	if err := receiveMpesaReq.Validate(); err != nil {
		return "", err
	}
	data, err := json.Marshal(receiveMpesaReq)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/%s", sdk.baseURL, mpesaEndpoint)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	locationURL, err := sdk.getHeaderParams(req, token)
	if err != nil {
		return "", err
	}
	id := strings.TrimPrefix(locationURL, fmt.Sprintf("%s/%s/", sdk.baseURL, mpesaEndpoint))
	return id, nil
}

// ProcessIncommingMpesaPayment Process Incoming Payment Result
func (sdk kSDK) ProcessIncommingMpesaPayment(grantType string) (string, error) {
	panic("Not implemented")
}

// QueryIncommingMpesaPayment Query Incoming Payment Status
func (sdk kSDK) QueryIncommingMpesaPayment(token, id string) (IncomingPaymentEvent, error) {
	if id == "" {
		return IncomingPaymentEvent{}, ErrEmptyID
	}

	url := fmt.Sprintf("%s/%s/", sdk.baseURL, mpesaEndpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return IncomingPaymentEvent{}, err
	}
	resp, err := sdk.getBodyParams(req, token)
	if err != nil {
		return IncomingPaymentEvent{}, err
	}
	var ipe IncomingPaymentEvent
	if err := json.Unmarshal(resp, &ipe); err != nil {
		return ipe, err
	}
	return ipe, nil
}
