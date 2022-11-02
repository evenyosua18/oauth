package constant

// TableName Table
type TableName string

const (
	EndpointTable          TableName = "endpoints"
	AuthorizationCodeTable TableName = "authorization_codes"
	AccessTokenTable       TableName = "access_tokens"
	RefreshTokenTable      TableName = "refresh_tokens"
	OauthClientTable       TableName = "oauth_clients"
	ScopeTable             TableName = "scopes"
	RoleTable              TableName = "roles"
	UserTable              TableName = "users"
)

// General Column
const (
	CreatedAt = `created_at`
	DeletedAt = `deleted_at`
	UpdatedAt = `updated_at`
)
