package main

import (
	"fmt"
	"strconv"
	// "os"
	// "bufio"
)

func main(){
	fmt.Println("最初はグーじゃんけん?"+"\n[1]ぐー[2]チョキ[3]パー")
	// number := 3
	var number int64 = 3
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	//Scan methodを使って入力待ちを実現

	// if scanner.Text() == 1{
	// }

	if number == 1{
	fmt.Println("勝利")
	}else if number == 2{
	fmt.Println("負け")
	}else{
		 // var num string = string(number)
	fmt.Println("アイコ")
  }
  var st_num = strconv.FormatInt(number,10)
  //formatIntはint64型でないとだめ！ 第二引数には進数を入れる
  fmt.Println(st_num+"の数字が選ばれました")
}
