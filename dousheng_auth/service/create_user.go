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
	"dousheng/auth/dal/etcdc"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func genJob(ctx context.Context, req *auth.RegisterRequest) etcdc.NeedDistLockingJob {
	return func() (interface{}, error) {
		user, err := db.QueryUser(ctx, req.Username)
		if err != nil {
			return nil, err
		}
		// TODO:未来要加入分布式锁
		if user != nil {
			return nil, errno.UserAlreadyExistErr
		}

		h := md5.New()
		if _, err = io.WriteString(h, req.Password); err != nil {
			return nil, err
		}
		passWord := fmt.Sprintf("%x", h.Sum(nil))
		res := &db.User{
			UserName: req.Username,
			Password: passWord,
		}
		err = db.RegisterUser(ctx, res)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *auth.RegisterRequest) (*db.User, error) {
	job := genJob(s.ctx, req)
	lockKey := "/creatuser/" + req.Username
	res, err := etcdc.RunJobWithDistLock(s.ctx, lockKey, job)
	if err != nil {
		return nil, err
	}
	return res.(*db.User), nil
}
