package config

import (
	"github.com/spf13/viper"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var Configs struct {
	MysqlConfig  MysqlConfigStruct
	MogoDBConfig MogoDBConfigStruct
	EtcdConfig   EtcdConfigStruct
	RpcConfig    RpcConfigStruct
	JaegerConfig jaegercfg.Configuration
	HertzConfig  HertzConfigStruct
	MinioConfig  MinioConfigStruct
}

func LoadConfig(filePath string) {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.UnmarshalKey("configs", &Configs); err != nil {
		panic(err)
	}
	// logrus.Info(Configs)
	// fmt.Print(Configs)
}
