package main

import (
	"fmt"
)

func main () {
	sli := make([]int,0)
	sli = append(sli,22,23,24)
	// fmt.Println(len(sli))

	for num := range sli{
		fmt.Println("番号",num)
	}
}