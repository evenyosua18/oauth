package access_token

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/model"
	"github.com/evenyosua18/oauth/config"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryAccessToken) InsertAccessToken(context context.Context, in interface{}) error {
	//tracer
	_, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//decode to model
	var req *model.AccessToken
	if err := mapstructure.Decode(in, &req); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return err
	}

	//query preparation
	db := r.db

	if config.GetConfig().Server.Debug == constant.True {
		db = db.Debug()
	}

	//call db
	if err := db.Create(&req).Error; err != nil {
		tracer.LogError(sp, tracer.CallDatabase, err)
		return err
	}

	return nil
}
