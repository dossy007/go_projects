package main
import (
	"fmt"
)

func main(){
	//ここではmapを主に使用する

	// m := map[string]int{"apple":159,"banana":123} //map[type]type
	// fmt.Println(m)//map自体
	// fmt.Println(m["apple"])
	// fmt.Println(m)

	// v, ok := map[key]

	// v,ok := m["apple"] //vには値が okにはbool値が
	// fmt.Println(v)
	// fmt.Println(ok)

	// mm := map[string]int{}
	// mm["apple"] = 123
	// mm["banana"] = 1221
	// fmt.Println(mm["banana"])
	// // _, ok := mm["apple"]
	// // fmt.Println(ok)

	// // v,_ := mm["apple"]
	// _, ok := mm["apple"]
	// fmt.Println(ok)

	// delete(mm,"banana")
	// small := func(mm map[string]int, k string)bool {
		// return (mm[k ]< 200)
	// }
	// delete_if(mm,small)
	m := map[string]int{"apple": 150, "banana": 300}
	// delete(m, "banana")

	// 存在しないときはエラーを表示
	f := func(k string) {
			fmt.Printf("%s not found\n", k)
	}
	delete_if_exist(m, "banana", f) // => "banana not found"
	fmt.Println(m)            
	      // => "map[apple:150]"

	// 200より小さい値を持つエントリを削除
	// m = map[string]int{"apple": 150, "banana": 300, "lemon": 400}
	// f_small := func(m map[string]int, k string) bool {
	// 		return (m[k] < 200)
	// }
	// delete_if(m, f_small)
	// fmt.Println(m) // => "map[banana:300 lemon:400]"

}





// 指定したキーが存在しなければf()を実行
func delete_if_exist(m map[string]int, k string, f func(string)) int {
	v, ok := m[k]

	if ok {
			delete(m, k)
			return v
	} else {
			f(k)
			return 0
	}
}

// func_if()がtrueの場合は対象を消す
func delete_if(
	m map[string]int,
	func_if func(map[string]int, string) bool,
) {
	for k := range m {
			ok := func_if(m, k)
			if ok {
					delete(m, k)
			} else {
			}
	}
	

	

}