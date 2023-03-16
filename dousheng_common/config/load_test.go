package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	LoadConfig("./config.yaml")
	configs := Configs
	assert.NotEqual(t, configs.MysqlConfig.URL, "")
	assert.NotEqual(t, configs.MogoDBConfig.URI, "")
	assert.NotEqual(t, len(configs.EtcdConfig.Endpoints), 0)
	assert.NotEqual(t, configs.EtcdConfig.Endpoints[0], "")

}
