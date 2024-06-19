package config

import (
	"strings"

	"github.com/spf13/viper"
)

var Single *Config

func init() {

	Single = &Config{v: getViper()}
}

func getViper() *viper.Viper {

	v := viper.New()

	v.SetConfigType("yaml")

	v.AddConfigPath("configs")

	v.SetConfigName("config")

	v.AutomaticEnv()
	v.SetEnvPrefix("engine")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := v.ReadInConfig()
	if err != nil {
		panic("error in reading config file: " + err.Error())
	}

	_ = v.AllSettings()

	return v
}
