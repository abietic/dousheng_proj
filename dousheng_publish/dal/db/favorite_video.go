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

	"gorm.io/gorm"
)

type FavoriteVideo struct {
	gorm.Model
	UserID  uint64 `json:"user_id"`
	VideoID uint64 `json:"favorite_video_id"`
}

func (u *FavoriteVideo) TableName() string {
	// return constants.UserTableName
	return `favorite_video`
}

// 查询用户所有喜欢的视频
func QueryUserFavoriteVideoList(ctx context.Context, userID uint64) ([]uint64, error) {
	var res []uint64
	if err := DB.WithContext(ctx).Model(&FavoriteVideo{}).Where("user_id = ?", userID).Select("favorite_video_id").Scan(&res).Error; err != nil {
		// if errors.Is(err, gorm.ErrRecordNotFound) {
		// 	return nil, nil
		// }
		return nil, err
	}
	return res, nil
}
