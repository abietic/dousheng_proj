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

package rpc

import (
	"context"
	"time"

	"dousheng/common/config"
	"dousheng/common/errno"
	"dousheng/common/middleware"
	publishconsts "dousheng/publish/consts"
	"dousheng/router/kitex_gen/publish"
	"dousheng/router/kitex_gen/publish/publishservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var publishClient publishservice.Client

func initPublishRpc() {
	// r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	// r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	r, err := etcd.NewEtcdResolver(config.Configs.EtcdConfig.Endpoints)
	if err != nil {
		panic(err)
	}

	c, err := publishservice.NewClient(
		// constants.UserServiceName,
		publishconsts.PublishServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	publishClient = c
}

// 用户发布视频
func PublishVideo(ctx context.Context, req *publish.PublishVideoRequest) (uint64, error) {
	resp, err := publishClient.PublishVideo(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return uint64(resp.VideoId), nil
}

// 用户已发布的视频
func ListPublished(ctx context.Context, req *publish.ListVideosRequest) ([]*publish.VideoInfo, error) {
	resp, err := publishClient.ListVideos(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}
