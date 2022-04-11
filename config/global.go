package config

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/viper"
)

type ENVConfig struct {
	Port           int    `mapstructure:"PORT"`
	Message        string `mapstructure:"MESSAGE"`
	SecretMessage  string `mapstructure:"SECRET_MESSAGE"`
	PrivateMessage string `mapstructure:"PRIVATE_MESSAGE"`
}

var EnvConfig ENVConfig

func Init() error {
	fmt.Println(os.Environ())
	_, path, _, _ := runtime.Caller(0)
	fmt.Println(path)
	fmt.Println(os.Getwd())
	viper.AddConfigPath(fmt.Sprintf("%s/..", path))
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	if err := viper.Unmarshal(&EnvConfig); err != nil {
		return err
	}

	return nil
}
