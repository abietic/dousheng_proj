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

// 这个包来自cloudwego，主要功能是使用log和ps在cpu利用率大于80%时返回错误
// 这是作为一个kitex server的WithBoundHandle参数输入用的，kitex server会在入站访问时检查
package bound

import (
	"context"
	"fmt"
	"net"

	"dousheng/common/errno"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/shirou/gopsutil/cpu"
)

var _ remote.InboundHandler = &cpuLimitHandler{}

const (
	CPURateLimit float64 = 80.0
)

type cpuLimitHandler struct{}

func NewCpuLimitHandler() remote.InboundHandler {
	return &cpuLimitHandler{}
}

// OnActive implements the remote.InboundHandler interface.
func (c *cpuLimitHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
	return ctx, nil
}

// OnRead implements the remote.InboundHandler interface.
func (c *cpuLimitHandler) OnRead(ctx context.Context, conn net.Conn) (context.Context, error) {
	p := cpuPercent()
	klog.CtxInfof(ctx, "current cpu is %.2g", p)
	if p > CPURateLimit {
		return ctx, errno.ServiceErr.WithMessage(fmt.Sprintf("cpu = %.2g", c))
	}
	return ctx, nil
}

// OnInactive implements the remote.InboundHandler interface.
func (c *cpuLimitHandler) OnInactive(ctx context.Context, conn net.Conn) context.Context {
	return ctx
}

// OnMessage implements the remote.InboundHandler interface.
func (c *cpuLimitHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	return ctx, nil
}

func cpuPercent() float64 {
	percent, _ := cpu.Percent(0, false)
	return percent[0]
}
