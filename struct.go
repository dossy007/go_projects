package main
import ("fmt")


type Vertex struct {
	x int
	y int
	s string
}

func main() {
	fmt.Println("111")
	//構造体を使用

	var x struct {
		i1,i2 int //宣言
		y int //値は0
		yy int
	}
	x.i1 = 1
	x.i2 = 22

	fmt.Println(x.i1)

	// v2

	v2 := Vertex{x: 1}
	v2.s = "wwwwww"
	fmt.Println(v2.s)
	fmt.Printf("%T %v\n",v2.s,v2.s)
}