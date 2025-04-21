package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Env    string       `yaml:"env"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (c *Config) validate() error {
	if c.Env != EnvDevelopment && c.Env != EnvProduction {
		return fmt.Errorf("invalid env parameter: should be one of (%s, %s)", EnvDevelopment, EnvProduction)
	}

	return nil
}

func MustLoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("unable to load config: %w", err))
	}

	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var cfg Config
	viper.Unmarshal(&cfg)

	if err := cfg.validate(); err != nil {
		panic(err)
	}

	return &cfg
}
