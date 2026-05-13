package core_db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

const (
	schema = `
	CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date CHAR(8) NOT NULL DEFAULT "",
		title varchar(100) NOT NULL DEFAULT "",
		comment TEXT,
		repeat VARCHAR(128) NOT NULL DEFAULT ""
	);
	CREATE INDEX IF NOT EXISTS idx_date ON scheduler (date)
	`
)

type Database struct {
	DB *sql.DB
}

func Init(dbFile string) (Database, error) {
	if dbFile == "" {
		return Database{}, fmt.Errorf("failed to get database path")
	}

	dir := filepath.Dir(dbFile)
	if dir != "." && dir != "/" {
		os.MkdirAll(dir, 0755)
	}

	_, err := os.Stat(dbFile)
	if err != nil {
		file, err := os.Create(dbFile)
		if err != nil {
			return Database{}, err
		}
		file.Close()
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		db.Close()
		return Database{}, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return Database{}, fmt.Errorf("failed to ping database: %w", err)
	}

	_, err = db.Exec(schema)
	if err != nil {
		db.Close()
		panic(err)
	}

	return Database{
		DB: db,
	}, nil
}

func (d *Database) Close() {
	d.DB.Close()
}
