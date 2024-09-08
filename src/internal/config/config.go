package config

import (
	"github.com/joho/godotenv"
	"os"
)

func InitializeConfig() {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	err = godotenv.Load(dir + "/.env")

	if err != nil {
		panic(err)
	}
}

func GetPort() string {
	return os.Getenv("PORT")
}
