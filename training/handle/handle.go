package handle
import (
    "fmt"
    "net/http"
    "strings"
    "html/template"
)

//外部fileにexportするときには、Sayのように頭文字は大文字にすること
func SayhelloName(w http.ResponseWriter, r *http.Request) {
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
