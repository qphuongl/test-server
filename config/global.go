package config

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/viper"
)

type ENVConfig struct {
	Port    int    `mapstructure:"PORT"`
	Message string `mapstructure:"MESSAGE"`
}

var EnvConfig ENVConfig

func Init() error {
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