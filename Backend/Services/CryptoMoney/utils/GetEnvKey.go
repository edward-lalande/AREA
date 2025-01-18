package utils

import (
	"github.com/spf13/viper"
)

func GetEnvKey(key string) string {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return ""
	}

	return viper.Get(key).(string)
}
