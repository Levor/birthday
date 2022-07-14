package config

import (
	"os"
	"strings"
)

type Cors struct {
	Methods []string
	Origins []string
	Headers []string
}

type Config struct {
	DbHost        string
	DbPort        string
	DbUser        string
	DbPassword    string
	DbScheme      string
	IsDebugMode   bool
	ServerHost    string
	ServerPort    string
	Cors          *Cors
	SecretKeyPath string
	PublicKeyPath string
}

func Read() *Config {
	cfg := &Config{}
	cfg.DbHost = getEnv("BIRTHDAY_MYSQL_HOST", "localhost")
	cfg.DbPort = getEnv("BIRTHDAY_MYSQL_PORT", "3306")
	cfg.DbUser = getEnv("BIRTHDAY_MYSQL_DB_USER", "root")
	cfg.DbScheme = getEnv("BIRTHDAY_MYSQL_DB_NAME", "gres_wp")
	cfg.DbPassword = getEnv("BIRTHDAY_MYSQL_DB_PASS", "secret")
	cfg.ServerHost = getEnv("BIRTHDAY_SERVER_HOST", "localhost")
	cfg.ServerPort = getEnv("BIRTHDAY_SERVER_PORT", "8081")
	cfg.IsDebugMode = true
	cfg.SecretKeyPath = getEnv("BIRTHDAY_SECRET_KEY_PATH", "./keys/jwt.pem")
	cfg.PublicKeyPath = getEnv("BIRTHDAY_PUBLIC_KEY_PATH", "./keys/jwt.pub")
	cfg.Cors = ReadCorsConfig()
	return cfg
}

func ReadCorsConfig() *Cors {
	cors := Cors{
		Methods: strings.Split(getEnv("BIRTHDAY_CORS_METHODS", "GET,POST,PUT,PATCH,DELETE,OPTIONS"), ","),
		Origins: strings.Split(getEnv("BIRTHDAY_CORS_ORIGINS", "*"), ","),
		Headers: strings.Split(getEnv("BIRTHDAY_CORS_HEADERS", "*"), ","),
	}
	return &cors
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
