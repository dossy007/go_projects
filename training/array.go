package main
import (
	"fmt"
)

func main(){


	// var array [1]byte
	//変数定義 変数名 [配列の要素数] 要素型
	var array1 []*int
	var array2 [8][3]int64
	var array3 [1]struct{x,y,z int}
	// fmt.Println(array)
	// fmt.Println(len(array))
	fmt.Println(array1)
	fmt.Println(len(array1))
	fmt.Println(len(array2))
	fmt.Println(array2)	
	fmt.Println(array3)

	// var array[3]string

	// array[0] = "日曜"
	// array[1] = "月曜"
	// array[2] = "火曜"

	// fmt.Println(array[0])

	// for i := 0; i < len(array); i++ {
	// 	fmt.Println(array[i])
	// }

	// array5 := [6]int{1,2,3,4}

	// fmt.Println(array5[3])
	// array6 := []int{1,2}
	// fmt.Println(array6)

	//配列は[]のとき{}で中身を入れれる。入れた分だけlengthが定義される
	//[3]でlengthを定義した時{}は4以上は無理、2以下の場合定義していないものは初期値0で定義される

	//goの配列は最初に配列の要素数を定義してしまうので、追加という概念が無い
	//だから基本的にはsliceを使用するものかと思うよ。


	//slice 要素数は未定義で空にする
	// rr := []int{1,2,3}

	// r1 := rr[:] //番号2未満をslice
	// fmt.Println(r1)

	// r1 = append(r1,5) //sliceに5をappend sliceの末尾に追加
	// r1[0] = 78 //sliceのupdateはこれ！！！
	// fmt.Println(len(r1))
	//sliceの初期値はnil 配列は0

	// sli := make([]int,5,5) //make関数でsliceを定義すれば初期値は0
	// make([]Type,length,capacity)

	//capacity sliceをsliceした時の上限

	slice := []int{0,1,2,3}
	sss := []int{5}
	fmt.Println(sss)
	// slice = append(slice,1212) //追加
	// // slice = slice[:len(slice)-1]
	// // fmt.Println(slice[:4-1]) == slice[:3] //3未満を取得
	// fmt.Println(slice[3])
	// fmt.Println(slice[len(slice)-1]) //末尾取得
	// fmt.Println(slice[:len(slice)-1]) //末尾削除 末尾未満になる



	fmt.Println("-----------------------")

	
	// slice , slice[0] = append(slice[:1],slice[0:]...),"0" //先頭に追加
	fmt.Println(slice)
	pos := 2
	// slice = append(slice[:pos+1])
	slice = append(slice[:pos+1],slice[pos:]...)

	// slice[pos] = 4343
	fmt.Println(slice)
	// capacity := []int{1,2}
	// fmt.Println(capacity)
	// fmt.Println(cap(capacity))  //capは2

	// capacity2 := capacity[:1]
	// capacity2 = append(capacity2)
	// fmt.Println(capacity2)
	// fmt.Println(cap(capacity2))

	// values := []int{0,1,2,3,4}

	// s1 := values[1:3]//1以上3未満
	// fmt.Println(s1)
	// fmt.Println("s1のcap: ",cap(s1))

	// s2 := s1[1:4]
	// fmt.Println(s2)
	// fmt.Println(cap(s2))	

	// s2 = append(s2,44)

	// fmt.Println(cap(s2))
}