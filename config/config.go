package config

import (
	"gorilla/utils"

	"github.com/spf13/viper"
)

type ViperConfig struct {
	*viper.Viper
}

var Gorilla *ViperConfig
var Logger *utils.Logger

func init() {
	Gorilla = initViperConfig()
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
