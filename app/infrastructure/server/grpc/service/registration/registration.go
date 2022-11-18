package registration

import registrationUC "github.com/evenyosua18/oauth/app/usecase/registration"

type ServiceRegistration struct {
	uc registrationUC.InputPortRegistration
}

func NewServiceRegistration(uc registrationUC.InputPortRegistration) *ServiceRegistration {
	return &ServiceRegistration{uc: uc}
}
