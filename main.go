package main

import "fmt"

var tiangan  = [10]string{"甲","乙","丙","丁","戊","己","庚","辛","壬","癸"}
var dizhi = [12]string{"子","丑","寅","卯","辰","巳","午","未","申","酉","戌","亥"}

var wuxing1 = [10]string{"木","木", "火","火", "土", "土", "金", "金", "水","水",}
var wuxing2 = [12]string{"水","土","木","木","土","火","火","土","金","金","土","水"}
func main() {
	var tian = 2
	var di = 10
	var cnt = 60
	for cnt > 0 {
		fmt.Printf(tiangan[tian] + dizhi[di])
		fmt.Printf(" ")
		tian--
		if tian < 0 {
			tian = 9
		}
		di--;
		if di <0 {
			di = 11
		}
		cnt--
		if cnt%10 == 0{
			fmt.Println()
		}
	}
	fmt.Println()
	 tian = 2
	 di = 10
	 cnt = 60
	for cnt > 0 {
		fmt.Printf(wuxing1[tian] + wuxing2[di])
		fmt.Printf(" ")
		tian--
		if tian < 0 {
			tian = 9
		}
		di--;
		if di <0 {
			di = 11
		}
		cnt--
		if cnt%10 == 0{
			fmt.Println()
		}
	}
}
