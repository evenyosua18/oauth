package accessToken

import (
	"context"
	"github.com/evenyosua18/oauth/app/domain/entity"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
)

func (i *InteractionAccessToken) PasswordGrant(context context.Context, in interface{}) (interface{}, error) {
	//tracer
	ctx, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//decode
	var req *entity.PasswordGrantRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}
	tracer.LogRequest(sp, req)

	//get oauth client
	oauthClientResponse, err := i.repo.GetOauthClient(ctx, &entity.GetOauthClientRequest{
		ClientId: req.ClientId,
	})

	if err != nil {
		tracer.LogError(sp, tracer.CallRepository, err)
		return nil, err
	}
	tracer.LogResponse(sp, oauthClientResponse)

	var oauthClient *entity.GetOauthClientResponse
	if err := mapstructure.Decode(oauthClientResponse, &oauthClient); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}
	tracer.LogResponse(sp, oauthClient)

	//check client secret

	//check scopes

	return i.out.AccessTokenResponse(&entity.AccessTokenResponse{})
}
