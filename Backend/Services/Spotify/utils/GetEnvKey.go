package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func GetEnvKey(key string) string {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error on oppeing .env")
		return ""
	}

	return viper.Get(key).(string)
}
