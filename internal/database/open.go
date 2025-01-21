package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var dbPath = "../../internal/database/demo.db"

func Open() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", dbPath)

	return
}
