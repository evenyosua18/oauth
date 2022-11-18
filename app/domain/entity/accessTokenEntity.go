package entity

import "time"

type PasswordGrantRequest struct {
	ClientId     string
	ClientSecret string
	Username     string
	Password     string
	Scopes       string
}

type AccessTokenRequest struct {
	IpAddress     string
	ExpireAt      time.Time
	UserId        string
	OauthClientId string
}

type AccessTokenResponse struct {
	AccessToken  string
	RefreshToken string
	ExpireAt     string
}
