package database

import (
	"fmt"
	"golang-kafka/util/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var (
	db *gorm.DB
)

type connection struct {
	Type     string
	Host     string
	Port     string
	DBName   string
	UserName string
	Password string
}

func getDatabaseConfig() connection {
	return connection{
		Type:     os.Getenv("MYSQL_DB_TYPE"),
		Host:     os.Getenv("MYSQL_DB_HOST"),
		Port:     os.Getenv("MYSQL_DB_PORT"),
		DBName:   os.Getenv("MYSQL_DB_NAME"),
		UserName: os.Getenv("MYSQL_DB_USER"),
		Password: os.Getenv("MYSQL_DB_PASSWORD"),
	}
}

func InitDatabase() {
	var err error

	config := getDatabaseConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.UserName, config.Password, config.Host, config.Port, config.DBName,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Errorf("Database connection failed: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("Failed to get sqlDB: %v", err)
	}

	sqlDB.SetMaxIdleConns(11)
	sqlDB.SetMaxOpenConns(33)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if db != nil {
		sqlDB, err := db.DB()
		if err == nil {
			_ = sqlDB.Close()
		}
	}
}
