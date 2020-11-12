package helper

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // sql behavior modified
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "gituser"
	password = "passme"
	dbname   = "mydb"
)

// InitDB initializes database table
func InitDB() (*sql.DB, error) {
	var connectionString = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
