package handle

import (
	"net/http"
	// "net/url"
	// "sort"
	// "strconv"
	"text/template"
	"github.com/dossy007/go_projects/todo/serv"


	// "./serv"
	// "../serv"
)

func Showindex(w http.ResponseWriter, r *http.Request) {
	tem, _ := template.ParseFiles("index.html")

	// p := serv.Connected()
	//serv packageのConnected funcでdbの情報を受けとる

	// if r.Method == "POST" { //sort time
	// 	r.ParseForm()
	// 	form := r.PostForm
	// 	s, _ := strconv.Atoi(form["sort_id"][0])
	// 	switch s {
	// 	case 0: //new
	// 		sort.Slice(p,
	// 			func(i, j int) bool {
	// 				return p[i].Updated_time.After(p[j].Updated_time)
	// 			})
	// 	case 1: //old
	// 		sort.Slice(p,
	// 			func(i, j int) bool {
	// 				return p[i].Updated_time.Before(p[j].Updated_time)
	// 			})
	// 	}
	// }
	// p:= 55
	p:= serv.Number()
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
		// form := r.PostForm
		// body := form["body"][0]
		// image := form["image"][0] //form is map value is []string of sliceで来るから[0]で取得

		// fmt.Fprintf(w, "フォーム1：\n%v\n", form["say"])
		//         wに描く format string  書き込む内容
		// serv.Create(body, image) //insert to db

		http.Redirect(w, r, "/", http.StatusMovedPermanently) //code 301
	}
}

// func Edit(w http.ResponseWriter, r *http.Request) {
// 	tem, _ := template.ParseFiles("edit.html")
// 	params := r.URL.Query()       //Queryで取得 map
// 	e := serv.Edit(Getid(params)) //return []serv.Vertex
// 	tem.Execute(w, e[0])
// }

// func Update(w http.ResponseWriter, r *http.Request) {
// 	if r.PostFormValue("_method") == "PUT" {
// 		r.ParseForm()

// 		form := r.PostForm
// 		params := r.URL.Query()
// 		id := Getid(params)
// 		body := form["body"][0]
// 		image := form["image"][0]

// 		serv.Update(id, body, image)
// 		http.Redirect(w, r, "/", http.StatusMovedPermanently)
// 	}
// }

// func Delete(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "DELETE" {
// 		params := r.URL.Query()    //to use Getid
// 		serv.Delete(Getid(params)) //use sql delete func

// 		http.Redirect(w, r, "/", http.StatusMovedPermanently)
// 	}
// }

// func Getid(params url.Values) int { //get id from httprequest
// 	var num string
// 	for k, _ := range params { //mapのkeyをget
// 		num = k
// 	}

// 	i, _ := strconv.Atoi(num) //string to int
// 	return i
// }
