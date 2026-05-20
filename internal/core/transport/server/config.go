package core_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	Addr     int           `envconfig:"ADDR"`
	Timeout  time.Duration `envconfig:"TIMEOUT"`
	TimeZone string        `envconfig:"TIME_ZONE"`
}

func getServerConfig() (ServerConfig, error) {
	var conf ServerConfig
	if err := envconfig.Process("SERVER", &conf); err != nil {
		return ServerConfig{}, fmt.Errorf("failed reading server config: %w", err)
	}

	return conf, nil
}

func GetServerConfigMust() ServerConfig {
	conf, err := getServerConfig()
	if err != nil {
		panic(err)
	}

	return conf
}
