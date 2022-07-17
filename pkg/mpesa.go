package kopokopo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (sdk kSDK) ReceiveMpesaPayment(token string, receiveMpesaReq ReceiveMpesaReq) (string, error) {
	if err := receiveMpesaReq.Validate(); err != nil {
		return "", err
	}
	data, err := json.Marshal(receiveMpesaReq)
	if err != nil {
		return "", err
	}
	endpoint := fmt.Sprintf("%s/api/v1/incoming_payments", sdk.baseURL)
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	url, err := sdk.getHeaderParams(req, token)
	if err != nil {
		return "", err
	}
	id := strings.TrimPrefix(url, fmt.Sprintf("%s/api/v1/incoming_payments/", sdk.baseURL))
	return id, nil
}

// func (sdk kSDK) ProcessIncommingMpesaPayment(grantType string) (string, error) {
// 	panic("Not implemented")
// }

func (sdk kSDK) QueryIncommingMpesaPayment(token, id string) (IncomingPaymentEvent, error) {
	if id == "" {
		return IncomingPaymentEvent{}, errors.New("empty id")
	}

	endpoint := fmt.Sprintf("%s/api/v1/incoming_payments/", sdk.baseURL)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
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
