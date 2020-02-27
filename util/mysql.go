package mysqlUtil

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func NewMysql(datasource string, f func(*sql.DB)) *error {
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	f(db)
	return &err
}

func BBB(str string, v interface{}){

}




