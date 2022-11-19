package usecase

import (
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/infrastructure/server/grpc/builder"
	"github.com/evenyosua18/oauth/app/repository"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/access_token"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/endpoint"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/oauth_client"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/user"
	accessTokenUC "github.com/evenyosua18/oauth/app/usecase/accessToken"
	endpointUC "github.com/evenyosua18/oauth/app/usecase/endpoint"
	registrationUC "github.com/evenyosua18/oauth/app/usecase/registration"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, err := di.NewBuilder()

	if err != nil {
		panic(err)
	}

	if err = builder.Add([]di.Def{
		{Name: string(constant.EndpointCTN), Build: endpointInteraction},
		{Name: string(constant.AccessTokenCTN), Build: accessTokenInteraction},
		{Name: string(constant.RegistrationCTN), Build: registrationInteraction},
	}...); err != nil {
		panic(err)
	}

	return &Container{ctn: builder.Build()}
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

/*interaction container*/

// endpoint interaction
func endpointInteraction(_ di.Container) (interface{}, error) {
	repo := endpoint.NewEndpointRepository(repository.OauthDB)
	out := &builder.EndpointBuilder{}
	return endpointUC.NewInteractionEndpoint(repo, out), nil
}

// access token interaction
func accessTokenInteraction(_ di.Container) (interface{}, error) {
	repo := oauth_client.NewOauthClientRepository(repository.OauthDB)
	userRepo := user.NewUserRepository(repository.OauthDB)
	accessTokenRepo := access_token.NewAccessTokenRepository(repository.OauthDB)
	out := &builder.AccessTokenBuilder{}
	return accessTokenUC.NewInteractionAccessToken(repo, userRepo, accessTokenRepo, out), nil
}

// registration interaction
func registrationInteraction(_ di.Container) (interface{}, error) {
	repo := user.NewUserRepository(repository.OauthDB)
	out := &builder.RegistrationBuilder{}
	return registrationUC.NewInteractionRegistration(repo, out), nil
}
