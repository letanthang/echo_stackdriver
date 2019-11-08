package types

import "github.com/dgrijalva/jwt-go"

type errorResponse struct {
	ErrorCode string `json:"err_code"`
	Message   string `json:"message"`
}
type ErrorResponse struct {
	ErrorCode string `json:"err_code,omitempty"`
	Message   string `json:"message,omitempty"`
}
type OkStatus struct {
	Ok bool `json:"ok"`
}

type okResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func PayloadResponse(code string, message string) errorResponse {
	return errorResponse{
		ErrorCode: code,
		Message:   message,
	}
}

type JWTProfileClaims struct {
	AccountType string `json:"typ"`
	ClientName  string `json:"noc,omitempty"`
	ClientPhone string `json:"poc,omitempty"`
	ClientID    int    `json:"cid"`
	GroupID     int    `json:"gid"`
	jwt.StandardClaims
}

type JWTProfilePartnerClaims struct {
	AccountType string `json:"typ"`
	ClientName  string `json:"noc,omitempty"`
	ClientID    int    `json:"cid"`
	PartnerId   int    `json:"pid"`
	GroupID     int    `json:"gid"`
	jwt.StandardClaims
}

type Paging struct {
	Current   int `json:"current"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
	Limit     int `json:"limit"`
	Skip      int `json:"skip"`
}
