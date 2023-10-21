package config

import (
	"log"
	"os"
)

type Configuration struct {
	DatabaseName      string
	DatabaseHost      string
	DatabaseUser      string
	DatabasePassword  string
	MigrateToVersion  string
	MigrationLocation string
	RedisPassword     string
	RedisDatabase     string
	RedisHost         string
}

func GetConfiguration() Configuration {
	return Configuration{
		DatabaseName:      getOrFail("DB_NAME"),
		DatabaseHost:      getOrFail("DB_HOST"),
		DatabaseUser:      getOrFail("DB_USER"),
		DatabasePassword:  getOrFail("DB_PASSWORD"),
		MigrateToVersion:  getOrDefault("MIGRATE", "latest"),
		MigrationLocation: getOrDefault("MIGRATION_LOCATION", "migrations"),
		RedisPassword:     getOrFail("RD_PASSWORD"),
		RedisDatabase:     getOrFail("RD_BASE"),
		RedisHost:         getOrFail("RD_HOST"),
	}
}

func getOrFail(key string) string {
	env, set := os.LookupEnv(key)
	if !set || env == "" {
		log.Fatalf("%s env var is missing", key)
	}
	return env
}

func getOrDefault(key, defaultVal string) string {
	env, set := os.LookupEnv(key)
	if !set {
		return defaultVal
	}
	return env
}
