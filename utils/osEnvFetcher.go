package utils

import (
	"log"
	"os"
)

const (
	PRODUCTION string = "production"
)

// EnvFetcher environment variables fetcher
type EnvFetcher interface {
	Getenv(key string) string
}

type DotenvLoader func(filenames ...string) (err error)

type osEnvFetcher struct {
	dotenvLoader DotenvLoader
}

func NewOsEnvFetcher(dotenvLoader DotenvLoader) *osEnvFetcher {
	if os.Getenv("APP_ENV") != PRODUCTION {
		err := dotenvLoader()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	return &osEnvFetcher{}
}

func (f *osEnvFetcher) Getenv(key string) string {
	return os.Getenv(key)
}
