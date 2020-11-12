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
