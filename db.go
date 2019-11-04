// package main
// import (
//     "database/sql"
//     "fmt"
//     _ "github.com/go-sql-driver/mysql"
//     _"log"
// )
// type Message struct {
//     Id   int
//     Content string
// }
// var Db *sql.DB
// func init() {
//     var err error
//     Db, err = sql.Open("mysql", "root:@/test_sample_development")
//     if err != nil {
//         panic(err)
//     }
// }
// func main() {
//     err := Create()
//     if err != nil {
//         log.Fatal("Create error!")
//     }
//     fmt.Println("Createできたよ!")
//     user, err := GetUser(1)
//     if err != nil {
//         log.Fatal("GetPost error!")
//     }
//     fmt.Print("ID1の値を取得したよ:")
//     fmt.Println(user)
// }
// func Create() (err error) {
//     _, err = Db.Exec("INSERT INTO <table> (name) VALUES ('nyaa')")
//     return
// }
// func GetUser(id int) (user User, err error) {
//     user = User{}
//     err = Db.QueryRow("SELECT id, name FROM <table> WHERE id = ?", id).Scan(&user.Id, &user.Name)
//     return
// }




package main
 
import (
    "database/sql"
    "fmt"
 
    _ "github.com/go-sql-driver/mysql"
)
 
func main() {
    db, err := sql.Open("mysql", "root@/test_sample_development")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
 
    rows, err := db.Query("SELECT * FROM messages")
    defer rows.Close()
    if err != nil {
        panic(err.Error())
    }
 
    for rows.Next() {
        var id int
        var body string
        if err := rows.Scan(&id, &body); err != nil {
            panic(err.Error())
        }
        fmt.Println(id, body)
    }
}
