package db

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDB(dbConn *sql.DB, dir string, dbString string) error {
	driver, err := mysql.WithInstance(dbConn, &mysql.Config{
		DatabaseName: dbString,
	})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		dir,
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	if err := migrator.Up(); err != nil && err.Error() != "no change" {
		return err
	}

	fmt.Println("Database migration completed successfully.")
	return nil
}
