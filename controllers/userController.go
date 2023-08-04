package controllers

import (
	"database/sql"
	"fakhry/go-sql/entities"
	"log"
)

func GetAllUserController(db *sql.DB) []entities.User {
	// READ DATA --> SELECT
	// proses menjalankan query select
	rows, errSelect := db.Query("select id, name, email, password, address from users")

	if errSelect != nil { // ketika terjadi error saat menjalankan select
		log.Fatal("error run query select ", errSelect.Error())
	}
	// variabel untuk menyimpan semua data yang dibaca di db.Query
	var allUsers []entities.User
	// membaca per baris
	for rows.Next() {
		var dataUserRow entities.User // variabel untuk menyimpan data per baris
		errScan := rows.Scan(&dataUserRow.Id, &dataUserRow.Name, &dataUserRow.Email, &dataUserRow.Password, &dataUserRow.Address)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		// menambahkan tiap baris data yang dibaca ke slice
		allUsers = append(allUsers, dataUserRow)
	}
	return allUsers
}

func AddUserController(db *sql.DB) {

}
