package authentication

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/domain/entity"
	"github.com/evenyosua18/oauth/util/encryption"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel/trace"
	"strings"
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

	var scope string
	if accessToken.GrantType == constant.ClientCredentials {
		//get oauth client
		oauthClient, err := i.getOauthClient(ctx, sp, accessToken.OauthClientId)

		if err != nil {
			return err
		}

		scope = oauthClient.Scopes

	} else {
		//get user
		user, err := i.getUser(ctx, sp, accessToken.UserId)

		if err != nil {
			return err
		}

		//check role has permission can access the target

		scope = user.Scopes + constant.ScopeSeparator + user.Role.Scopes
	}

	//check request scope is exist
	tracer.LogInfo(sp, scope)
	scopeExist := checkScope(accessToken.Scope, scope)
	tracer.LogResponse(sp, scopeExist)

	if !scopeExist {
		tracer.LogError(sp, tracer.Checking, constant.ErrInvalidScope(""))
		return constant.ErrInvalidScope("")
	}

	//check scope has permission to access target endpoint

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

func (i *InteractionAuthentication) getOauthClient(ctx context.Context, sp trace.Span, id string) (*entity.GetOauthClientResponse, error) {
	//get oauth client by id
	oauthClientResponse, err := i.oauthClient.GetOauthClient(ctx, entity.GetOauthClientRequest{ClientId: id})

	if err != nil {
		tracer.LogError(sp, tracer.CallRepository, err)
		return nil, err
	}

	//decode oauth client
	var oauthClient *entity.GetOauthClientResponse
	if err := mapstructure.Decode(oauthClientResponse, &oauthClient); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}

	return oauthClient, nil
}

func (i *InteractionAuthentication) getUser(ctx context.Context, sp trace.Span, id string) (*entity.GetUserResponse, error) {
	//get user
	userResponse, err := i.user.GetUser(ctx, entity.GetUserRequest{Id: id})

	if err != nil {
		tracer.LogError(sp, tracer.CallRepository, err)
		return nil, err
	}

	//decode
	var user *entity.GetUserResponse
	if err := mapstructure.Decode(userResponse, &user); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}

	return user, nil
}

func checkScope(reqScope, listScope string) bool {
	//split req scope
	listRequestScope := strings.Split(reqScope, constant.ScopeSeparator)

	//loop every request scope
	for _, scope := range listRequestScope {
		if !strings.Contains(scope, listScope) {
			return false
		}
	}

	return true
}
