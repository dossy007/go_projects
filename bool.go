package main
import "fmt"
func main(){
	var b bool
	//bool型の変数 bを定義
	//bool型にはリテラル変数 true falseしか入りません

	b = true
	b = false
	// fmt.Println(b)

	b = true || false
	//bにはtrueもしくはfalseが入っているのでtrueが返される

	fmt.Println(b)
}
