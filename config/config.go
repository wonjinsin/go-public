package config

import (
	"gorilla/utils"
	"os"

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

	err := v.ReadInConfig()
	if err != nil {
		Logger.Logging().Errorw("fatal error config file: default")
		os.Exit(1)
	}

	return &ViperConfig{
		Viper: v,
	}
}
