package migrate

import (
	"fmt"
	"log"
  
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //nolint
	_ "github.com/golang-migrate/migrate/v4/source/file"       //nolint
  )
  
  func Up(databaseURL, migrationsURL string) {
	migrator, err := migrate.New(migrationsURL, databaseURL)
	if err != nil {
	  log.Fatal("migrate.New", err)
	}
	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
	  log.Fatal("migrator.Up", err)
	}
  
	version, dirty, err := migrator.Version()
	if err != nil {
	  log.Fatal("migrator.Version", err)
	}
	fmt.Printf("Applied version %d, dirty %t\n", version, dirty)
  }

func CreateCleanDB(migrationsURL, databaseURL string) error {
	migrator, err := migrate.New(migrationsURL, databaseURL)
	if err != nil {
		return err
	}

	if err := migrator.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	version, dirty, err := migrator.Version()
	if err != nil {
		log.Fatal("migrate.Version: ", err)
	}
	fmt.Printf("Applied version %d, dirty %t\n", version, dirty)
	return nil
}
