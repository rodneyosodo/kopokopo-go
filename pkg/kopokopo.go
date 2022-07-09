package kopokopo

import (
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

func (sdk kSDK) makeRequest(req *http.Request, token string) ([]byte, error) {
	if token != "" {
		req.Header.Add("Authorization", "Basic "+token)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := sdk.client.Do(req)
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
