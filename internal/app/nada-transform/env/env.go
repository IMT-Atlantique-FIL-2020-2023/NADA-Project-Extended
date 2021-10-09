package env

import (
  godotenv "github.com/joho/godotenv"
  "os"
  "log"
)

func GetEnv(key string) string {
	return os.Getenv(key)
  }

  func init(){
	err := godotenv.Load("internal/app/nada-transform/env/.env")
	if err != nil {
		log.Fatal(err)
	}
  }