package config

import (
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		HTTP  HTTPConfig  `mapstructure:"http"`
		Redis RedisConfig `mapstructure:"redis"`
		GRPC  GRPCConfig  `mapstructure:"grpc"`
	}
	HTTPConfig struct {
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegaBytes"`
	}
	RedisConfig struct {
		Host       string `mapstructure:"host"`
		Port       string `mapstructure:"port"`
		CacheLimit int    `mapstructure:"cacheLimit"`
	}
	GRPCConfig struct {
		Port string `mapstructure:"port"`
	}
)

func Init(configsDir, environment string) (*Config, error) {
	viper.AddConfigPath(configsDir)
	viper.SetConfigName(environment)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
