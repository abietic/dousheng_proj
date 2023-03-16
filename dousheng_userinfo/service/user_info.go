// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package service

import (
	"context"

	"dousheng/common/errno"
	"dousheng/userinfo/dal/mongo"

	"dousheng/userinfo/kitex_gen/userinfo"

	"dousheng/userinfo/dal/db"
	"dousheng/userinfo/pack"
)

type UserInfoService struct {
	ctx context.Context
}

// NewUserInfoService new UserInfoService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{
		ctx: ctx,
	}
}

const (
	remote_call_count = 3
)

func (s *UserInfoService) visitorFollowingHost(visitorID, hostID uint64) (bool, error) {
	if visitorID == hostID {
		return false, nil
	}
	res, err := db.QueryUserFollowed(s.ctx, visitorID)
	if err != nil {
		err = errno.GetUserFollowedErr
		return false, err
	}
	for _, uid := range res {
		if uid == hostID {
			return true, nil
		}
	}
	return false, nil
}

func (s *UserInfoService) getUserStats(hostID uint64) (*db.UserStats, error) {
	res, err := db.QueryUserStats(s.ctx, hostID)
	if err != nil {
		err = errno.GetUserStatsErr
		return nil, err
	}
	return res, nil
}

func (s *UserInfoService) getUserData(hostID uint64) (*mongo.UserData, error) {
	res, err := mongo.QueryUserData(s.ctx, hostID)
	if err != nil {
		err = errno.GetUserDataErr
		return nil, err
	}
	return res, nil
}

// 查找用户的相关信息
func (s *UserInfoService) UserInfo(req *userinfo.UesrInfoRequest) (*userinfo.UserInfo, error) {
	resChan := make(chan interface{}, remote_call_count)
	// 读取用户的统计数据
	go func() {
		res, err := s.getUserStats(uint64(req.HostUserId))
		if err != nil {
			resChan <- err
			return
		}
		resChan <- res
	}()
	// 读取用户的个人信息
	go func() {
		res, err := s.getUserData(uint64(req.HostUserId))
		if err != nil {
			resChan <- err
			return
		}
		resChan <- res
	}()
	// 读取用户的关注关系
	go func() {
		res, err := s.visitorFollowingHost(uint64(req.VisitorUserId), uint64(req.HostUserId))
		if err != nil {
			resChan <- err
			return
		}
		resChan <- res
	}()
	var us *db.UserStats
	var ud *mongo.UserData
	var followed bool
	// 将所有查询结果收集起来整合为结果
	for i := 0; i < remote_call_count; i++ {
		resOrErr := <-resChan
		switch v := resOrErr.(type) {
		case errno.ErrNo:
			return nil, v
		case *mongo.UserData:
			ud = v
		case *db.UserStats:
			us = v
		case bool:
			followed = v
		default:
			return nil, errno.ServiceErr
		}
	}
	res := pack.UserInfo(us, ud, followed)
	return res, nil
}
