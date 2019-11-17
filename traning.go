package main

import (
	"fmt"
)

func main() {
	//1.11をintに
	// f := 1.11
	// f = f int
	// fmt.Printf("%T\n",f)

	//q2
	//5,6

	//q3
	// m := map[string]int {
	// // m := make(map[string]int)
	// 	"Mike":20,
	// 	"Nncy":24,
	// 	"Messi":30,
	// }
// m["MIke"] = 2330
	//mapを作成するとき
	// makeを使うときには初期値は入れれんかわりに、lenとcapを指定
	//値は後から入れる

	//mapを普通使用の時には、初期値は入れれる
	// m1["MIMI"] = 2121

	// fmt.Printf("%T %v\n",m,m) //=>map[string]int map[MIke:20]
	// fmt.Println(m)
	fmt.Println("答えは？")
	// return1()
	fmt.Printf("%#v",return1())
	fmt.Println("\nreturnは?")
	result := return1()
	result["Mico"] = 230
	fmt.Println(result["Mico"])
	fmt.Println(result)
	//funcを呼ぶ時には()も絶対に必要
	// fmt.Println(number()+2)
}



///fucnの返り値にintじゃなくて、mapを返す、structを返すのはどうか？

func return1()(result map[string]int) {
	//ちゃんと返り値の型指定を厳格にしないとだめ！ mapだけだとng
	//map[string]intと明示しなければならない

	
// func return1() (result int) {
	m := map[string]int{
		// m := make(map[string]int)
			"Mike":20,
			"Nncy":24,
			"Messi":30,
		}
		result = m
		return result
}


func number() int {
return 11
}