package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

// Make sure config file exist
func ConfigInitialization() error {
	configDirPath, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	appDir := filepath.Join(configDirPath, "restui")
	err = os.MkdirAll(appDir, 0o755)
	if err != nil {
		return err
	}

	dataPath := filepath.Join(appDir, "data.db")
	_, err = os.Stat(dataPath)
	firstTime := false
	if os.IsNotExist(err) {
		firstTime = true
	}

	db, err := sql.Open("sqlite3", dataPath)
	if err != nil {
		return err
	}
	defer db.Close()

	if firstTime {
		err = initializeDatabase(db)
		if err != nil {
			return err
		}
	}

	return err
}
func DatabaseInitialize() (*sql.DB, error) {
	configDirPath, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	appDir := filepath.Join(configDirPath, "restui")
	err = os.MkdirAll(appDir, 0o755)
	if err != nil {
		return nil, err
	}

	dataPath := filepath.Join(appDir, "data.db")

	db, err := sql.Open("sqlite3", dataPath)
	if err != nil {
		return nil, err
	}

	return db, err
}

func initializeDatabase(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS suggestions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    text TEXT NOT NULL UNIQUE,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
	`
	_, err := db.Exec(schema)
	return err
}
