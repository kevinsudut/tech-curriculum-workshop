package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kevinsudut/tech-curriculum-workshops/config"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		config.DB_DRIVER,
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_DBNAME),
	)
	if err != nil {
		return db, errors.AddTrace(err)
	}

	err = db.Ping()
	if err != nil {
		return db, errors.AddTrace(err)
	}

	return db, nil
}
