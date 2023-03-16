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

	// "github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"

	"gorm.io/gorm"
)

type UserFollowing struct {
	gorm.Model
	UserID         uint64 `json:"user_id"`
	FollowedUserID uint64 `json:"followed_user_id"`
}

func (u *UserFollowing) TableName() string {
	// return constants.UserTableName
	return `user_followed`
}

// 查询用户关注的所有用户
func QueryUserFollowed(ctx context.Context, userID uint64) ([]uint64, error) {
	var res []uint64
	if err := DB.WithContext(ctx).Model(&UserFollowing{}).Where("user_id = ?", userID).Select("followed_user_id").Scan(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
