package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type LoggerConfig struct {
	Level  string `envconfig:"LEVEL" default:"DEBUG"`
	Folder string `envconfig:"FOLDER" default:"./logs"`
}

func newLoggerConfig() (LoggerConfig, error) {
	var conf LoggerConfig
	if err := envconfig.Process("LOGGER", &conf); err != nil {
		return LoggerConfig{}, fmt.Errorf("error read logger config: %w", err)
	}

	return conf, nil
}

func NewLoggerConfigMust() LoggerConfig {
	conf, err := newLoggerConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
