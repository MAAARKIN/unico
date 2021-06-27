package db

import (
	"fmt"
	"log"

	"github.com/MAAARKIN/unico/config"
	"github.com/gobuffalo/packr/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func StartDatabase(cfg config.Config) *sqlx.DB {

	dbCfg := cfg.Database

	uri := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.Dbname)
	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
		log.Fatalln(err)
	}

	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./migrations"),
	}

	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)

	return db
}
