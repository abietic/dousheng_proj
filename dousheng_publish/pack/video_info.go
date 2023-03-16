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
	"dousheng/publish/dal/db"
	"dousheng/publish/dal/mongo"
	"dousheng/publish/kitex_gen/publish"
)

// 转换model数据到展示数据
func VideoInfo(vs *db.VideoStats, vd *mongo.VideoData, isFavorite bool) *publish.VideoInfo {
	if vs == nil || vd == nil {
		return nil
	}

	return &publish.VideoInfo{
		VideoId:       int64(vs.ID),
		PlayUrl:       vd.PlayUrl,
		CoverUrl:      vd.CoverUrl,
		Title:         vd.Title,
		FavoriteCount: int64(vs.FavoriteCount),
		CommentCount:  int64(vs.CommentCount),
		IsFavorite:    isFavorite,
	}
}
