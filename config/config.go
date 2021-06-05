package config

import (
	"giraffe/utils"
	"path"
	"path/filepath"
	"runtime"

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
	v.AddConfigPath("../../config/")
	v.AddConfigPath("../../../config/")

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		Logger.Logging().Warnw("fatal error config file: default", "err", err)
	}

	if v.GetString("giraffe_env") == "local" {
		v.Set("absPath", getRootDir())
	}

	return &ViperConfig{
		Viper: v,
	}
}

func getRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
