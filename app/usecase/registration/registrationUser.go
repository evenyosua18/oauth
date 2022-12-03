package registration

import (
	"context"
	"github.com/evenyosua18/oauth/app/domain/entity"
	"github.com/evenyosua18/oauth/util/encryption"
	"github.com/evenyosua18/oauth/util/tracer"
	"github.com/mitchellh/mapstructure"
)

type insertUser struct {
	Name        string
	Password    string
	Phone       string
	Email       string
	RoleId      string
	IsActive    bool
	IsSuperRole bool
}

func (i *InteractionRegistration) RegistrationUser(ctx context.Context, in interface{}) (interface{}, error) {
	//tracer
	ctx, sp := tracer.ChildTracer(ctx)
	defer sp.End()
	tracer.LogRequest(sp, in)

	//decode
	var req *entity.RegistrationUserRequest
	if err := mapstructure.Decode(in, &req); err != nil {
		tracer.LogError(sp, tracer.DecodeObject, err)
		return nil, err
	}

	insertUser := insertUser{
		Name:        req.Name,
		Phone:       req.Phone,
		Email:       req.Email,
		RoleId:      i.DefaultRoleId,
		IsActive:    true,  //default
		IsSuperRole: false, //default
	}

	//hash password
	newPassword, err := encryption.HashString(req.Password)
	if err != nil {
		tracer.LogError(sp, tracer.CallUtility, err)
		return nil, err
	}

	insertUser.Password = newPassword

	//call repository
	tracer.LogObject(sp, tracer.Before(tracer.CallRepository), insertUser)
	insertUserResponse, err := i.repo.InsertUser(ctx, insertUser)
	if err != nil {
		tracer.LogError(sp, tracer.CallRepository, err)
		return nil, err
	}

	tracer.LogResponse(sp, insertUserResponse)
	return i.out.RegistrationUserResponse(insertUserResponse)
}
