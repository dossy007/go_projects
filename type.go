package main
import "fmt"
func main() {

	type myinteger int
	//type 型名　基になる型
//int型からmyinteger型を宣言
	var i myinteger = 1234

	i+= 1
	fmt.Println(i)
	var in int = int(i)
	var sr string = string(in)
	fmt.Println(sr)

	var b []byte = []byte("ab")
	//stringからスライス型（配列）に変換
	fmt.Println(b)


	type mm struct {
		a int
		b int
	}
}

//型の宣言は何の型で使うかの集合体
//一々変数型宣言宣言しなくても良いメリットがあるのではないか？

