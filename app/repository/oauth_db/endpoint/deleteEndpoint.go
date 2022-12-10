package endpoint

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/model"
	"github.com/evenyosua18/oauth/config"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type deleteEndpointRequest struct {
	Id         string
	SoftDelete bool
}

func (r *RepositoryEndpoint) DeleteEndpoint(ctx context.Context, in interface{}) error {
	//tracer
	ctx, sp := tracer.ChildTracer(ctx)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//decode
	var req deleteEndpointRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return err
	}

	//query preparation
	db := r.db

	if config.GetConfig().Server.Debug == constant.True {
		db = db.Debug()
	}

	if !req.SoftDelete {
		db = db.Unscoped()
	}

	//call db
	db = db.Where("id = ?", req.Id).Delete(&model.Endpoint{})

	if db.Error != nil {
		tracer.LogError(sp, tracer.CallDatabase, db.Error)
		return db.Error
	}

	tracer.LogResponse(sp, db.RowsAffected)

	if db.RowsAffected == 0 {
		tracer.LogError(sp, tracer.AfterQuery, gorm.ErrRecordNotFound)
		return gorm.ErrRecordNotFound
	}

	return nil
}
