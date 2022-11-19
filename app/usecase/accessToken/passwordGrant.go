package accessToken

import (
	"context"
	"fmt"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/domain/entity"
	"github.com/evenyosua18/oauth/util/encryption"
	"github.com/evenyosua18/oauth/util/str"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel/trace"
	"time"
)

func (i *InteractionAccessToken) PasswordGrant(context context.Context, in interface{}) (interface{}, error) {
	//tracer
	ctx, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//set up response
	response := entity.AccessTokenResponse{}

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

	//check scopes

	//get user
	user, err := i.getUser(ctx, sp, req.Username)
	if err != nil {
		return nil, err //tracer already inside the function
	}

	//check password
	if isPasswordOk := encryption.CompareHashString(req.Password, user.Password); !isPasswordOk {
		tracer.LogError(sp, tracer.Checking, constant.ErrInvalidPassword)
		return nil, constant.ErrInvalidPassword
	}

	//generate refresh token
	response.RefreshToken = str.GenerateString(32)
	tracer.LogObject(sp, tracer.Generator, response.RefreshToken)

	//generate token
	tokenId, err := uuid.NewUUID()

	if err != nil {
		tracer.LogError(sp, tracer.Generator, err)
		return nil, err
	}

	token, err := encryption.GenerateToken(i.ExpiredTime, tokenId.String(), user.Name)
	if err != nil {
		tracer.LogError(sp, tracer.Generator, err)
		return nil, err
	}
	response.AccessToken = token
	tracer.LogResponse(sp, token)

	//get claims for get expired time
	claims, err := encryption.ValidateToken(token)
	if err != nil {
		tracer.LogError(sp, tracer.Checking, err)
		return nil, err
	}
	tracer.LogObject(sp, tracer.PrintInformation, claims)
	response.ExpireAt = fmt.Sprintf("%.0f", claims[constant.ClaimsExpired].(float64))

	//save token
	accessToken := entity.InsertAccessTokenRequest{
		Id:            claims[constant.ClaimsId].(string),
		IpAddress:     "",
		ExpireAt:      time.Unix(int64(claims[constant.ClaimsExpired].(float64)), 0),
		UserId:        user.Id,
		OauthClientId: oauthClient.Id,
	}

	if err = i.accToken.InsertAccessToken(ctx, accessToken); err != nil {
		tracer.LogError(sp, tracer.CallRepository, err)
		return nil, err
	}

	//save refresh token

	tracer.LogResponse(sp, response)
	return i.out.AccessTokenResponse(response)
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

func (i *InteractionAccessToken) getUser(ctx context.Context, sp trace.Span, username string) (*entity.GetUserResponse, error) {
	//get user by username or email or phone
	tracer.LogObject(sp, tracer.Before(tracer.CallRepository), username)
	userResponse, err := i.user.GetUser(ctx, entity.GetUserRequest{
		Value: username,
	})

	if err != nil {
		tracer.LogError(sp, tracer.CallDatabase, err)
		return nil, err
	}

	var user *entity.GetUserResponse
	if err := mapstructure.Decode(userResponse, &user); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}

	tracer.LogResponse(sp, user)
	return user, nil
}
