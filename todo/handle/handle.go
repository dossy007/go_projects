package handle

import (
	"fmt"
	"net/http"
	"text/template"

	"../serv"
)

type Vertex struct {
	Id   int
	Name string
	//servをimportしているからvertex使えるのか？
}

func Showindex(w http.ResponseWriter, r *http.Request) {
	tem, _ := template.ParseFiles("index.html")

	p := serv.Posts()
	//serv packageのPosts funcでdbの情報を受け取っている
	// fmt.Println(r.Method)

	// var pp []Vertex //slice Vertex型を指定
	// v := Vertex{Id: 1, Name: "https://d1f5hsy4d47upe.cloudfront.net/3c/3cfb747bdc87c4dacd090210a0b95d5e_t.jpeg"}
	// v2 := Vertex{Id: 2, Name: "https://fotopus.com/neko/images/top/neko_subject_04.jpg"}
	// pp = append(pp, v)
	// fmt.Println(p[0])
	// v1 := Vertex{Id: 3, Name: "3ですよ〜〜"}
	// fmt.Println(p[0][0])
	// fmt.Printf("%T\n %#v", p[1], p[1])
	// fmt.Printf("%T\n", p)
	// fmt.Printf("%T\n", pp)
	// fmt.Printf("%T\n", v)
	// fmt.Printf("---------------")
	// fmt.Println(p)
	tem.Execute(w, p)
	//execute is template to act and http.RequestWriter に書き出す
}

func New(w http.ResponseWriter, r *http.Request) {

	tem, _ := template.ParseFiles("new.html")
	po := 1
	tem.Execute(w, po) //poは2つ引数いるみたい
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "POST" {
		fmt.Println("post ok")
	}
	// fmt.FPrintf(w, "クエリ:%s\n", r.URL.RawQuery)
	// fmt.Fprintf(w, "クエリ：%s\n", r.URL.RawQuery)
	// Bodyデータを扱う場合には、事前にパースを行う
	r.ParseForm()

	// Formデータを取得.
	form := r.PostForm
	forms := form["say"][0]
	fmt.Printf("%T\n", forms) //form is map value is []string of sliceで来るから[0]で取得

	// fmt.Fprintf(w, "フォーム1：\n%v\n", form["say"])
	//         wに描く format string  書き込む内容

	// Showindex(w http.ResponseWriter, r *http.Request)
	// tem, _ := template.ParseFiles("index.html")
	// po := 1
	// tem.Execute(w, po) //poは2つ引数いるみたい
}
