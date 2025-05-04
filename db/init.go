package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func DB() (*sql.DB, error) {
	dbpath := os.Getenv("DBPATH")
	if dbpath == "" {
		return nil, fmt.Errorf("DBPATH environment variable not set")
	}

	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDB() error {
	db, dberr := DB()
	if dberr != nil {
		return fmt.Errorf("DB_CONNECTION_FAIL: %w", dberr)
	}

	if err := createQuestionsTables(db); err != nil {
		return fmt.Errorf("CREATE_QUESTIONS_TABLE_FAIL: %w", err)
	}

	if err := createUsersTables(db); err != nil {
		return fmt.Errorf("CREATE_USERS_TABLE_FAIL: %w", err)
	}

	if err := createLangsTables(db); err != nil {
		return fmt.Errorf("CREATE_LANGS_TABLE_FAIL: %w", err)
	}

	if err := createAnswersTables(db); err != nil {
		return fmt.Errorf("CREATE_ANSWERS_TABLE_FAIL: %w", err)
	}

	return nil
}

func createQuestionsTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS questions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userid INTEGER,
		title TEXT,
		description TEXT,
		lang TEXT,
		difficulty TEXT,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func createUsersTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		email TEXT,
		password TEXT,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func createLangsTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS langs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		type TEXT,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func createAnswersTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS answers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		qid INTEGER,
		answer TEXT,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
