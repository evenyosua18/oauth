package oauth_client

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/model"
	"github.com/evenyosua18/oauth/config"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
)

type request struct {
	ClientId string
}

type response struct {
	Id           string
	ClientId     string
	ClientSecret string
	URI          string
	Scopes       string
	ClientType   string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}

func (r *RepositoryOauthClient) GetOauthClient(context context.Context, in interface{}) (interface{}, error) {
	//tracer
	_, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//mapping request
	var req request
	if err := mapstructure.Decode(in, &req); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}

	//query preparation
	db := r.db

	if config.GetConfig().Server.Debug == constant.True {
		db = db.Debug()
	}

	var oauthClient model.OauthClient
	if err := db.Where(oauthClient.GetClientIdColumn()+" = ?", req.ClientId).Take(&oauthClient).Error; err != nil {
		tracer.LogError(sp, tracer.CallDatabase, err)
		return nil, err
	}
	tracer.LogResponse(sp, oauthClient)

	var deletedAt string
	if oauthClient.DeletedAt != nil {
		deletedAt = oauthClient.DeletedAt.String()
	}

	return &response{
		Id:           oauthClient.Id,
		ClientId:     oauthClient.ClientId,
		ClientSecret: oauthClient.ClientSecret,
		URI:          oauthClient.URI,
		Scopes:       oauthClient.Scopes,
		ClientType:   oauthClient.ClientType,
		CreatedAt:    oauthClient.CreatedAt.String(),
		UpdatedAt:    oauthClient.UpdatedAt.String(),
		DeletedAt:    deletedAt,
	}, nil
}
