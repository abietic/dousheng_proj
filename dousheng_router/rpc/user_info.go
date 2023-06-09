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
	"dousheng/router/kitex_gen/userinfo"
	"dousheng/router/kitex_gen/userinfo/userinfoservice"
	userinfoconsts "dousheng/userinfo/consts"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userinfoClient userinfoservice.Client

func initUserInfoRpc() {
	// r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	// r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	r, err := etcd.NewEtcdResolver(config.Configs.EtcdConfig.Endpoints)
	if err != nil {
		panic(err)
	}

	c, err := userinfoservice.NewClient(
		// constants.UserServiceName,
		userinfoconsts.UserInfoServiceName,
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
	userinfoClient = c
}

// 读取用户信息
func GetUserInfo(ctx context.Context, req *userinfo.UesrInfoRequest) (*userinfo.UserInfo, error) {
	resp, err := userinfoClient.GetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserInfo, nil
}
