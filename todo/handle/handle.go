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

	p := serv.Connected()
	//serv packageのConnected funcでdbの情報を受け取っている
	// fmt.Println(r.Method)

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

	// Bodyデータを扱う場合には、事前にパースを行う
	r.ParseForm()

	// Formデータを取得
	form := r.PostForm
	body := form["body"][0]
	image := form["image"][0] //form is map value is []string of sliceで来るから[0]で取得

	// fmt.Fprintf(w, "フォーム1：\n%v\n", form["say"])
	//         wに描く format string  書き込む内容
	serv.Posts(body, image) //insert to db

	http.Redirect(w, r, "/", http.StatusMovedPermanently) //redierct to root
}
