package main

import (
	"log"
	_ "log"
	"net/http"

	"./handle"
)

func main() {
	http.HandleFunc("/", handle.Showindex) //index.htmlを表示
	http.Handle("/stylesheet/", http.StripPrefix("/stylesheet/", http.FileServer(http.Dir("stylesheet/"))))
	//Prefix(prefix string, h handler)
	//(prefix stringはurlpathで受け取った、requestを削除し,handler(後ろの引数)をinvokeしている)
	//cssを反映する

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
