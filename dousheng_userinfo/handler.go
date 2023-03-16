package main

import (
	"context"
	"dousheng/common/errno"
	userinfo "dousheng/userinfo/kitex_gen/userinfo"
	"dousheng/userinfo/pack"
	"dousheng/userinfo/service"
)

// UserInfoServiceImpl implements the last service interface defined in the IDL.
type UserInfoServiceImpl struct{}

// GetUserInfo implements the UserInfoServiceImpl interface.
func (s *UserInfoServiceImpl) GetUserInfo(ctx context.Context, req *userinfo.UesrInfoRequest) (resp *userinfo.UserInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(userinfo.UserInfoResponse)
	userInfo, err := service.NewUserInfoService(ctx).UserInfo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserInfo = userInfo

	return resp, nil
}
