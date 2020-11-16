package helper

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "gituser"
	password = "passme123"
	dbname   = "mydb"
)


type Shipment struct {
	gorm.Model
	Packages []Package
	Data     string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Package struct {
	gorm.Model
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}
