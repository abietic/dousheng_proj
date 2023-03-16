package etcdc

import (
	clientv3 "go.etcd.io/etcd/client/v3"

	"dousheng/common/config"
	"time"
)

var EtcdC *clientv3.Client

func Init() {
	cli, err := clientv3.New(clientv3.Config{
		// Endpoints:   []string{"127.0.0.1:2379"},
		Endpoints:   config.Configs.EtcdConfig.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		panic(err)
	}
	EtcdC = cli
}
