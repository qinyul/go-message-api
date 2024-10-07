package configuration

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"strconv"

	"github.com/qinyul/messaging-api/models"
	"github.com/uptrace/bun/driver/pgdriver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host         string
	User         string
	Password     string
	Port         int
	SSLMode      bool
	DatabaseName string
	DB           *gorm.DB
}

func NewDatabaseConfig(cfg *Config) *DatabaseConfig {
	dbPort, err := strconv.Atoi(cfg.DB_PORT)

	if err != nil {
		log.Fatal("NewDatabaseConfig:: failed to convert db port to int")
	}
	return &DatabaseConfig{
		Host:         cfg.DB_HOST,
		User:         cfg.DB_USER,
		Password:     cfg.DB_PASSWORD,
		DatabaseName: cfg.DATABASE_NAME,
		Port:         dbPort,
		SSLMode:      cfg.SSL_MODE == "true",
	}
}

func (cfg *DatabaseConfig) ConnectDatabase() error {
	if cfg.DB != nil {
		slog.Info("DB already initialized")
		return nil
	}

	dataSource := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, "postgres")

	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dataSource)))

	defer db.Close()

	if _, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", cfg.DatabaseName)); err != nil {
		if err.Error() != fmt.Sprintf("ERROR: database \"%s\" already exists (SQLSTATE=42P04)", cfg.DatabaseName) {
			log.Fatal("Failed to create database: ", err)
		}
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", cfg.Host, cfg.Port, cfg.User, cfg.DatabaseName, cfg.Password)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	cfg.DB = DB

	slog.Info("Database connection established")
	return nil
}

func (cfg *DatabaseConfig) Migrate() {
	if cfg.DB == nil {
		log.Fatal("DB not initialized")
	}

	cfg.DB.AutoMigrate(&models.Message{})
}
