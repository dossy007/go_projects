package main

import (
	"log"
	_ "log"
	"net/http"

	"./handle"
)

func main() {
	http.HandleFunc("/", handle.Showindex) //index.htmlを表示
	// http.HandleFunc関数を使い、第1引数にパスを、第2引数に処理を指定することでリクエストを処理できるようになります。

	http.HandleFunc("/new", handle.New)
	http.HandleFunc("/create", handle.Create)
	http.HandleFunc("/delete", handle.Delete)
	http.Handle("/stylesheet/", http.StripPrefix("/stylesheet/", http.FileServer(http.Dir("stylesheet/"))))
	//Prefix(prefix string, h handler)
	//(prefix stringはurlpathで受け取った、requestを削除し,handler(後ろの引数)をinvokeしている)
	//cssを反映する

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
