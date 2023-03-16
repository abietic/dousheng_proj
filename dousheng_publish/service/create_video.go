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

	"dousheng/publish/kitex_gen/publish"

	"dousheng/publish/dal/db"
)

type CreateVideoService struct {
	ctx context.Context
}

// NewCreateVideoService new CreateVideoService
func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

// 指定用户发布视频
func (s *CreateVideoService) CreateVideo(req *publish.PublishVideoRequest) (uint64, error) {
	vid, err := db.CreateVideo(s.ctx, uint64(req.UserId), req.Title, req.VideoUrl, req.CoverUrl)
	if err != nil {
		return 0, err
	}
	return vid, nil
}
