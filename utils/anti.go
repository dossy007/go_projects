package utils

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectD() *sql.DB {
	db , err := sql.Open("mysql", "root:@/mydb")
	if err != nil {
		log.Fatal("エラーです")
	}	
	return db
}

//package antなのでantiの塊として使える？


