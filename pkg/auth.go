package kopokopo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var tokenEndpoint = "oauth"

// GetToken Request application authorization
func (sdk kSDK) GetToken() (tokenResp, error) {
	q := url.Values{}
	q.Add("client_id", sdk.credentials.AppID)
	q.Add("client_secret", sdk.credentials.Secret)
	q.Add("grant_type", "client_credentials")
	url := fmt.Sprintf("%s/%s/token?%s", sdk.baseURL, tokenEndpoint, q.Encode())

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return tokenResp{}, err
	}
	resp, err := sdk.getBodyParams(req, "")
	if err != nil {
		return tokenResp{}, err
	}

	var tr tokenResp
	if err := json.Unmarshal(resp, &tr); err != nil {
		return tokenResp{}, err
	}
	return tr, nil
}

// RevokeToken Revoke application's access token
func (sdk kSDK) RevokeToken(token string) error {
	if token == "" {
		return ErrEmptyToken
	}
	q := url.Values{}
	q.Add("client_id", sdk.credentials.AppID)
	q.Add("client_secret", sdk.credentials.Secret)
	q.Add("token", token)
	url := fmt.Sprintf("%s/%s/revoke?%s", sdk.baseURL, tokenEndpoint, q.Encode())

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}
	_, err = sdk.getBodyParams(req, "")
	if err != nil {
		return err
	}
	return nil
}

// TokenIntrospection Request token introspection
func (sdk kSDK) TokenIntrospection(token string) (tokenIntrospectionResp, error) {
	if token == "" {
		return tokenIntrospectionResp{}, ErrEmptyToken
	}
	q := url.Values{}
	q.Add("client_id", sdk.credentials.AppID)
	q.Add("client_secret", sdk.credentials.Secret)
	q.Add("token", token)
	url := fmt.Sprintf("%s/%s/introspect?%s", sdk.baseURL, tokenEndpoint, q.Encode())

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return tokenIntrospectionResp{}, err
	}
	resp, err := sdk.getBodyParams(req, "")
	if err != nil {
		return tokenIntrospectionResp{}, err
	}
	var tir tokenIntrospectionResp
	if err := json.Unmarshal(resp, &tir); err != nil {
		return tir, err
	}
	return tir, nil
}

// TokenInformation Request token information
func (sdk kSDK) TokenInformation(token string) (tokenInfo, error) {
	if token == "" {
		return tokenInfo{}, ErrEmptyToken
	}

	endpoint := fmt.Sprintf("%s/%s/token/info", sdk.baseURL, tokenEndpoint)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return tokenInfo{}, err
	}
	resp, err := sdk.getBodyParams(req, token)
	if err != nil {
		return tokenInfo{}, err
	}
	var ti tokenInfo
	if err := json.Unmarshal(resp, &ti); err != nil {
		return ti, err
	}
	return ti, nil
}
