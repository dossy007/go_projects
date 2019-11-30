package handle

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"text/template"

	"../serv"
)

// type Vertex struct {
// 	Id   int
// 	Name string
// 	//servでtypeを使用していて、その情報を持ってきてるのでいらない
// }

func Showindex(w http.ResponseWriter, r *http.Request) {
	tem, _ := template.ParseFiles("index.html")

	p := serv.Connected()
	//serv packageのConnected funcでdbの情報を受け取っている
	if r.Method == "POST" { //sort time
		r.ParseForm()
		form := r.PostForm
		s, _ := strconv.Atoi(form["sort_id"][0])
		// fmt.Printf("%T\n", s)
		switch s {
		case 0:
			fmt.Println("0上から新しいよ")
			sort.Slice(p,
				func(i, j int) bool {
					return p[i].Updated_time.After(p[j].Updated_time)
				})
		case 1:
			fmt.Println("1上から古いよ")
			sort.Slice(p,
				func(i, j int) bool {
					return p[i].Updated_time.Before(p[j].Updated_time)
				})
		}
	}

	tem.Execute(w, p)
	//execute is template to act and http.RequestWriter に書き出す
}

func New(w http.ResponseWriter, r *http.Request) {

	tem, _ := template.ParseFiles("new.html")
	tem.Execute(w, "")
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		r.ParseForm() // Bodyデータを扱うには、事前にパースを行う

		// Formデータを取得
		form := r.PostForm
		body := form["body"][0]
		image := form["image"][0] //form is map value is []string of sliceで来るから[0]で取得

		// fmt.Fprintf(w, "フォーム1：\n%v\n", form["say"])
		//         wに描く format string  書き込む内容
		serv.Posts(body, image) //insert to db

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	tem, _ := template.ParseFiles("edit.html")
	params := r.URL.Query()       //Queryで取得 map
	e := serv.Edit(Getid(params)) //return []serv.Vertex
	tem.Execute(w, e[0])
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.PostFormValue("_method") == "PUT" {
		r.ParseForm()

		form := r.PostForm
		params := r.URL.Query()
		id := Getid(params)
		body := form["body"][0]
		image := form["image"][0]

		serv.Update(id, body, image)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		params := r.URL.Query()    //to use Getid
		serv.Delete(Getid(params)) //use sql delete func

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func Getid(params url.Values) int { //requestからidを取得
	var num string
	for k, _ := range params { //mapのkeyをget
		num = k
	}

	i, _ := strconv.Atoi(num) //string to int
	return i
}
