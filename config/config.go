package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ViperConfig struct {
	*viper.Viper
}

var Gorilla *ViperConfig

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
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	return &ViperConfig{
		Viper: v,
	}
}
