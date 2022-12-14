package constant

// layer
const (
	RepositoryLayer  = `Repository`
	InteractionLayer = `Interaction`
	ServiceLayer     = `Service`
)

// string boolean
const (
	True  = `true`
	False = `false`
)

// claims
const (
	ClaimsId       = `jti`
	ClaimsExpired  = `exp`
	ClaimsUsername = `username`
)

// date time format
const (
	DefaultDateFormat     = "2006-02-01"                                //dd-mm-yyyy
	DefaultTimeFormat     = "15:04:05"                                  //hh:mm:ss
	DefaultDateTimeFormat = DefaultDateFormat + " " + DefaultTimeFormat //dd-mm-yyyy hh:mm:ss
	ReadableDateFormat    = "02 Jan 2006"
)

// scope separator
const (
	ScopeSeparator = ","
)

// grant type
const (
	Password          = `password`
	ClientCredentials = `client`
	AuthorizationCode = `authorization`
	ImplicitGrant     = `implicit`
)
