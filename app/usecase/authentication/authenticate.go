package authentication

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/domain/entity"
	"github.com/evenyosua18/oauth/util/encryption"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel/trace"
	"time"
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
	accessToken, err := i.getAccessToken(ctx, sp, claims[constant.ClaimsId].(string))
	if err != nil {
		return err
	}

	//check expired time
	if time.Unix(int64(claims[constant.ClaimsExpired].(float64)), 0).Format(constant.DefaultDateTimeFormat) != accessToken.ExpireAt {
		tracer.LogError(sp, tracer.Checking, constant.ErrInvalidToken(constant.ErrMessageInvalidExpiredTime))
		return constant.ErrInvalidToken(constant.ErrMessageInvalidExpiredTime)
	}

	tracer.LogResponse(sp, accessToken)
	return nil
}

func (i *InteractionAuthentication) getAccessToken(ctx context.Context, sp trace.Span, id string) (*entity.GetAccessTokenResponse, error) {
	//get access token by id
	tracer.LogObject(sp, tracer.Before(tracer.CallRepository), id)
	accessTokenResponse, err := i.accessToken.GetAccessToken(ctx, entity.GetAccessTokenRequest{Id: id})
	if err != nil {
		tracer.LogError(sp, tracer.CallRepository, err)
		return nil, err
	}
	tracer.LogResponse(sp, accessTokenResponse)

	//decode
	var accessToken *entity.GetAccessTokenResponse
	if err := mapstructure.Decode(accessTokenResponse, &accessToken); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}

	tracer.LogResponse(sp, accessToken)
	return accessToken, nil
}
