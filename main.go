package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id          string
	Name        string
	Email       string
	Password    string
	Address     string
	PhoneNumber string
}

func main() {
	var db *sql.DB
	var err error
	// <username>:<password>@tcp(<hostname>:<port-db>)/<db-name>
	// var connectionString = "root:qwerty123@tcp(127.0.0.1:3306)/db_loanee"
	var connectionString = os.Getenv("CONNECTION_DB")
	fmt.Println("connectionstring:", connectionString)
	db, err = sql.Open("mysql", connectionString)
	if err != nil { // jika terjadi error
		log.Fatal("error open connection to db ", err.Error())
	}
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db ", errPing.Error())
	} else {
		fmt.Println("success connect to db")
	}
	//close db conn
	defer db.Close()

	//MENU
	fmt.Println("Pilih Menu:\n1. Read Data.\n2. Insert Data")
	var pilihan int
	fmt.Println("input pilihan anda:")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		// READ DATA --> SELECT
		// proses menjalankan query select
		rows, errSelect := db.Query("select id, name, email, password, address from users")

		if errSelect != nil { // ketika terjadi error saat menjalankan select
			log.Fatal("error run query select ", errSelect.Error())
		}
		// variabel untuk menyimpan semua data yang dibaca di db.Query
		var allUsers []User
		// membaca per baris
		for rows.Next() {
			var dataUserRow User // variabel untuk menyimpan data per baris
			errScan := rows.Scan(&dataUserRow.Id, &dataUserRow.Name, &dataUserRow.Email, &dataUserRow.Password, &dataUserRow.Address)
			if errScan != nil {
				log.Fatal("error scan select", errScan.Error())
			}
			// menambahkan tiap baris data yang dibaca ke slice
			allUsers = append(allUsers, dataUserRow)
		}
		// fmt.Println("data all users:\n", allUsers)
		for _, v := range allUsers {
			fmt.Println("nama:", v.Name)
		}
	case 2:
		fmt.Println("insert data")

	case 3:
		fmt.Println("Update data")
		// write your code here

	case 4:
		fmt.Println("read data user by id")
		// input: 1000 OR 1=1; DELETE from users;
		// select id, name, email from users where id = 1000 OR 1=1
		// write your code here

	case 5:
		fmt.Println("delete data")
		// write your code here

	} // EOF switch
}
