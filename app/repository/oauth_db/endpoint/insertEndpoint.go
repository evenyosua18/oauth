package endpoint

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/model"
	"github.com/evenyosua18/oauth/config"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
)

func (r *RepositoryEndpoint) InsertEndpoint(context context.Context, in interface{}) (interface{}, error) {
	//tracer
	_, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//decode
	var endpoint *model.Endpoint
	if err := mapstructure.Decode(in, &endpoint); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}

	//query preparation
	db := r.db

	if config.GetConfig().Server.Debug == constant.True {
		db = db.Debug()
	}

	//call db
	if err := db.Create(&endpoint).Error; err != nil {
		tracer.LogError(sp, tracer.CallDatabase, err)
		return nil, err
	}

	tracer.LogResponse(sp, endpoint)
	return endpoint, nil
}
