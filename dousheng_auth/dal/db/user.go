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

package db

import (
	"context"
	"errors"

	// "github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"

	"dousheng/auth/dal/mongo"

	"gorm.io/gorm"
)

const (
	defaut_user_signature        = "这个人很懒什么都没写。"
	defaut_user_avatar           = "https://pkg.go.dev/static/shared/gopher/pilot-bust-1431x901.svg"
	defaut_user_background_image = "https://static.zhihu.com/heifetz/assets/liukanshan-peek.a71ecf3e.png"
)

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	// return constants.UserTableName
	return `user`
}

// 创建用户
func CreateUser(ctx context.Context, users *User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// 查询用户,没找到时user和error都返回nil
func QueryUser(ctx context.Context, userName string) (*User, error) {
	var res User
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Take(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}

func CreateUserWithTx(tx *gorm.DB, ctx context.Context, users *User) error {
	return tx.WithContext(ctx).Create(users).Error
}

func RegisterUser(ctx context.Context, users *User) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		err := CreateUserWithTx(tx, ctx, users)
		if err != nil {
			return err
		}
		var userStats UserStats
		userStats.ID = users.ID
		err = CreateUserStatsWithTx(tx, ctx, &userStats)
		if err != nil {
			return err
		}
		var userData = &mongo.UserData{
			UserID:          uint64(users.ID),
			Username:        users.UserName,
			Avatar:          defaut_user_avatar,
			BackgroundImage: defaut_user_background_image,
			Signature:       defaut_user_signature,
		}
		err = mongo.CreateUserData(ctx, userData)
		if err != nil {
			return err
		}
		return nil
	})
}
