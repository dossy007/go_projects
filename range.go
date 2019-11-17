package main

import (
	"fmt"
)

func main () {
	sli := make([]int,0)
	sli = append(sli,22,23,24,25)

	for num, block := range sli{
		fmt.Println("番号",num,"slice-number",block)
	}
	//explanation
	//rangeはrubyで言う所のeachみたいなもん
	//rangeで回せるのは、配列、slice,文字列,map,channelのみ。
	//for文でrangeを用いると配列などから要素を1つずつ取り出してroopできる。

	//for num , block := range 回したい要素 {
		// fmt.Println(num,block)
		//numが0から要素の数分のlength
		//blockが１つずつの要素
	// }


}