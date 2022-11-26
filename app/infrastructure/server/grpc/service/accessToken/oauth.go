package accessToken

import accessTokenUC "github.com/evenyosua18/oauth/app/usecase/accessToken"

type ServiceAccessToken struct {
	uc accessTokenUC.InputPortAccessToken
}

func NewServiceAccessToken(uc accessTokenUC.InputPortAccessToken) *ServiceAccessToken {
	return &ServiceAccessToken{uc: uc}
}
