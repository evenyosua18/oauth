package oauth

import accessTokenUC "github.com/evenyosua18/oauth/app/usecase/accessToken"

type ServiceAccessToken struct {
	uc accessTokenUC.InputPortAccessToken
}

type ServiceAuthentication struct {
}

func NewServiceAccessToken(uc accessTokenUC.InputPortAccessToken) *ServiceAccessToken {
	return &ServiceAccessToken{uc: uc}
}

func NewServiceAuthentication() *ServiceAuthentication {
	return &ServiceAuthentication{}
}
