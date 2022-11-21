package access_token

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/model"
	"github.com/evenyosua18/oauth/config"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
	"time"
)

type insertAccessToken struct {
	Id            string
	IpAddress     string
	ExpireAt      string
	UserId        string
	OauthClientId string
	RefreshToken  string
}

func (r *RepositoryAccessToken) InsertAccessToken(context context.Context, in interface{}) error {
	//tracer
	_, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//decode to model
	var req *insertAccessToken
	if err := mapstructure.Decode(in, &req); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return err
	}
	tracer.LogObject(sp, tracer.DecodeObject, req)

	//create access token model
	expireAt, err := time.Parse(constant.DefaultDateTimeFormat, req.ExpireAt)

	if err != nil {
		tracer.LogError(sp, tracer.Time, err)
		return err
	}

	accessToken := model.AccessToken{
		Id:            req.Id,
		IpAddress:     req.IpAddress,
		ExpireAt:      expireAt,
		UserId:        req.UserId,
		OauthClientId: req.OauthClientId,
		RefreshTokens: []model.RefreshToken{
			{
				RefreshToken:  req.RefreshToken,
				AccessTokenId: req.Id,
			},
		},
	}

	//query preparation
	db := r.db

	if config.GetConfig().IsDebugMode() {
		db = db.Debug()
	}

	//call db
	if err := db.Create(&accessToken).Error; err != nil {
		tracer.LogError(sp, tracer.CallDatabase, err)
		return err
	}

	return nil
}
