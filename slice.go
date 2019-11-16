package main
import (
	"fmt"
)

func main() {
	// sliceから取りたい値を取得する方法を探る
	//まずは、intで値を入れて取り出したい値を取得する

	// slice := []int
	slice := make([]int,1)
	slice = append(slice,22)
	slice[0] = 11  //update
	slice = append(slice,33)
	// fmt.Println(slice)

	num := 2


	//funcにintの値を渡す処理
	// fmt.Printf("%T\n",slice[num])
	//sliceは []intの型
	//slice[0]はintの型


	sliced(slice[num])
	fmt.Println(slice[num])
	// fmt.Println(slice2)

	slice2(&slice[num]) //point渡しにする
	fmt.Println(slice[num])

}

func sliced (x int) {
	x = 222
	fmt.Println(x)
	//値渡し
}


func slice2 (x *int) {
	*x = 111
	fmt.Println(*x)
}