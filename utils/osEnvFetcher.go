package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvFetcher interface {
	Getenv(key string) string
}

type osEnvFetcher struct{}

func (f *osEnvFetcher) NewOsEnvFetcher() *osEnvFetcher {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &osEnvFetcher{}
}

func (f *osEnvFetcher) Getenv(key string) string {
	return os.Getenv(key)
}
