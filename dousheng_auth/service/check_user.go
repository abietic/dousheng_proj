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
	"crypto/md5"
	"fmt"
	"io"

	"dousheng/common/errno"

	"dousheng/auth/kitex_gen/auth"

	"dousheng/auth/dal/db"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// 用户的登录服务
func (s *CheckUserService) CheckUser(req *auth.LoginRequest) (int64, error) {
	// 将传入的用户密码进行加密
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	// 根据用户名到数据库中查找用户信息
	userName := req.Username
	user, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if user == nil {
		return 0, errno.UserNotExistErr
	}
	u := user
	if u.Password != passWord {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(u.ID), nil
}
