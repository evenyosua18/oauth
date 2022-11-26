package access_token

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/model"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
)

type request struct {
	Id string
}

type response struct {
	Id            string
	IpAddress     string
	ExpireAt      string
	UserId        string
	OauthClientId string
	RefreshToken  string
	GrantType     string
}

func (r *RepositoryAccessToken) GetAccessToken(context context.Context, in interface{}) (interface{}, error) {
	//tracer
	_, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//decode
	var req request
	if err := mapstructure.Decode(in, &req); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}

	//call db
	var accessToken model.AccessToken
	if err := r.db.Where(accessToken.GetColumnId()+" = ?", req.Id).Take(&accessToken).Error; err != nil {
		tracer.LogError(sp, tracer.CallDatabase, err)
		return nil, err
	}
	tracer.LogResponse(sp, accessToken)
	return &response{
		Id:            accessToken.Id,
		IpAddress:     accessToken.IpAddress,
		ExpireAt:      accessToken.ExpireAt.Format(constant.DefaultDateTimeFormat),
		UserId:        accessToken.UserId,
		OauthClientId: accessToken.OauthClientId,
		GrantType:     accessToken.GrantType,
	}, nil
}
