package entity

import "time"

type InsertAccessTokenRequest struct {
	Id            string
	IpAddress     string
	ExpireAt      time.Time
	UserId        string
	OauthClientId string
}

// AccessTokenResponse response get access token
type AccessTokenResponse struct {
	AccessToken  string
	RefreshToken string
	ExpireAt     string
}
