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

type GetAccessTokenRequest struct {
	Id string
}

type GetAccessTokenResponse struct {
	Id            string
	IpAddress     string
	ExpireAt      string
	UserId        string
	OauthClientId string
	GrantType     string
	CreatedAt     string
	UpdateAt      string
	DeletedAt     string
}

// AccessTokenResponse response get access token
type AccessTokenResponse struct {
	AccessToken  string
	RefreshToken string
	ExpireAt     string
}
