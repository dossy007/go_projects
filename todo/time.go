package main

import (
	"fmt"
	"sort"
	"time"
)

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

var past = time.Date(2010, time.May, 18, 23, 0, 0, 0, time.Now().Location())
var present = time.Now()
var future = time.Now().Add(24 * time.Hour)

var dateSlice timeSlice = []time.Time{present, future, past}
var times = []time.Time{present, future, past}

//今日、明日,過去の3つがtime型のsliceで管理

func main() {

	// fmt.Println("Past : ", past)
	// fmt.Println("Present : ", present)
	// fmt.Println("Future : ", future)

	// fmt.Println("Before sorting : ", dateSlice)

	// sort.Sort(dateSlice)

	// fmt.Println("\nAfter sorting : ", dateSlice)

	// sort.Sort(sort.Reverse(dateSlice))

	// fmt.Println("\nAfter REVERSE sorting : ", dateSlice)

	fmt.Println(times)
	fmt.Println("--------------")

	sort.Slice(times,
		func(i, j int) bool {
			fmt.Println(i, j)
			return times[i].Before(times[j])
			// return time[i].create_time < times[j].create_time
			//バブルsortかな 比べて小さいものを最初に持ってくる
			//1と2を比べ,1と３を比べ1と4を比べる
			//trueの場合にsliceに入れるの繰り返し
		})

	fmt.Println(times)
}
