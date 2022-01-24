package config

import "os"

type Config struct {
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbScheme    string
	IsDebugMode bool
	ServerHost  string
	ServerPort  string
}

func Read() *Config {
	cfg := &Config{}
	cfg.DbHost = os.Getenv("BIRTHDAY_MYSQL_HOST")
	cfg.DbPort = os.Getenv("BIRTHDAY_MYSQL_PORT")
	cfg.DbUser = os.Getenv("BIRTHDAY_MYSQL_DB_USER")
	cfg.DbScheme = os.Getenv("BIRTHDAY_MYSQL_DB_NAME")
	cfg.DbPassword = os.Getenv("BIRTHDAY_MYSQL_DB_PASS")
	cfg.ServerHost = os.Getenv("BIRTHDAY_SERVER_HOST")
	cfg.ServerPort = os.Getenv("BIRTHDAY_SERVER_PORT")
	cfg.IsDebugMode = true

	return cfg
}
