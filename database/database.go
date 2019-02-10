package database

import (
	"github.com/jinzhu/gorm"
)

const (
	driver = "mysql"
	dsn    = "root:root@/gpi_validate_v2?parseTime=true"
)

// Connect ...
func Connect() *gorm.DB {
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		panic(err)
	}
	return db
}
