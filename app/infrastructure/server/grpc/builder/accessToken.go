package builder

import (
	"github.com/evenyosua18/oauth/app/infrastructure/proto/pb"
	"github.com/mitchellh/mapstructure"
)

type AccessTokenBuilder struct{}

func (*AccessTokenBuilder) AccessTokenResponse(in interface{}) (interface{}, error) {
	var res *pb.AccessTokenResponse
	if err := mapstructure.Decode(in, &res); err != nil {
		return nil, err
	}
	return res, nil
}
