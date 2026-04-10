package db

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(dbConn *sql.DB, dir string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: dir,
	}

	// Step 1: DOWN (rollback all)
	_, err := migrate.Exec(dbConn, "mysql", migrations, migrate.Down)
	if err != nil {
		return fmt.Errorf("Down migration failed: %w", err)
	}

	// Step 2: UP (apply fresh)
	number, err := migrate.Exec(dbConn, "mysql", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("Up migration failed: %w", err)
	}

	fmt.Printf("DB migration completed. %d migrations applied.\n", number)
	return nil
}
