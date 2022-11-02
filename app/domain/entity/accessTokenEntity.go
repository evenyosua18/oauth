package entity

type PasswordGrantRequest struct {
	ClientId     string
	ClientSecret string
	Username     string
	Password     string
	Scopes       string
}

type AccessTokenResponse struct {
	AccessToken  string
	RefreshToken string
	ExpireAt     string
}
