package educert

import (
	"fmt"

	"github.com/emmanuelbenson/gpi-validate-v2/database"
)

/**
dao is Data Access Object - Like a repository
*/

// Create is a dao func adds new record of Educational Certificate check request
func Create(ec *EducationalCertificate) []error {

	db := database.Connect()
	defer db.Close()
	rs := db.Create(&ec).GetErrors()

	if len(rs) > 0 {
		fmt.Println(rs[0])
		return rs
	}

	return nil

	// db := databases.Connect()

	// stmt, err := db.Prepare("INSERT INTO educational_certificate VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")

	// if err != nil {

	// 	db.Close()
	// 	return err
	// }

	// _, err = stmt.Exec(0,
	// 	ec.UserID,
	// 	ec.FirstName,
	// 	ec.OtherName,
	// 	ec.LastName,
	// 	ec.Title,
	// 	ec.Type,
	// 	"NOTSTARTED",
	// 	ec.Verified,
	// 	ec.Document,
	// 	time.Now(),
	// 	"0000-00-00 00:00:00",
	// )

	// defer stmt.Close()
	// defer db.Close()

	// if err != nil {
	// 	return err
	// }

	// return nil
}

// FetchAll lists all user Educational Certificate
func FetchAll(userID int) ([]*EducationalCertificate, error) {
	ecs := []*EducationalCertificate{}

	db := database.Connect()

	defer db.Close()

	errs := db.Find(&ecs).GetErrors()

	if len(errs) > 0 {
		return ecs, errs[0]
	}

	return ecs, nil

	// rows, err := db.Query("SELECT id, first_name, other_name, last_name, title, type, status FROM educational_certificate WHERE user_id = ? ORDER BY created_at DESC", userID)
	// if err != nil {
	// 	db.Close()
	// 	rows.Close()
	// 	return nil, err
	// }

	// for rows.Next() {
	// 	ec := EducationalCertificate{}
	// 	if err := rows.Scan(
	// 		&ec.ID,
	// 		&ec.FirstName,
	// 		&ec.OtherName,
	// 		&ec.LastName,
	// 		&ec.Title,
	// 		&ec.Type,
	// 		&ec.Status,
	// 	); err != nil {
	// 		log.Println("Error @ educert.dao line 61", err.Error())
	// 		return nil, err
	// 	}

	// 	ecs = append(ecs, ec)
	// }
	// return ecs, nil
}
