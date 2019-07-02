package db

import (
	_ "database/sql/driver"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"gogo/log"
)

var DB *sqlx.DB

func init() {
	var err error
	DB, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}
}
