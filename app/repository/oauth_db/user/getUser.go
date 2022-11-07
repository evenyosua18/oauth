package user

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/model"
	"github.com/evenyosua18/oauth/config"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
)

type request struct {
	Id    string
	Value string
}

func (r *RepositoryUser) GetUser(context context.Context, in interface{}) (interface{}, error) {
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

	//query preparation
	var user model.User
	db := r.db

	if config.GetConfig().Server.Debug == constant.True {
		db = db.Debug()
	}

	//query condition
	if req.Id != "" {
		db = db.Where(user.GetIdColumn()+" = ?", req.Id)
	}

	if req.Value != "" {
		db = db.Where(user.GetEmailColumn()+" = ? OR "+user.GetPhoneColumn()+" = ? OR "+user.GetNameColumn()+" = ?", req.Value, req.Value, req.Value)
	}

	//call db
	if err := db.Take(user).Error; err != nil {
		tracer.LogError(sp, tracer.CallDatabase, err)
		return nil, err
	}

	tracer.LogResponse(sp, user)
	return nil, nil
}
