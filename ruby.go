package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"log"
)

func input()(result string) {
	//scannerを使う
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	result = scanner.Text()
	return result
}

type R1 struct {
	title string
	genre string
	review string
}



func write_review() map[string]string{

	fmt.Println("本のタイトルを入力")
	title := input()
	fmt.Println("本のジャンルを入力")
	genre := input()
	fmt.Println("本の感想を入力")
	review := input()

	r1 := map[string]string{
			"title":title,
			"genre":genre,
			"review":review,
		}


	return r1

}


func read_review(s []map[string]string) {
	for num := range s {
		fmt.Println(num,"titleは",s[num]["title"])
	}
	fmt.Println("みたい番号を取得")
	number := scan()

	if len(s) <= number { //打ち込んだnumがlength以上であれば
		log.Fatal("存在しない番号やんか、あかんで")
	}
	display:= s[number]

	fmt.Println("タイトルは ",display["title"],"\nジャンルは",display["genre"],"\n感想は",display["review"])
}


func scan()(conv_num int) { //打ち込んだ文字を取得しintに変換
	scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() //入力待ち
		num := scanner.Text()
		
		conv_num,_	= strconv.Atoi(num)
		return conv_num
}


func main() {
	var slice []map[string]string
	forbun_label: //labelをfor文に対して宣言
	//forの中にswitchを入れ子にしているのでlabelを宣言すべき

	for {

		fmt.Println("番号を選んでください!")
		fmt.Println("[1]レビューを書く\n[2]レビューをみる\n[3]アプリ終了")

		conv_num := scan()
		switch conv_num {
		case 1:
			fmt.Println("1を選んだね")
			slice = append(slice, write_review())
		case 2:
			fmt.Println("2を選んだね")
			read_review(slice)
	  case 3:
			fmt.Println("3を選んだね")
			break forbun_label  //labelのbreak
		default:
			fmt.Println("もう一度選べや")
	  }
  }
}