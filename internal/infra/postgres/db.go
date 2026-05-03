package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/hassan-alidoost/oms-project/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg config.DatabaseConfig) (*gorm.DB, func(), error) {
	connection := getConnectionString(cfg)
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
        return nil, nil, err
    }

	if err := sqlDb.Ping(); err != nil {
        return nil, nil, fmt.Errorf("database unreachable: %w", err)
    }

	sqlDb.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.ConnMaxLifeTime * time.Minute)

	log.Println("postgres db connection established.")

	cleanup := func() {
        sqlDb.Close()
    }

	return db, cleanup, nil
}

func getConnectionString(cfg config.DatabaseConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)
}