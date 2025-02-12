package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadConfig(path string) {
	env := os.Getenv("ENV_ENVIRONMENT")

	if "" == env {
		err := godotenv.Load(path + "/env.local")
		if err != nil {
			fmt.Println("Env file doesn't exist")
			os.Exit(3)
		}
	}
}
