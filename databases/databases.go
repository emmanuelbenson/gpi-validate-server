package databases

import (
	"database/sql"
	"log"
)

const (
	driver = "mysql"
	dsn    = "root:root@/gpi_validate_v2?parseTime=true"
)

// Connect ...
func Connect() (db *sql.DB) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatalln("Error opening SQL: ", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("Error pinging:", err.Error())
	}
	return db
}
