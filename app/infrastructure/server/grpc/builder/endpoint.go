package builder

import (
	"github.com/evenyosua18/oauth/app/infrastructure/server/grpc/proto/pb"
	"github.com/mitchellh/mapstructure"
)

type EndpointBuilder struct{}

func (*EndpointBuilder) GetEndpointsResponse(in interface{}) (interface{}, error) {
	var res *pb.GetEndpointsResponse
	if err := mapstructure.Decode(in, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (*EndpointBuilder) InsertEndpointResponse(in interface{}) (interface{}, error) {
	var res *pb.InsertEndpointResponse
	if err := mapstructure.Decode(in, &res); err != nil {
		return nil, err
	}

	return res, nil
}
