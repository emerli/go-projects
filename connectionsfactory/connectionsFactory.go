package connectionsfactory

import (
	"fmt"

	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	MaxOpenConnections = 40
	MaxIdleConnections = 10
	DSNFormat          = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
)

var (
	usernameEnv   = "postgres"
	passwordEnv   = "postgres"
	dbNameEnv     = "test"
	host          = "localhost"
	port          = 5432
	enableLogs    = true
)

func NewDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(DSNFormat, host, usernameEnv, passwordEnv, dbNameEnv, port)

	config := gorm.Config{}
	if enableLogs {
		config = gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	}

	config.NamingStrategy = schema.NamingStrategy{
		SingularTable: true,
	}

	newDB, err := gorm.Open(postgres.Open(dsn), &config)
	if err != nil {
		//log.Fatalf("exceptionsError connecting to database [%s]: %s\n", os.Getenv("DB_NAME"), err.Error())
		return &gorm.DB{}, err
	}
	dbConfig, _ := newDB.DB()

	dbConfig.SetMaxIdleConns(MaxIdleConnections)
	dbConfig.SetMaxOpenConns(MaxOpenConnections)
	dbConfig.SetConnMaxLifetime(time.Hour)
	return newDB, nil
}

func init() {
}
