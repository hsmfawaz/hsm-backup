package utilities

import (
	"github.com/spf13/viper"
)

type Config struct {
	SSH_KEY_PATH string `mapstructure:"ssh_key_path"`
}

func NewConfig() *Config {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}
