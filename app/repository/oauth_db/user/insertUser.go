package user

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/model"
	"github.com/evenyosua18/oauth/config"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
)

type response struct {
	Id string
}

func (r *RepositoryUser) InsertUser(context context.Context, in interface{}) (interface{}, error) {
	//tracer
	_, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//decode
	var user *model.User
	if err := mapstructure.Decode(in, &user); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}

	//query preparation
	db := r.db

	if config.GetConfig().Server.Debug == constant.True {
		db = db.Debug()
	}

	if err := db.Create(&user).Error; err != nil {
		tracer.LogError(sp, tracer.CallDatabase, err)
		return nil, err
	}

	tracer.LogResponse(sp, user)
	return &response{Id: user.Id}, nil
}
