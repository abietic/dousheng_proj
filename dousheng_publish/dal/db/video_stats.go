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
	"time"

	// "github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"

	"dousheng/publish/dal/mongo"

	"gorm.io/gorm"
)

type VideoStats struct {
	gorm.Model
	AuthorID      uint64 `json:"author_id"`
	CommentCount  uint64 `json:"comment_count"`
	FavoriteCount uint64 `json:"favorite_count"`
}

func (u *VideoStats) TableName() string {
	// return constants.UserTableName
	return `video_stats`
}

// 查询用户发布的所有视频的统计数据
func QueryVideoStatsList(ctx context.Context, authorID uint64) ([]VideoStats, error) {
	var res []VideoStats
	if err := DB.WithContext(ctx).Where("author_id = ?", authorID).Find(&res).Error; err != nil {
		// if errors.Is(err, gorm.ErrRecordNotFound) {
		// 	return nil, nil
		// }
		return nil, err
	}
	return res, nil
}

// 为用户创建一个视频的统计行
func CreateVideoStatsWithTx(tx *gorm.DB, ctx context.Context, videoStats *VideoStats) error {
	return tx.WithContext(ctx).Create(videoStats).Error
}

func CreateVideo(ctx context.Context, authorID uint64, title, playUrl, coverUrl string) (uint64, error) {
	var vid uint64
	err := DB.Transaction(func(tx *gorm.DB) error {
		err := UpdateUserWorkCountWithTx(tx, ctx, 1, authorID)
		if err != nil {
			return err
		}
		videoStats := VideoStats{AuthorID: authorID}
		publishTime := time.Now()
		videoStats.CreatedAt = publishTime
		err = CreateVideoStatsWithTx(tx, ctx, &videoStats)
		if err != nil {
			return err
		}
		videoData := mongo.VideoData{VideoID: uint64(videoStats.ID), AuthorID: authorID, Title: title, PlayUrl: playUrl, CoverUrl: coverUrl, PublishTime: publishTime}
		err = mongo.CreateVideoData(ctx, &videoData)
		if err != nil {
			return err
		}
		vid = uint64(videoStats.ID)
		return nil
	})
	if err != nil {
		return 0, err
	}
	return vid, nil
}
