package config

import (
	"context"
	"fmt"
	"log"
	"mangojek-backend/exception"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresConfig struct {
	ConnConfig *pgx.ConnConfig
}

func NewPostgresDatabase(configuration Config) *gorm.DB {
	postgresPoolMin, err := strconv.Atoi(configuration.Get("POSTGRES_POOL_MIN"))
	exception.PanicIfNeeded(err)
	postgresPoolMax, err := strconv.Atoi(configuration.Get("POSTGRES_POOL_MAX"))
	exception.PanicIfNeeded(err)
	postgresMaxIdleTime, err := strconv.Atoi(configuration.Get("POSTGRES_MAX_IDLE_TIME_SECOND"))
	exception.PanicIfNeeded(err)
	postgresDBName := configuration.Get("POSTGRES_DBNAME")
	postgresPort := configuration.Get("POSTGRES_PORT")
	postgresPass := configuration.Get("POSTGRES_PASS")
	postgresHost := configuration.Get("POSTGRES_HOST")
	postgresUser := configuration.Get("POSTGRES_USER")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", postgresHost, postgresUser, postgresPass, postgresDBName, postgresPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// disable log mode
		// Logger: logger.Default.LogMode(logger.Silent),

		// skip default transaction
		SkipDefaultTransaction: true,
	})
	exception.PanicIfNeeded(err)
	sqlDB, err := db.DB()

	sqlDB.SetMaxOpenConns(postgresPoolMax)
	sqlDB.SetConnMaxIdleTime(time.Duration(postgresMaxIdleTime) * time.Second)
	sqlDB.SetMaxIdleConns(postgresPoolMin)
	sqlDB.SetConnMaxLifetime(time.Duration(postgresMaxIdleTime) * time.Second)

	return db
}

func NewRunMigration(db *gorm.DB) {
	for _, entity := range RegisterEntities() {
		err := db.AutoMigrate(entity.Entity)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func NewPostgresContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
