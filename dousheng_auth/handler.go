package main

import (
	"context"
	auth "dousheng/auth/kitex_gen/auth"
	"dousheng/auth/pack"
	"dousheng/auth/service"
	"dousheng/common/errno"
)

// UserAuthServiceImpl implements the last service interface defined in the IDL.
type UserAuthServiceImpl struct{}

// Register implements the UserAuthServiceImpl interface.
func (s *UserAuthServiceImpl) Register(ctx context.Context, req *auth.RegisterRequest) (resp *auth.RegisterResponse, err error) {
	// TODO: Your code here...

	resp = new(auth.RegisterResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userModel, err := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserAuth = pack.User(userModel)
	return resp, nil
}

// Login implements the UserAuthServiceImpl interface.
func (s *UserAuthServiceImpl) Login(ctx context.Context, req *auth.LoginRequest) (resp *auth.LoginResponse, err error) {
	// TODO: Your code here...

	resp = new(auth.LoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserAuth = &auth.UserAuth{UserId: uid, Username: req.Username}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
