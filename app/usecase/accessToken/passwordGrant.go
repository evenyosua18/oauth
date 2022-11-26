package accessToken

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/domain/entity"
	"github.com/evenyosua18/oauth/util/encryption"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel/trace"
	"strings"
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
	oauthClient, err := i.getOauthClient(ctx, sp, req.ClientId)
	if err != nil {
		return nil, err //tracer already inside the function
	}

	//check client secret
	if req.ClientSecret != oauthClient.ClientSecret {
		tracer.LogError(sp, tracer.Checking, constant.ErrInvalidClientSecret)
		return nil, constant.ErrInvalidClientSecret
	}

	//get user
	user, err := i.getUser(ctx, sp, req.Username, req.Scopes)
	if err != nil {
		return nil, err //tracer already inside the function
	}

	//check password
	if isPasswordOk := encryption.CompareHashString(req.Password, user.Password); !isPasswordOk {
		tracer.LogError(sp, tracer.Checking, constant.ErrInvalidPassword)
		return nil, constant.ErrInvalidPassword
	}

	//manage access token
	accessToken, err := i.generateAccessToken(ctx, sp, user.Name, user.Id, oauthClient.Id)
	if err != nil {
		return nil, err
	}

	tracer.LogResponse(sp, accessToken)
	return i.out.AccessTokenResponse(accessToken)
}

func (i *InteractionAccessToken) getOauthClient(ctx context.Context, sp trace.Span, clientId string) (*entity.GetOauthClientResponse, error) {
	//get oauth client
	tracer.LogObject(sp, tracer.Before(tracer.CallRepository), clientId)
	oauthClientResponse, err := i.repo.GetOauthClient(ctx, &entity.GetOauthClientRequest{
		ClientId: clientId,
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
	return oauthClient, nil
}

func (i *InteractionAccessToken) getUser(ctx context.Context, sp trace.Span, username, reqScopes string) (*entity.GetUserResponse, error) {
	//get user by username or email or phone
	tracer.LogObject(sp, tracer.Before(tracer.CallRepository), username)
	userResponse, err := i.user.GetUser(ctx, entity.GetUserRequest{
		Value: username,
	})

	if err != nil {
		tracer.LogError(sp, tracer.CallDatabase, err)
		return nil, err
	}

	//decode
	var user *entity.GetUserResponse
	if err := mapstructure.Decode(userResponse, &user); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}

	//check is active
	if !user.IsActive {
		tracer.LogError(sp, tracer.Checking, constant.ErrInactiveUser)
		return nil, constant.ErrInactiveUser
	}

	//check scopes
	scopes := strings.Split(reqScopes, constant.ScopeSeparator)

	for _, scope := range scopes {
		if !strings.Contains(user.Scopes, scope) {
			return nil, constant.ErrInvalidScope(scope)
		}
	}

	tracer.LogResponse(sp, user)
	return user, nil
}
