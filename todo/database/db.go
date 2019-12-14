package database

import (
	"database/sql"
	"log"
	"os"
	// "fmt"

	_ "github.com/go-sql-driver/mysql"
	//sql: unknown driver "mysql" (forgotten import?)と言われる
)
var db *sql.DB

func ConnectDB() *sql.DB {
	// db, err := sql.Open("mysql", "root:@/mydb?parseTime=true")
	var (
		connectionName = os.Getenv("CLOUDSQL_CONNECTION_NAME")
		user           = os.Getenv("CLOUDSQL_USER")
		name           = os.Getenv("CLOUDSQL_DATABASE_NAME")
		password       = os.Getenv("CLOUDSQL_PASSWORD") // NOTE: password may be empty
)

var err error
// db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName))

var dbopenstring = user + ":"+ password + "@unix(/cloudsql/"+connectionName+")/"+name
// user:password@unix(/cloudsql/INSTANCE_CONNECTION_NAME)/dbname

db,err = sql.Open("mysql",dbopenstring)
// db, err := sql.Open("mysql", "testuser:testpass@unix(/cloudsql/testproj:asia-northeast1:testinstance)/testdb")

if err != nil {
		log.Fatalf("Could not open db: %v", err)
}
	// cfg.DBName = "DBName"
	// cfg.ParseTime = true

	// const timeout = 10 * time.Second
	// cfg.Timeout = timeout
	// cfg.ReadTimeout = timeout
	// cfg.WriteTimeout = timeout

	// db, err := mysql.DialCfg(cfg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
