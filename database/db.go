package database

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	db *sql.DB
}

var (
	sqlite     *Sqlite
	initDBOnce sync.Once
)

const file string = "db.sqlite"
const create string = `
  CREATE TABLE "urls" (
	"short_url"	TEXT NOT NULL UNIQUE,
	"original_url"	TEXT NOT NULL UNIQUE,
	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY("original_url")
	);
	`

func newDB() {
	initDBOnce.Do(func() {
		db, err := sql.Open("sqlite3", file)
		if err != nil {
			panic(err)
		}
		if _, err := db.Exec(create); err != nil {
			panic(err)
		}

		sqlite = &Sqlite{
			db: db,
		}
	})
}

func (q *Sqlite) Close() {
	q.db.Close()
}

func GetDB() *Sqlite {
	if sqlite == nil {
		newDB()
	}

	return sqlite
}
