package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/yaza-putu/online-test-dot/src/config"
	"github.com/yaza-putu/online-test-dot/src/database"
	"github.com/yaza-putu/online-test-dot/src/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

func NewServerTest(method string, target string, params string) (echo.Context, *httptest.ResponseRecorder) {
	// setup env
	envTesting()

	// call database
	databaseTesting()

	e := echo.New()
	req := &http.Request{}
	rec := httptest.NewRecorder()

	switch method {
	case "POST":
		req = httptest.NewRequest(http.MethodPost, target, strings.NewReader(params))
		break
	case "PUT":
		req = httptest.NewRequest(http.MethodPut, target, strings.NewReader(params))
		break
	case "PATCH":
		req = httptest.NewRequest(http.MethodPatch, target, strings.NewReader(params))
		break
	case "DELETE":
		req = httptest.NewRequest(http.MethodPatch, target, strings.NewReader(params))
		break
	default:
		req = httptest.NewRequest(http.MethodGet, target, strings.NewReader(params))
		break
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return e.NewContext(req, rec), rec
}

func envTesting() {
	workingdir, _ := os.Getwd()
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(workingdir + "/../../../../")
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func databaseTesting() {
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
}
