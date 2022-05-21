package common

import (
	"os"

	"github.com/spf13/viper"
)

func LoadFromFile(v *viper.Viper) {
	configPath, err := os.Getwd()
	if err != nil {
		return
	}

	v.AddConfigPath(configPath)
	v.AddConfigPath(configPath + "/configs")

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	_ = v.ReadInConfig()

}
