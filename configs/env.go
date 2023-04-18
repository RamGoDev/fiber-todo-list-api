package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err.Error())
	}

	return os.Getenv(key)
}

func GetEnvInt(key string) int {
	data, err := strconv.Atoi(GetEnv(key))

	if err != nil {
		log.Fatal(err.Error())
	}

	return data
}
