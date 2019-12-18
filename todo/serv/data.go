package serv

import (
	"log"
	"time"
  "fmt"
	// "../database"
	"github.com/dossy007/go_projects/todo/database"
)

type Vertex struct { //for save
	Id           int
	Body         string
	Image        string
	Created_time time.Time
	Updated_time time.Time
}

type Stringtime struct { //for show date
	Id int
	Body string
	Image string
	Create_time string
	Updated_time string
}

const time_layout = "2006-01-02 15:04:05"

	
func Connected() []Vertex { //2重slice 全件取得
	db := database.ConnectDB()
	defer db.Close()

	//sql
	rows, err := db.Query(`SELECT id,body,image,created_at,updated_at FROM posts`)
	if err != nil {
		log.Fatal(err)
	}

	var sli []Vertex
	var v1 Vertex //structをv1で使用する宣言

	for rows.Next() { //nextはscanを使う為

		if err := rows.Scan(&v1.Id, &v1.Body, &v1.Image, &v1.Created_time, &v1.Updated_time); err != nil {
			log.Fatal(err)
		}
	
		sli = append(sli, v1)
	}
	fmt.Println(sli)
	return sli
}

func Create(body string, image string) {
	db := database.ConnectDB()
	defer db.Close()

	create_time := time.Now() //time.Time
	rows, err := db.Prepare("INSERT INTO posts(body,image,created_at,updated_at) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	rows.Exec(body, image, create_time, create_time)
	// Exec()にプリペアードステートメントを指定してSQLを実行する
}

func Edit(id int) []Vertex {
	db := database.ConnectDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM posts WHERE id = ?", id)
	//where分は配列で取得しよる
	if err != nil {
		log.Fatal(err)
	}
	
	var sli []Vertex  // sliceをVertexで定義
	var v1 Vertex     //structをv1で使用する宣言
	for rows.Next() { //next is to use scan
		if err := rows.Scan(&v1.Id, &v1.Body, &v1.Image, &v1.Created_time, &v1.Updated_time); err != nil {
			log.Fatal(err)
		}
		sli = append(sli, v1) // [{v1}]
	}
	return sli
}

func Update(id int, body string, image string) {
	db := database.ConnectDB()
	defer db.Close()
	update_time := time.Now()

	rows, err := db.Prepare("UPDATE posts SET body =?,image =?,updated_at = ? WHERE id = ?")

	if err != nil {
		log.Fatal(err)

	}
	rows.Exec(body, image, update_time, id)
}

func Delete(id int) {
	db := database.ConnectDB()
	defer db.Close()
	rows, err := db.Prepare("DELETE FROM posts WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	rows.Exec(id)
}


func Changelayout(t []Vertex)[]Stringtime { // for use time(2006-01-02 15:04:05), serv []Vertex ,return []Stringtime
	var sli []Stringtime
	for _, v := range t { //get map_key
		stime := v.Created_time.Format(time_layout)
		v1 := Stringtime{v.Id,v.Body,v.Image,stime,stime} 
		sli = append(sli, v1)
	}
	return sli
}