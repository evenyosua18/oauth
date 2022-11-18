package builder

import (
	"github.com/evenyosua18/oauth/app/infrastructure/proto/pb"
	"github.com/mitchellh/mapstructure"
)

type RegistrationBuilder struct{}

func (*RegistrationBuilder) RegistrationUserResponse(in interface{}) (interface{}, error) {
	var res *pb.RegistrationUserResponse
	if err := mapstructure.Decode(in, &res); err != nil {
		return nil, err
	}
	return res, nil
}
