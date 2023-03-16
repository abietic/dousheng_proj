package main

import (
	"dousheng/auth/consts"
	auth "dousheng/auth/kitex_gen/auth/userauthservice"

	// "log"
	"dousheng/auth/dal"
	"dousheng/common/bound"
	"dousheng/common/config"
	"dousheng/common/middleware"
	tracer2 "dousheng/common/tracer"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	config.LoadConfig("./config.yaml")
	// tracer2.InitJaeger(constants.UserServiceName)
	tracer2.InitJaegerWithConfig(consts.UserServiceName)
	dal.Init()
}

func main() {
	// svr := auth.NewServer(new(UserAuthServiceImpl))

	// err := svr.Run()

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	// r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})

	Init()
	r, err := etcd.NewEtcdRegistry(config.Configs.EtcdConfig.Endpoints)
	if err != nil {
		panic(err)
	}
	// addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	addr, err := net.ResolveTCPAddr(config.Configs.RpcConfig.ServiceNetwork, config.Configs.RpcConfig.ServiceAddress)
	if err != nil {
		panic(err)
	}

	svr := auth.NewServer(new(UserAuthServiceImpl),
		// server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}), // server name
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                          // middleware
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
