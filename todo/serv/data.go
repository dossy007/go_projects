package serv

import (
	"fmt"
	"log"

	"../database"
)

type Vertex struct {
	Id    int
	Body  string
	Image string
}

func Connected() []Vertex { //2重slice 全件取得
	db := database.ConnectDB()
	defer db.Close()

	//sql
	rows, err := db.Query(`SELECT id,body,image FROM posts`)
	if err != nil {
		log.Fatal(err)
	}

	// double_s := [][]string
	// var double_s [][]string
	var sli []Vertex
	// 2重sliceをstringで定義
	var v1 Vertex //structをv1で使用する宣言

	for rows.Next() { //nextはscanを使う為

		if err := rows.Scan(&v1.Id, &v1.Body, &v1.Image); err != nil {
			log.Fatal(err)
		}
		// rows.Scan(&id,&body,&image)
		// fmt.Println(v1.Id)
		// d := []string{v1.Body, v1.Image} //sliceを定義
		sli = append(sli, v1)
		// fmt.Println(sli)
	}
	return sli
}

func Posts(body string, image string) {
	db := database.ConnectDB()
	defer db.Close()
	fmt.Println(body, image)
	rows, err := db.Prepare("INSERT INTO posts(body,image) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	rows.Exec(body, image)
	// Exec()にプリペアードステートメントを指定してSQLを実行する
}
