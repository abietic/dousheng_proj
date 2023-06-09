package main

import (
	"dousheng/userinfo/dal"
	userinfo "dousheng/userinfo/kitex_gen/userinfo/userinfoservice"

	// "log"
	"dousheng/common/bound"
	"dousheng/common/config"
	"dousheng/common/middleware"
	tracer2 "dousheng/common/tracer"
	"dousheng/userinfo/consts"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

const (
	EtcdAddress = "127.0.0.1:2379"
)

func Init() {
	config.LoadConfig("./config.yaml")
	// tracer2.InitJaeger(constants.UserServiceName)
	// tracer2.InitJaeger(consts.UserInfoServiceName)
	tracer2.InitJaegerWithConfig(consts.UserInfoServiceName)
	dal.Init()
}

func main() {
	// svr := userinfo.NewServer(new(UserInfoServiceImpl))

	// err := svr.Run()

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	Init()
	// r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
	r, err := etcd.NewEtcdRegistry(config.Configs.EtcdConfig.Endpoints)
	if err != nil {
		panic(err)
	}
	// addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	addr, err := net.ResolveTCPAddr(config.Configs.RpcConfig.ServiceNetwork, config.Configs.RpcConfig.ServiceAddress)
	if err != nil {
		panic(err)
	}
	svr := userinfo.NewServer(new(UserInfoServiceImpl),
		// server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}), // server name
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserInfoServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                              // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r),                                             // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
