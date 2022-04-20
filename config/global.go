package config

import (
	"fmt"
	"runtime"

	"github.com/spf13/viper"
)

type ENVConfig struct {
	Port           int    `mapstructure:"PORT"`
	Message        string `mapstructure:"MESSAGE"`
	PrivateMessage string `mapstructure:"PRIVATE_MESSAGE"`
	RedisHost      string `mapstructure:"REDIS_HOST"`
	RedisPass      string `mapstructure:"REDIS_PASS"`
	MongoHost      string `mapstructure:"MONGO_HOST"`
	MongoPass      string `mapstructure:"MONGO_PASS"`
}

var EnvConfig ENVConfig

func Init() error {
	// fmt.Println(os.Environ())
	_, path, _, _ := runtime.Caller(0)
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
