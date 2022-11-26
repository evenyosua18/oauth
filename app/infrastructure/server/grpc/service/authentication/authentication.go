package authentication

import authenticationUC "github.com/evenyosua18/oauth/app/usecase/authentication"

type ServiceAuthentication struct {
	uc authenticationUC.InputPortAuthentication
}

func NewServiceAuthentication(uc authenticationUC.InputPortAuthentication) *ServiceAuthentication {
	return &ServiceAuthentication{uc: uc}
}
