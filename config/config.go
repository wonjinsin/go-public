package config

import (
	"giraffe/utils"

	"github.com/spf13/viper"
)

type ViperConfig struct {
	*viper.Viper
}

var Giraffe *ViperConfig
var Logger *utils.Logger

func init() {
	Giraffe = initViperConfig()
}

func initViperConfig() *ViperConfig {
	v := viper.New()
	v.SetConfigName("local")
	viper.SetConfigType("json")
	v.AddConfigPath("./config/")
	v.AddConfigPath("../config/")

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		Logger.Logging().Warnw("fatal error config file: default", "err", err)
	}

	return &ViperConfig{
		Viper: v,
	}
}
