package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kelseyhightower/envconfig"
	"net/http"
)

type AppConfig struct {
	Port string `required:"true"`
}

type DbConfig struct {
	Host    string `required:"true"`
	Port    string `required:"true"`
	User    string `required:"true"`
	Pass    string `required:"true"`
	Name    string `required:"true"`
	Dialect string `required:"true"`
	SSLMode string `default:"disable"`
}

type Quote struct {
	Id     uint
	Quote  string
	Author string
}

func loadDbConfig() (*DbConfig, error) {
	config := &DbConfig{}
	err := envconfig.Process("DB", config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot load db config, %s", err))
	}

	return config, nil
}

func loadAppConfig() (*AppConfig, error) {
	config := &AppConfig{}
	err := envconfig.Process("APP", config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot load app config, %s", err))
	}

	return config, nil
}

func setupDb(dbConfig *DbConfig) (*gorm.DB, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Name, dbConfig.Pass, dbConfig.SSLMode)

	db, err := gorm.Open(dbConfig.Dialect, connString)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(Quote{})

	return db, nil
}

func setupRouter(db *gorm.DB) (*gin.Engine, error) {
	ginEngine := gin.Default()
	ginEngine.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	ginEngine.GET("/quote", func(c *gin.Context) {
		var quotes []Quote
		query := db.Find(&quotes)

		if query.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching rows"})
			return
		}

		c.JSON(http.StatusOK, quotes)
	})

	return ginEngine, nil
}

func main() {
	appConfig, err := loadAppConfig()
	if err != nil {
		panic(err)
	}

	dbConfig, err := loadDbConfig()
	if err != nil {
		panic(err)
	}

	db, err := setupDb(dbConfig)
	if err != nil {
		panic(err)
	}

	ginEngine, err := setupRouter(db)
	if err != nil {
		panic(err)
	}

	err = ginEngine.Run(fmt.Sprintf(":%s", appConfig.Port))

	if err != nil {
		panic(err)
	}
}
