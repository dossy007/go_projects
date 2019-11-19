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
}

func Showindex(w http.ResponseWriter, r *http.Request) {
	tem, _ := template.ParseFiles("index.html")

	p := serv.Posts()
	//serv packageのPosts funcでdbの情報を受け取っている
	// fmt.Println(r.Method)

	var pp []Vertex //slice Vertex型を指定
	v := Vertex{Id: 1, Name: "https://d1f5hsy4d47upe.cloudfront.net/3c/3cfb747bdc87c4dacd090210a0b95d5e_t.jpeg"}
	v2 := Vertex{Id: 2, Name: "https://fotopus.com/neko/images/top/neko_subject_04.jpg"}
	pp = append(pp, v, v2)
	// fmt.Println(p[0])
	// v1 := Vertex{Id: 3, Name: "3ですよ〜〜"}
	// fmt.Println(p[0][0])
	// fmt.Printf("%T\n %#v", p[1], p[1])
	fmt.Printf("%T\n", p)
	tem.Execute(w, pp)
	//execute is template to act and http.RequestWriter に書き出す
}
