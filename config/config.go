package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type AppConfig struct {
	Server ServerConfig `yaml:"server" env-prefix:"SERVER_"`
	DB     DBConfig     `yaml:"db" env-prefix:"DB_"`
	Logger *zap.SugaredLogger
}

type ServerConfig struct {
	Port int    `yaml:"port" env:"PORT" env-default:"3000"`
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
}

type DBConfig struct {
	URI string `yaml:"uri" env:"URI" env-required:"true"`
}

func LoadConfig(configPath string) (*AppConfig, error) {
	var cfg AppConfig

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	cfg.Logger = zap.Must(zap.NewDevelopment()).Sugar()
	return &cfg, nil
}
