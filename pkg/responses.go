package kopokopo

// type ErrorResp struct {
// 	Code    int    `json:"error_code"`
// 	Message string `json:"error_message"`
// }

type tokenResp struct {
	AccessToken string `json:"access_token,omitempty"`
	Type        string `json:"token_type,omitempty"`
	Expiry      uint64 `json:"expires_in,omitempty"`
	Creation    uint64 `json:"created_at,omitempty"`
}

type tokenIntrospectionResp struct {
	Active   bool   `json:"active,omitempty"`
	Scope    string `json:"scope,omitempty"`
	ClientID string `json:"client_id,omitempty"`
	Type     string `json:"token_type,omitempty"`
	Expiry   uint64 `json:"exp,omitempty"`
	Creation uint64 `json:"iat,omitempty"`
}
type application struct {
	UID string `json:"uid,omitempty"`
}
type tokenInfo struct {
	OwnerID     string      `json:"resource_owner_id,omitempty"`
	Scope       []string    `json:"scope,omitempty"`
	Expiry      uint64      `json:"expires_in,omitempty"`
	Application application `json:"application,omitempty"`
	Creation    uint64      `json:"created_at,omitempty"`
}
