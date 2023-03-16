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

	"gorm.io/gorm"
)

type UserStats struct {
	gorm.Model
	// UserID         uint64 `json:"user_id"`
	WorkCount      uint64 `json:"work_count"`
	FollowCount    uint64 `json:"follow_count"`
	FollowerCount  uint64 `json:"follower_count"`
	FavoriteCount  uint64 `json:"favorite_count"`
	TotalFavorited uint64 `json:"total_favorited"`
}

func (u *UserStats) TableName() string {
	// return constants.UserTableName
	return `user_stats`
}

// 查询用户,没找到时user和error都返回nil
func QueryUserStats(ctx context.Context, userID uint64) (*UserStats, error) {
	var res UserStats
	if err := DB.WithContext(ctx).Where("id = ?", userID).Take(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}
