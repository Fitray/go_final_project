package core_server

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr string `envconfig:"TODO_PORT" required="true"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("failed to get config: %w", err)
	}
	return config, nil
}
