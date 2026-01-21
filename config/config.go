package config

import "os"

type Config struct {
	AppPort   string
	MongoURI  string
	MongoDB   string
	JWTSecret string
}

func Load() *Config {
	return &Config{
		AppPort:   getEnv("APP_PORT", "8080"),
		MongoURI:  mustEnv("MONGO_URI"),
		MongoDB:   mustEnv("MONGO_DB"),
		JWTSecret: mustEnv("JWT_SECRET"),
	}
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(key + " is required")
	}
	return v
}
