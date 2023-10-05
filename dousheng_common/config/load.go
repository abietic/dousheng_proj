// 这个包是封装了yaml读取库viper的配置读取库
package config

import (
	"github.com/spf13/viper"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// 应用各种组件的配置结构体
var Configs struct {
	MysqlConfig  MysqlConfigStruct
	MogoDBConfig MogoDBConfigStruct
	EtcdConfig   EtcdConfigStruct
	RpcConfig    RpcConfigStruct
	JaegerConfig jaegercfg.Configuration
	HertzConfig  HertzConfigStruct
	MinioConfig  MinioConfigStruct
}

// 给出相对于应用启动目录的配置文件路径，加载配置文件
func LoadConfig(filePath string) {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&Configs); err != nil {
		panic(err)
	}
}
