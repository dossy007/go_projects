//このfileでserverは立てれている!!!

package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    "database/sql"
    _"github.com/go-sql-driver/mysql"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    t,_:= template.ParseFiles("index.html")

    type Mydata struct {
        Name string
        Age int
        Message string
    }

    //構造体を定義 構造体で渡さないとあかんねんて(集合体)
    pers := Mydata{
        Name: "dossy",
        Age: 100,
        Message: "変数を持って来ているよ",
    }

    t.Execute(w,pers)
}


func main() {
    db, err := sql.Open("mysql","root:@/mydb")
    if err != nil {
        log.Fatal("エラーでっせ")
    }
    defer db.Close()

  type Post struct { //post構造体を定義
    Id int
    Body string
    Image string
  }

    //Queryで複数のレコードを取得
    rows,err := db.Query(`
        SELECT id,body,image FROM posts
        `)
        if err != nil{
            log.Fatal(err)
        }
        defer rows.Close()

for rows.Next() { //for文
    var id int
    var name string
    var image string
    //空の変数を定義し、Scanで値をコピー
    //id ,body, imageの順番で取得しているので、その順番で入る
    if err := rows.Scan(&id, &name, &image); err != nil {
        log.Fatal(err)
    }
    fmt.Println(id, name,image)
    //ここでターミナルにdbの情報を表示させている
}

    if err != nil {
        panic(err.Error())
    }

    http.HandleFunc("/", sayhelloName)       //アクセスのルーティングを設定します
    error := http.ListenAndServe(":3000", nil) //監視するポートを設定します
    if error != nil {
        log.Fatal("ListenAndServe: ", error)
    }
}


