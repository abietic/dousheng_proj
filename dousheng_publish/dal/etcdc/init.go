package etcdc

import (
	"dousheng/common/config"

	clientv3 "go.etcd.io/etcd/client/v3"

	"time"
)

var EtcdC *clientv3.Client

func Init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Configs.EtcdConfig.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		panic(err)
	}
	EtcdC = cli
}
