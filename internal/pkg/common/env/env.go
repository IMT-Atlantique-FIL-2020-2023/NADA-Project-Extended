package env

import (
	"fmt"

	viper "github.com/spf13/viper"
)

func GetEnv(key string) string {
	return viper.GetString(key)
}

func Init(path string, name string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}
