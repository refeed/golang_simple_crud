package lib

import (
	"log"
	"os"
)

func GetEnvOrDefault(envVar string, defaultValue string) string {
	value, set := os.LookupEnv(envVar)
	if !set {
		log.Printf("%v not set, using default val \"%v\"", envVar, defaultValue)
		return defaultValue
	}
	return value
}
