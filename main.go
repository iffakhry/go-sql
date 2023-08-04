package main

import (
	"database/sql"
	"fakhry/go-sql/controllers"
	"fakhry/go-sql/entities"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

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
		result := controllers.GetAllUserController(db)
		for _, v := range result {
			fmt.Println("nama:", v.Name)
		}

	case 2:
		newUser := entities.User{}
		fmt.Println("Input new ID:")
		fmt.Scanln(&newUser.Id)
		fmt.Println("Input new Name:")
		fmt.Scanln(&newUser.Name)
		fmt.Println("Input new Email:")
		fmt.Scanln(&newUser.Email)
		fmt.Println("Input new Password:")
		fmt.Scanln(&newUser.Password)
		fmt.Println("Input new Address:")
		fmt.Scanln(&newUser.Address)
		fmt.Println("Input new Phone Number:")
		fmt.Scanln(&newUser.PhoneNumber)

		// result, errInsert := db.Exec("INSERT INTO users (id, name, email, password, address, phone_number) VALUES (?, ?, ?, ?,?,?)", newUser.Id, newUser.Name, newUser.Email, newUser.Password, newUser.Address, newUser.PhoneNumber)

		statement, errPrepare := db.Prepare("INSERT INTO users (id, name, email, password, address, phone_number) VALUES (?, ?, ?, ?,?,?)")
		if errPrepare != nil {
			log.Fatal("error prepare insert", errPrepare.Error())
		}
		result, errInsert := statement.Exec(newUser.Id, newUser.Name, newUser.Email, newUser.Password, newUser.Address, newUser.PhoneNumber)

		if errInsert != nil {
			log.Fatal("error insert", errInsert.Error())
		} else {
			row, _ := result.RowsAffected()
			if row > 0 {
				fmt.Println("success Insert data")
			} else {
				fmt.Println("failed to insert data")
			}
		}

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
