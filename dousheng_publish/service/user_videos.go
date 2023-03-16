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

	"dousheng/publish/kitex_gen/publish"

	"dousheng/publish/dal/db"
	"dousheng/publish/dal/mongo"
	"dousheng/publish/pack"

	"github.com/cloudwego/kitex/pkg/klog"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new PublishListService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{
		ctx: ctx,
	}
}

func (s *PublishListService) fetchVideoDataList(hostUserID uint64) ([]mongo.VideoData, error) {
	videoDataList, err := mongo.QueryVideoDataList(s.ctx, uint64(hostUserID))
	if err != nil {
		klog.Error(err)
		return nil, errno.GetVideoDataListErr
	}
	return videoDataList, nil
}

func (s *PublishListService) fetchVideoStatusList(hostUserID uint64) ([]db.VideoStats, error) {
	videoStatusList, err := db.QueryVideoStatsList(s.ctx, hostUserID)
	if err != nil {
		klog.Error(err)
		return nil, errno.GetVideoStatsListErr
	}
	return videoStatusList, nil
}

func (s *PublishListService) fetchUserFavoriteVideoList(visitorUserID uint64) ([]uint64, error) {
	favoriteList, err := db.QueryUserFavoriteVideoList(s.ctx, visitorUserID)
	if err != nil {
		klog.Error(err)
		return nil, errno.GetUserFavoriteVideoListErr
	}
	return favoriteList, nil
}

const (
	remote_call_count = 3
)

// 获取用户发布的所有视频及视频信息
func (s *PublishListService) GetPublishList(req *publish.ListVideosRequest) ([]*publish.VideoInfo, error) {
	var videoStatusList []db.VideoStats
	var mapping map[uint64]*mongo.VideoData
	var favoriteSet map[uint64]struct{}
	resChannel := make(chan interface{}, remote_call_count)

	go func() {
		if res, err := s.fetchVideoStatusList(uint64(req.HostUserId)); err != nil {
			resChannel <- err
		} else {
			resChannel <- res
		}
	}()
	go func() {
		if res, err := s.fetchVideoDataList(uint64(req.HostUserId)); err != nil {
			resChannel <- err
		} else {
			vdl := len(res)
			mapping := make(map[uint64]*mongo.VideoData, vdl)
			for i := 0; i < vdl; i++ {
				mapping[res[i].VideoID] = &res[i]
			}
			resChannel <- mapping
		}
	}()
	go func() {
		if res, err := s.fetchUserFavoriteVideoList(uint64(req.VisitorUserId)); err != nil {
			resChannel <- err
		} else {
			favoriteSet := make(map[uint64]struct{}, len(res))
			for _, vid := range res {
				favoriteSet[vid] = struct{}{}
			}
			resChannel <- favoriteSet
		}
	}()

	for i := 0; i < remote_call_count; i++ {
		vo := <-resChannel
		switch v := vo.(type) {
		case errno.ErrNo:
			return nil, v
		case []db.VideoStats:
			videoStatusList = v
		case map[uint64]*mongo.VideoData:
			mapping = v
		case map[uint64]struct{}:
			favoriteSet = v
		default:
			return nil, errno.GetVideoInfoListNotConsistErr
		}
	}

	vsl := len(videoStatusList)

	res := make([]*publish.VideoInfo, vsl)

	for i := 0; i < vsl; i++ {
		vdp, exist := mapping[uint64(videoStatusList[i].ID)]
		if !exist {
			return nil, errno.GetVideoInfoListNotConsistErr
		}
		_, isFavorite := favoriteSet[uint64(videoStatusList[i].ID)]
		res[i] = pack.VideoInfo(&videoStatusList[i], vdp, isFavorite)
	}
	return res, nil
}
