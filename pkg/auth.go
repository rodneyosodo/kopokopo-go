package kopokopo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var tokenEndpoint = "oauth"

func (sdk kSDK) GetToken() (string, error) {
	q := url.Values{}
	q.Add("client_id", sdk.credentials.AppID)
	q.Add("client_secret", sdk.credentials.Secret)
	q.Add("grant_type", "client_credentials")
	endpoint := fmt.Sprintf("%s/%s/token?%s", sdk.baseURL, tokenEndpoint, q.Encode())

	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return "", err
	}
	resp, err := sdk.makeRequest(req, "")
	if err != nil {
		return "", err
	}

	var tr tokenResp
	if err := json.Unmarshal(resp, &tr); err != nil {
		return "", err
	}
	return tr.AccessToken, nil
}

func (sdk kSDK) RevokeToken(token string) error {
	if token == "" {
		return errors.New("empty token")
	}
	q := url.Values{}
	q.Add("client_id", sdk.credentials.AppID)
	q.Add("client_secret", sdk.credentials.Secret)
	q.Add("token", token)
	endpoint := fmt.Sprintf("%s/%s/revoke?%s", sdk.baseURL, tokenEndpoint, q.Encode())

	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return err
	}
	_, err = sdk.makeRequest(req, "")
	if err != nil {
		return err
	}
	return nil
}

func (sdk kSDK) TokenIntrospection(token string) (tokenIntrospectionResp, error) {
	if token == "" {
		return tokenIntrospectionResp{}, errors.New("empty token")
	}
	q := url.Values{}
	q.Add("client_id", sdk.credentials.AppID)
	q.Add("client_secret", sdk.credentials.Secret)
	q.Add("token", token)
	endpoint := fmt.Sprintf("%s/%s/token/info?%s", sdk.baseURL, tokenEndpoint, q.Encode())

	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return tokenIntrospectionResp{}, err
	}
	resp, err := sdk.makeRequest(req, "")
	if err != nil {
		return tokenIntrospectionResp{}, err
	}
	var tir tokenIntrospectionResp
	if err := json.Unmarshal(resp, &tir); err != nil {
		return tir, err
	}
	return tir, nil
}

func (sdk kSDK) TokenInformation(token string) (tokenInfo, error) {
	if token == "" {
		return tokenInfo{}, errors.New("empty token")
	}

	endpoint := fmt.Sprintf("%s/%s/token/info", sdk.baseURL, tokenEndpoint)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return tokenInfo{}, err
	}
	resp, err := sdk.makeRequest(req, token)
	if err != nil {
		return tokenInfo{}, err
	}
	var ti tokenInfo
	if err := json.Unmarshal(resp, &ti); err != nil {
		return ti, err
	}
	return ti, nil
}
