package endpoint

import (
	"context"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/evenyosua18/oauth/app/repository/oauth_db/model"
	"github.com/evenyosua18/oauth/config"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
)

type getEndpointRequest struct {
	EndpointType string
	Name         string
}

type response struct {
	Id           string
	Name         string
	EndpointType string
	Description  string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}

func (r *RepositoryEndpoint) GetEndpoints(context context.Context, in interface{}) (interface{}, error) {
	//tracer
	_, sp := tracer.ChildTracer(context)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//mapping request
	var req getEndpointRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}

	//query preparation
	var endpoint model.Endpoint
	db := r.db

	if config.GetConfig().Server.Debug == constant.True {
		db = db.Debug()
	}

	db = db.Table(string(constant.EndpointTable)).Where(constant.DeletedAt + " IS NULL")

	//query condition
	if req.EndpointType != "" {
		db = db.Where(endpoint.GetEndpointTypeColumn()+" = ?", req.EndpointType)
	}

	if req.Name != "" {
		db = db.Where(endpoint.GetNameColumn()+" LIKE ?", "%"+req.Name+"%")
	}

	//call db
	var endpoints []response
	rows, err := db.Rows()
	if err != nil {
		tracer.LogError(sp, tracer.CallDatabase, err)
		return nil, err
	}

	//loop every rows
	for rows.Next() {
		//scan row
		if err := rows.Scan(&endpoint.Id, &endpoint.Name, &endpoint.EndpointType, &endpoint.Description, &endpoint.CreatedAt, &endpoint.UpdatedAt, &endpoint.DeletedAt); err != nil {
			tracer.LogError(sp, tracer.ScanRow, err)
			return nil, err
		}

		//move to array
		var deletedAt string
		if endpoint.DeletedAt != nil {
			deletedAt = endpoint.DeletedAt.String()
		}
		endpoints = append(endpoints, response{
			Id:           endpoint.Id,
			Name:         endpoint.Name,
			EndpointType: endpoint.EndpointType,
			Description:  endpoint.Description,
			CreatedAt:    endpoint.CreatedAt.String(),
			UpdatedAt:    endpoint.UpdatedAt.String(),
			DeletedAt:    deletedAt,
		})
	}

	defer rows.Close()

	tracer.LogResponse(sp, endpoints)
	return endpoints, nil
}
