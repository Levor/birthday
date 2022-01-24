package db

import (
	"fmt"
	"github.com/Levor/birthday/internal/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CreateConnection creates a new db connection
func NewConnection(cfg *config.Config) (*gorm.DB, error) {

	// initialize a new db connection
	connection, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=True", // username:password@protocol(host)/dbname?param=value
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbScheme))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if cfg.IsDebugMode {
		connection.LogMode(true)
	}
	return connection, nil
}
