package db

import (
	"fmt"
	"time"

	"github.com/aralim11/go-crm-api/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetConnectionString(cfg config.DatabaseConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_general_ci",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.DBPort,
		cfg.DBName,
	)
}

func NewConnection(dbCOnfig config.DatabaseConfig) (*sqlx.DB, error) {
	connStr := GetConnectionString(dbCOnfig)

	db, err := sqlx.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Connection pool settings (VERY important in production)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
