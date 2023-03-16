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

package pack

import (
	"dousheng/userinfo/dal/db"
	"dousheng/userinfo/dal/mongo"
	"dousheng/userinfo/kitex_gen/userinfo"
)

// 转换model数据到展示数据
func UserInfo(us *db.UserStats, ud *mongo.UserData, followed bool) *userinfo.UserInfo {
	if us == nil || ud == nil {
		return nil
	}

	return &userinfo.UserInfo{
		Avatar:          ud.Avatar,
		BackgroundImage: ud.BackgroundImage,
		Signature:       ud.Signature,
		Username:        ud.Username,
		FollowCount:     int64(us.FollowCount),
		TotalFavorited:  int64(us.TotalFavorited),
		FollowerCount:   int64(us.FollowerCount),
		FavoriteCount:   int64(us.FavoriteCount),
		WorkCount:       int64(us.WorkCount),
		IsFollow:        followed,
	}
}
