package main
import "fmt"
import "unicode/utf8"
func main(){
	// var i int = 12345
	// var i64 int64 = int64(i)

	// fmt.Println(i64)
	var str string
	str = "aa"
	str = str + "bb"
	str+= "ご"
	var ss string = "ご"
	fmt.Println(len(ss))
	fmt.Println(str+"\nlen:",utf8.RuneCountInString(str))
}
