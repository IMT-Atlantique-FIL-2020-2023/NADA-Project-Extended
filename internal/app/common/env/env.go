package env

import (
  viper "github.com/spf13/viper"
  "fmt"
)

func GetEnv(key string) string {
	return viper.GetString(key)
  }

func Init(path string){
	viper.AddConfigPath("internal/app/nada-transform/env")
    viper.SetConfigName(".nada-transform.env")
    viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	 }
}