package accessToken

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/domain/entity"
	"github.com/evenyosua18/oauth/app/domain/repository/oauth_db"
	"github.com/evenyosua18/oauth/config"
	"github.com/evenyosua18/oauth/util/encryption"
	"github.com/evenyosua18/oauth/util/str"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/peer"
	"time"
)

const (
	defaultExpiredTime        = "1" //in hours
	defaultLengthExpiredToken = 32  //length of string the refresh token
)

type InputPortAccessToken interface {
	PasswordGrant(ctx context.Context, in interface{}) (interface{}, error)
}

type OutputPortAccessToken interface {
	AccessTokenResponse(in interface{}) (interface{}, error)
}

type InteractionAccessToken struct {
	ExpiredTime        string
	LengthRefreshToken int

	repo     oauth_db.OauthClientRepository
	user     oauth_db.UserRepository
	accToken oauth_db.AccessTokenRepository
	out      OutputPortAccessToken
}

func NewInteractionAccessToken(r oauth_db.OauthClientRepository, u oauth_db.UserRepository, at oauth_db.AccessTokenRepository, o OutputPortAccessToken) *InteractionAccessToken {
	interaction := InteractionAccessToken{
		repo:     r,
		user:     u,
		out:      o,
		accToken: at,
	}

	//expired time
	interaction.ExpiredTime = config.GetConfig().Server.Token.Expired

	if interaction.ExpiredTime == "" {
		interaction.ExpiredTime = defaultExpiredTime
	}

	//length of expired token
	interaction.LengthRefreshToken = config.GetConfig().Server.Token.LengthExpired

	if interaction.LengthRefreshToken == 0 {
		interaction.LengthRefreshToken = defaultLengthExpiredToken
	}

	if interaction.ExpiredTime == "" {
		interaction.ExpiredTime = defaultExpiredTime
	}

	return &interaction
}

// general function

func (i *InteractionAccessToken) generateAccessToken(ctx context.Context, sp trace.Span, username, userId, oauthClientId, grantType string) (response entity.AccessTokenResponse, err error) {
	//get ip
	ip, ok := peer.FromContext(ctx)

	if !ok {
		err = constant.ErrIP
		tracer.LogError(sp, tracer.CallUtility, constant.ErrIP)
		return
	}
	tracer.LogObject(sp, tracer.PrintInformation, ip)

	//generate refresh token
	response.RefreshToken = str.GenerateString(i.LengthRefreshToken)
	tracer.LogObject(sp, tracer.Generator, response.RefreshToken)

	//generate token id
	tokenId, err := uuid.NewUUID()

	if err != nil {
		tracer.LogError(sp, tracer.Generator, err)
		return
	}

	//generate token
	var token string
	var expiredAt time.Time

	token, expiredAt, err = encryption.GenerateToken(i.ExpiredTime, tokenId.String(), username)
	if err != nil {
		tracer.LogError(sp, tracer.Generator, err)
		return
	}
	response.AccessToken = token
	tracer.LogResponse(sp, token)

	//save token & refresh token
	accessToken := entity.InsertAccessTokenRequest{
		Id:            tokenId.String(),
		IpAddress:     ip.Addr.String(),
		ExpireAt:      expiredAt.Format(constant.DefaultDateTimeFormat),
		UserId:        userId,
		OauthClientId: oauthClientId,
		RefreshToken:  response.RefreshToken,
		GrantType:     grantType,
	}

	//add to response struct
	response.ExpireAt = accessToken.ExpireAt

	//call repository
	if err = i.accToken.InsertAccessToken(ctx, accessToken); err != nil {
		tracer.LogError(sp, tracer.CallRepository, err)
		return
	}

	return
}
