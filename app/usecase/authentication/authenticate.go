package authentication

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/domain/entity"
	"github.com/evenyosua18/oauth/util/encryption"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
)

func (i *InteractionAuthentication) Authenticate(context context.Context, in interface{}) error {
	//tracer
	ctx, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogResponse(sp, in)

	//decode
	var req *entity.AuthenticationRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return err
	}

	//extract claims
	claims, err := encryption.ValidateToken(req.AccessToken)
	if err != nil {
		tracer.LogError(sp, tracer.Checking, err)
		return err
	}
	tracer.LogObject(sp, tracer.PrintInformation, claims)

	//get access token
	accessToken, err := i.accessToken.GetAccessToken(ctx, entity.GetAccessTokenRequest{Id: claims[constant.ClaimsId].(string)})
	if err != nil {
		tracer.LogError(sp, tracer.CallRepository, err)
		return err
	}
	tracer.LogObject(sp, tracer.CallRepository, accessToken)

	return nil
}
