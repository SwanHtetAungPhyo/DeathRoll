package repo

import (
	"DeathRoll/internal/config"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	// Avoid for the global db instance
	instance *gorm.DB
	// Ensure the db connection is initialized once
	once sync.Once
)

func Init(log *logrus.Logger, cfg *config.Configuration) {
	once.Do(func() {
		log.Info("Initializing database connection...")
		dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
			cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbUser, cfg.DbPass, "false")
		var err error
		// Retry logic for establishing the connection
		for i := 0; i < 5; i++ {
			instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				break
			}
			log.Warnf("Connection attempt %d failed: %v", i+1, err)
			time.Sleep(5 * time.Second)
		}

		// If the connection is not established after retries, log a fatal error
		if err != nil {
			log.Fatal("Database connection failed after retries: ", err)
		}

		sqlDB, err := instance.DB()
		if err != nil {
			log.Error("Failed to retrieve SQL DB instance: ", err)
			return
		}
		// Connection Pooling Set up
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(5 * time.Minute)
		log.Info("Database connection initialized successfully.")
	})
}

// GetDBInstance retrieves the singleton instance of the database connection.
// This allows the application to perform database operations by obtaining a reference to the initialized instance.
func GetDBInstance() *gorm.DB {
	return instance
}
