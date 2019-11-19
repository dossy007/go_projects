package serv

import (
	"log"

	"../database"
)

func Posts() [][]string { //2重slice 全件取得
	db := database.ConnectDB()
	defer db.Close()

	//sql
	rows, err := db.Query(`SELECT id,body,image FROM posts`)
	if err != nil {
		log.Fatal(err)
	}

	// double_s := [][]string
	var double_s [][]string
	// 2重sliceをstringで定義

	for rows.Next() { //nextはscanを使う為

		var id int
		var body string
		var image string

		if err := rows.Scan(&id, &body, &image); err != nil {
			log.Fatal(err)
		}
		// rows.Scan(&id,&body,&image)

		d := []string{body, image} //sliceを定義
		double_s = append(double_s, d)
	}
	return double_s
}

//[][]stringではなく
//[]Vertexを返すように修正
