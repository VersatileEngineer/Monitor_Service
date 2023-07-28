package database

import "database/sql"

type SQLiteDB struct {
	Path string
	DB   *sql.DB
}

func NewSQLiteDB(path string) *SQLiteDB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}

	return &SQLiteDB{
		Path: path,
		DB:   db,
	}
}

func (db *SQLiteDB) Close() {
	db.DB.Close()
}
