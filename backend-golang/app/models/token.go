package models

//TokenSecurityModel is a blueprint or struct for token jwt security
type TokenSecurityModel struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
