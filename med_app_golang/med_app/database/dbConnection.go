package med_app

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

//DB connections
func LocalDbConnect() (*sql.DB, error) {
	log.Println("LocalDBConnect")
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "192.168.2.5", 3306, "guru")

	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Println("open connection failed: ", err)
		return db, err
	}
	log.Println("LocalDbConnect-")
	return db, err
}
