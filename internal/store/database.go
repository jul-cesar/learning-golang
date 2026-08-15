package store

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost user=myuser password=mypassword dbname=mydatabase")
	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}
	fmt.Println("database connection established")
	return db, nil
}