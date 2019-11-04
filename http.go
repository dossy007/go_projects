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
    r.ParseForm()       //urlが渡すオプションを解析します。POSTに対してはレスポンスパケットのボディを解析します（request body）
    //注意：もしParseFormメソッドがコールされなければ、以下でフォームのデータを取得することができません。
    fmt.Println(r.Form) //これらのデータはサーバのプリント情報に出力されます
    // fmt.Println("path", r.URL.Path)
    // fmt.Println("scheme", r.URL.Scheme)
    // fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    // fmt.Fprintf(w, "Hello astaxie!") //ここでwに書き込まれたものがクライアントに出力されます。
    t,_:= template.ParseFiles("index.html")
    // template.ParseFile関数は、edit.htmlの内容を読み込み、*template.Templateを返します。

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
    // t.Executeメソッドはテンプレートを実行し、生成したHTMLをhttp.ResponseWriterに書き出します。 ドット付きの識別子である.Titleと.Bodyは、それぞれp.Titleとp.Bodyを参照します。
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //リクエストを取得するメソッド
    if r.Method == "GET" {
        //rはhttprequestでmethodはgetの時にtemplateを読む
        t, _ := template.ParseFiles("login.gtpl.html")
        //同じfolderのhtmlをここで探す
        t.Execute(w, nil)
    } else {
        //必ず明示的にr.ParseForm()をコールした後でなければ、このフォームのデータに対して操作を行うことはできません
        r.ParseForm()

        //ログインデータがリクエストされ、ログインのロジック判断が実行されます。terminalにformの中身が表示
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
        fmt.Println("aaa",r.Form)
        //r.Formには全てのリクエストのデータが入っている
        // map[pass:111,name:111]]的な感じ
    }
}

func main() {
    db, err := sql.Open("mysql","root:@/mydb")
    if err != nil {
        log.Fatal("エラーでっせ")
    }
    defer db.Close()


// --------------------

    // var (
    // id int
    // body string
    // )
    // rows, err := db.Query("SELECT id, body FROM posts WHERE id = ?", 2)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer rows.Close()
    // for rows.Next() {
    //     err := rows.Scan(&id, &body)
    //     if err != nil {
    //         log.Fatal(err)
    //     }
    // }
    // log.Println(id,body)

    // err = rows.Err()
    // if err != nil {
    //     log.Fatal(err)
    // }
    //ここはmysqlから情報を引っ張る処理(引っ張ってその後どう処理するのかまではわからんかった)


    // Insert文発行
    // Resultが戻される
    // result, err := db.Exec(`
        // INSERT INTO posts(id, body) VALUES(5, 'Marketing')
  // `)

  //update文発行
  // result ,err := db.Exec(`
  //   UPDATE
  //   posts
  //   SET
  //   body = 898989
  //   WHERE
  //   id = '1'
  //   `)
  type Post struct { //post構造体を定義
    Id int
    Body string
    Image string
  }


  // var post Post

  //select文発行
  // result, err := db.Query(`
  //   SELECT
  //   id
  //   ,body
  //   FROM
  //   posts
  // `)

    //QueryRowで１行取得
    // id := 100
    // var body string

    // if err := db.QueryRow("SELECT body FROM posts WHERE id = 2").Scan(&body); err != nil {
    //     log.Fatal(err)
    // }

    // fmt.Println(id, body)

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
    // fmt.Println(hai)
}

  // for result.Next() {
  //       var dept Dept
  //       err := result.Scan(&(dept.Id), &(dept.Body))
  //       if err != nil {
  //           panic(err.Error())
  //       }
  //       fmt.Println(dept)
  //   }
  //posts_tableにid 4 body marketingをinsertするsqlをresultに入れる
    if err != nil {
        panic(err.Error())
    }
    // fmt.Println(id)

    // AutoIncrementの型で使える
    // 最後に挿入したキーを返す(が, 今回は主キーをAutoIncrementにしていないので使えない.例を誤った感)
    // id, err := result.LastInsertId()
    // if err != nil {
    //     panic(err.Error())
    // }
    // fmt.Println(id)
    //insertしたidを出力

    // 影響を与えた行数を返す
    // n, err := result.RowsAffected()
    // if err != nil {
    //     panic(err.Error())
    // }
    // fmt.Println(n)


// --------------------


    http.HandleFunc("/", sayhelloName)       //アクセスのルーティングを設定します
    http.HandleFunc("/login", login)         //アクセスのルーティングを設定します /loginのroutingが走ればlogin funcが動く
    error := http.ListenAndServe(":3000", nil) //監視するポートを設定します
    if error != nil {
        log.Fatal("ListenAndServe: ", error)
    }
}










