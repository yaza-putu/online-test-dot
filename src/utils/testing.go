package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/yaza-putu/online-test-dot/src/config"
	"github.com/yaza-putu/online-test-dot/src/database"
	"github.com/yaza-putu/online-test-dot/src/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func EnvTesting(path string) error {
	workingdir, _ := os.Getwd()
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(workingdir + path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.Set("app_debug", false)
	viper.Set("app_status", "test")

	return err
}

func DatabaseTesting() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB().User, config.DB().Password, config.DB().Host, config.DB().Port, config.DB().Name)

	sqlDB, err := sql.Open("mysql", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		if config.App().Debug == true {
			logger.New(err, logger.SetType(logger.FATAL))
		} else {
			logger.New(
				errors.New("Database connection error, please enable debug mode to view error"),
				logger.SetType(logger.FATAL),
			)
		}
	}

	database.Instance = db
	return err
}
