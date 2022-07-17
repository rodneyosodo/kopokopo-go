package kopokopo

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// ContentType represents all possible content types.
type ContentType string

var _ SDK = (*kSDK)(nil)

// SDK contains Kopokopo API.
type SDK interface {

	// The client credentials flow is the simplest OAuth 2 grant,
	// with a server-to-server exchange of your application’s client_id, client_secret
	// for an OAuth application access token. In order to execute this flow,
	// you will need to make an HTTP request from your application server,
	// to the Kopo Kopo authorization server.
	GetToken() (string, error)

	// The request is used to revoke a particular token at a time.
	RevokeToken(token string) error

	// It can be used to check the validity of your access tokens, and find out other
	// information such as which user and which scopes are associated with the token.
	// The client secret will not be displayed as that is to remain confidential with the application owner.
	TokenIntrospection(token string) (tokenIntrospectionResp, error)

	// Shows details about the token used for authentication.
	TokenInformation(token string) (tokenInfo, error)

	// Create a webhook subscription
	CreateWebhook(token string, webookReq CreateWebhookReq) (string, error)

	// Before processing webhook events, make sure that they originated from Kopo Kopo.
	// Each request is signed with the api_key you got when creating an oauth application on the platform.
	ValidateWebhook(webhookURL string) (BuyGoodsTrans, error)

	// Notifies your application when a Buygoods Transaction has been received.
	C2BSubscription(webhookURL string) (BuyGoodsTrans, error)

	// Notifies your application when a B2b (External Till to Till transaction) has been received.
	// These are payments recieved from other tills and not subscribers.
	B2BSubscription(webhookURL string) (BuyGoodsTrans, error)

	// Notifies your application when another Kopo Kopo merchant transfers funds
	// to your Kopo Kopo merchant account (Merchant to Merchant)
	M2MSubscription(webhookURL string) (BuyGoodsTrans, error)

	// Notifies your application when a Buygoods Transaction has been reversed
	C2BReversalSubscription(webhookURL string) (BuyGoodsTrans, error)

	// Settlement Transfer Completed
	SettlementSub(webhookURL string) (BuyGoodsTrans, error)

	// Customer Created
	CusomerCreationSub(webhookURL string) (CustomerReq, error)

	// Receive payments from M-PESA users via STK Push.
	ReceiveMpesaPayment(token string, receiveMpesaReq ReceiveMpesaReq) (string, error)

	// With an Incoming Payment location url, you can query what the status of the Incoming Payment is.
	// If a corresponding Incoming Payment Result exists, it will be bundled in the payload of the result.
	QueryIncommingMpesaPayment(token, id string) (IncomingPaymentEvent, error)
}

// Credentials contains the credentials
type Credentials struct {
	AppID  string `json:"app_id"` // Application key.
	Secret string `json:"secret"` // Application secret. Only revealed to the user when creating an application or during regeneration of client credentials.
	APIKey string `json:"api_key"`
}

type kSDK struct {
	baseURL     string
	credentials Credentials
	client      *http.Client
}

// Config contains sdk configuration parameters.
type Config struct {
	BaseURL         string
	Credentials     Credentials
	MaxIdleConns    int
	IdleConnTimeout time.Duration
}

// NewSDK returns new mainflux SDK instance.
func NewSDK(conf Config) SDK {
	return &kSDK{
		baseURL:     conf.BaseURL,
		credentials: conf.Credentials,
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:    conf.MaxIdleConns,
				IdleConnTimeout: conf.IdleConnTimeout,
			},
		},
	}
}

func (sdk kSDK) makeRequest(req *http.Request, token string) (*http.Response, error) {
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (sdk kSDK) getBodyParams(req *http.Request, token string) ([]byte, error) {
	resp, err := sdk.makeRequest(req, token)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return body, nil
}

func (sdk kSDK) getHeaderParams(req *http.Request, token string) (string, error) {
	resp, err := sdk.makeRequest(req, token)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusCreated {
		return "", errors.New("failed to created")
	}
	id := resp.Header.Get("Location")
	return id, nil
}
