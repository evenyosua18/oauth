package entity

type InsertAccessTokenRequest struct {
	Id            string
	IpAddress     string
	ExpireAt      string
	UserId        string
	OauthClientId string
	RefreshToken  string
	GrantType     string
}

// AccessTokenResponse response get access token
type AccessTokenResponse struct {
	AccessToken  string
	RefreshToken string
	ExpireAt     string
}
