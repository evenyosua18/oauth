package entity

// PasswordGrantRequest request for get access token by password grant
type PasswordGrantRequest struct {
	ClientId     string
	ClientSecret string
	Username     string
	Password     string
	Scopes       string
}
