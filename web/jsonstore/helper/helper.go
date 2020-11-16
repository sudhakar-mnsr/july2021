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

// GORM creates tables with plural names. Use this to suppress it
func (Shipment) TableName() string {
	return "Shipment"
}

func (Package) TableName() string {
	return "Package"
}

