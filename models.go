package main

// RequestBody represents the incoming request structure
type RequestBody struct {
	FlowID            string             `json:"flowId,omitempty"`
	RequestID         string             `json:"requestId,omitempty"`
	ActionType        string             `json:"actionType"`
	Event             Event              `json:"event"`
	AllowedOperations []AllowedOperation `json:"allowedOperations"`
}

// Event contains the context data for the pre-issue access token event
type Event struct {
	Request      Request       `json:"request"`
	Tenant       Tenant        `json:"tenant"`
	User         *User         `json:"user,omitempty"`
	UserStore    *UserStore    `json:"userStore,omitempty"`
	AccessToken  AccessToken   `json:"accessToken"`
	RefreshToken *RefreshToken `json:"refreshToken,omitempty"`
}

// Request contains OAuth2 request details
type Request struct {
	GrantType         string          `json:"grantType"`
	ClientID          string          `json:"clientId"`
	Scopes            []string        `json:"scopes,omitempty"`
	AdditionalHeaders []RequestHeader `json:"additionalHeaders,omitempty"`
	AdditionalParams  []RequestParam  `json:"additionalParams,omitempty"`
}

// RequestHeader represents additional HTTP headers
type RequestHeader struct {
	Name  string   `json:"name"`
	Value []string `json:"value"`
}

// RequestParam represents additional request parameters
type RequestParam struct {
	Name  string   `json:"name"`
	Value []string `json:"value"`
}

// Tenant represents the tenant information
type Tenant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// User represents the authenticated user
type User struct {
	ID string `json:"id"`
}

// UserStore represents the user store information
type UserStore struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// AccessToken represents the access token being issued
type AccessToken struct {
	TokenType string       `json:"tokenType"`
	Claims    []TokenClaim `json:"claims"`
	Scopes    []string     `json:"scopes,omitempty"`
}

// RefreshToken represents the refresh token being issued
type RefreshToken struct {
	Claims []TokenClaim `json:"claims"`
}

// TokenClaim represents a token claim with name and value
type TokenClaim struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// AllowedOperation defines permitted operations
type AllowedOperation struct {
	Op    string   `json:"op"`
	Paths []string `json:"paths"`
}

// SuccessResponse represents a successful response
type SuccessResponse struct {
	ActionStatus string      `json:"actionStatus"`
	Operations   []Operation `json:"operations,omitempty"`
}

// FailedResponse represents a failed response
type FailedResponse struct {
	ActionStatus       string `json:"actionStatus"`
	FailureReason      string `json:"failureReason"`
	FailureDescription string `json:"failureDescription"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	ActionStatus     string `json:"actionStatus"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorDescription string `json:"errorDescription"`
}

// Operation represents an operation to perform on token claims/scopes
type Operation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

// CustomClaim represents a custom claim value for add operations
type CustomClaim struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
