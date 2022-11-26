package authentication

import (
	"context"
	"github.com/evenyosua18/oauth/app/domain/entity"
	"github.com/evenyosua18/oauth/util/encryption"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
)

func (u *InteractionAuthentication) Authenticate(context context.Context, in interface{}) error {
	//tracer
	_, sp := tracer.ChildTracer(context)
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

	return nil
}
