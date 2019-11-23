package main

import "fmt"

func main() {
	fmt.Println("Hello Assemble time")
}

//GOOS=js GOARCH=wasm go build -o main.wasm
//このコマンドでmain.goのwasm fileを作成している
// そしてその時点のコードが実行される
