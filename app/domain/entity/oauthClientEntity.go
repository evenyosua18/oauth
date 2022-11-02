package entity

type GetOauthClientRequest struct {
	ClientId string
}

type GetOauthClientResponse struct {
	Id           string
	ClientId     string
	ClientSecret string
	URI          string
	Scopes       string
	ClientType   string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}
