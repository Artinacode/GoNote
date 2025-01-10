package main

import "fmt"
//去掉分号，for 就是c中的while的作用。
func main(){
	sum := 1
	for sum <1000{
		sum += sum
	}
	fmt.Println(sum)
}

//省略循环条件，无限循环
// for {
//}