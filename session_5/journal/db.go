package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	//go:embed sql/add.sql
	addSQL string

	//go:embed sql/last.sql
	lastSQL string
)

type Entry struct {
	Time    time.Time `json:"time"`
	User    string    `json:"user"`
	Content string    `json:"content"`
}

type DB struct {
	db *sql.DB
}

func NewDB(dsn string) (*DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db := DB{conn}
	if err := db.Health(); err != nil {
		conn.Close()
		return nil, err
	}

	return &db, nil
}

func strDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func (d *DB) Add(entry Entry) error {
	sql := fmt.Sprintf(addSQL, strDate(entry.Time), entry.User, entry.Content)
	_, err := d.db.Exec(sql)
	return err
}

func (d *DB) Last() (Entry, error) {
	row := d.db.QueryRow(lastSQL)
	var e Entry
	if err := row.Scan(&e.Time, &e.User, &e.Content); err != nil {
		return Entry{}, err
	}

	return e, nil
}

func (d *DB) Health() error {
	return d.db.Ping()
}

func (d *DB) Close() error {
	return d.db.Close()
}
