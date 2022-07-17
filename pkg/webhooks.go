// Package kopokopo Webhooks are a means of getting notified of events in the Kopo Kopo application.
// To receive webhooks, you need to create a webhook subscription.
package kopokopo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var webhookSubsEndpoint = "api/v1/webhook_subscriptions"

func (sdk kSDK) CreateWebhook(token string, webookReq CreateWebhookReq) (string, error) {
	if err := webookReq.Validate(); err != nil {
		return "", err
	}
	endpoint := fmt.Sprintf("%s/%s", sdk.baseURL, webhookSubsEndpoint)
	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return "", err
	}
	url, err := sdk.getHeaderParams(req, token)
	if err != nil {
		return "", err
	}
	id := strings.TrimPrefix(url, fmt.Sprintf("%s/%s/", sdk.baseURL, webhookSubsEndpoint))
	return id, nil
}

func (sdk kSDK) ValidateWebhook(webhookURL string) (BuyGoodsTrans, error) {
	if webhookURL == "" {
		return BuyGoodsTrans{}, errors.New("empty webhook url")
	}
	req, err := http.NewRequest(http.MethodPost, webhookURL, nil)
	if err != nil {
		return BuyGoodsTrans{}, err
	}
	req.Header.Add("X-KopoKopo-Signature", sdk.credentials.APIKey)
	resp, err := sdk.getBodyParams(req, "")
	if err != nil {
		return BuyGoodsTrans{}, err
	}
	var bgt BuyGoodsTrans
	if err := json.Unmarshal(resp, &bgt); err != nil {
		return BuyGoodsTrans{}, err
	}
	return bgt, nil
}

func (sdk kSDK) C2BSubscription(webhookURL string) (BuyGoodsTrans, error) {
	if webhookURL == "" {
		return BuyGoodsTrans{}, errors.New("empty webhook url")
	}
	req, err := http.NewRequest(http.MethodPost, webhookURL, nil)
	if err != nil {
		return BuyGoodsTrans{}, err
	}
	req.Header.Add("X-KopoKopo-Signature", sdk.credentials.APIKey)
	resp, err := sdk.getBodyParams(req, "")
	if err != nil {
		return BuyGoodsTrans{}, err
	}
	var bgt BuyGoodsTrans
	if err := json.Unmarshal(resp, &bgt); err != nil {
		return BuyGoodsTrans{}, err
	}
	return bgt, nil
}

func (sdk kSDK) B2BSubscription(webhookURL string) (BuyGoodsTrans, error) {
	if webhookURL == "" {
		return BuyGoodsTrans{}, errors.New("empty webhook url")
	}
	req, err := http.NewRequest(http.MethodPost, webhookURL, nil)
	if err != nil {
		return BuyGoodsTrans{}, err
	}
	req.Header.Add("X-KopoKopo-Signature", sdk.credentials.APIKey)
	resp, err := sdk.getBodyParams(req, "")
	if err != nil {
		return BuyGoodsTrans{}, err
	}
	var bgt BuyGoodsTrans
	if err := json.Unmarshal(resp, &bgt); err != nil {
		return BuyGoodsTrans{}, err
	}
	return bgt, nil
}

func (sdk kSDK) M2MSubscription(webhookURL string) (BuyGoodsTrans, error) {
	if webhookURL == "" {
		return BuyGoodsTrans{}, errors.New("empty webhook url")
	}
	req, err := http.NewRequest(http.MethodPost, webhookURL, nil)
	if err != nil {
		return BuyGoodsTrans{}, err
	}
	req.Header.Add("X-KopoKopo-Signature", sdk.credentials.APIKey)
	resp, err := sdk.getBodyParams(req, "")
	if err != nil {
		return BuyGoodsTrans{}, err
	}
	var bgt BuyGoodsTrans
	if err := json.Unmarshal(resp, &bgt); err != nil {
		return BuyGoodsTrans{}, err
	}
	return bgt, nil
}

func (sdk kSDK) C2BReversalSubscription(webhookURL string) (BuyGoodsTrans, error) {
	if webhookURL == "" {
		return BuyGoodsTrans{}, errors.New("empty webhook url")
	}
	req, err := http.NewRequest(http.MethodPost, webhookURL, nil)
	if err != nil {
		return BuyGoodsTrans{}, err
	}
	req.Header.Add("X-KopoKopo-Signature", sdk.credentials.APIKey)
	resp, err := sdk.getBodyParams(req, "")
	if err != nil {
		return BuyGoodsTrans{}, err
	}
	var bgt BuyGoodsTrans
	if err := json.Unmarshal(resp, &bgt); err != nil {
		return BuyGoodsTrans{}, err
	}
	return bgt, nil
}

func (sdk kSDK) SettlementSub(webhookURL string) (BuyGoodsTrans, error) {
	if webhookURL == "" {
		return BuyGoodsTrans{}, errors.New("empty webhook url")
	}
	req, err := http.NewRequest(http.MethodPost, webhookURL, nil)
	if err != nil {
		return BuyGoodsTrans{}, err
	}
	req.Header.Add("X-KopoKopo-Signature", sdk.credentials.APIKey)
	resp, err := sdk.getBodyParams(req, "")
	if err != nil {
		return BuyGoodsTrans{}, err
	}
	var bgt BuyGoodsTrans
	if err := json.Unmarshal(resp, &bgt); err != nil {
		return BuyGoodsTrans{}, err
	}
	return bgt, nil
}

func (sdk kSDK) CusomerCreationSub(webhookURL string) (CustomerReq, error) {
	if webhookURL == "" {
		return CustomerReq{}, errors.New("empty webhook url")
	}
	req, err := http.NewRequest(http.MethodPost, webhookURL, nil)
	if err != nil {
		return CustomerReq{}, err
	}
	req.Header.Add("X-KopoKopo-Signature", sdk.credentials.APIKey)
	resp, err := sdk.getBodyParams(req, "")
	if err != nil {
		return CustomerReq{}, err
	}
	var cr CustomerReq
	if err := json.Unmarshal(resp, &cr); err != nil {
		return CustomerReq{}, err
	}
	return cr, nil
}
