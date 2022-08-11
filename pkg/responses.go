package kopokopo

// type ErrorResp struct {
// 	Type    string `json:"error,omitempty"`
// 	Message string `json:"error_description,omitempty"`
// 	State   string `json:"state,omitempty"`
// }

type tokenResp struct {
	AccessToken string `json:"access_token,omitempty"` // Access Token
	Type        string `json:"token_type,omitempty"`   // Type of token
	Expiry      uint64 `json:"expires_in,omitempty"`   // Expiry duration of token
	Creation    uint64 `json:"created_at,omitempty"`   // Creation timestamp of token
}

type tokenIntrospectionResp struct {
	Active   bool   `json:"active,omitempty"`     // If the token is active or not
	Scope    string `json:"scope,omitempty"`      // The application scope of the token
	ClientID string `json:"client_id,omitempty"`  // The application id associated with the token
	Type     string `json:"token_type,omitempty"` // Type of the token
	Expiry   uint64 `json:"exp,omitempty"`        // Expiry timestamp of the token
	Creation uint64 `json:"iat,omitempty"`        // Creation timestamp of the token
}
type application struct {
	UID string `json:"uid,omitempty"` // The application id associated with the token
}
type tokenInfo struct {
	OwnerID     string      `json:"resource_owner_id,omitempty"` // The owner id associated with the token
	Scope       []string    `json:"scope,omitempty"`             // The application scope of the token
	Expiry      uint64      `json:"expires_in,omitempty"`        // Expiry duration of token
	Application application `json:"application,omitempty"`
	Creation    uint64      `json:"created_at,omitempty"` // Creation timestamp of the token
}
