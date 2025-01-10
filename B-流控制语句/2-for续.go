package main

import "fmt"
//init 和 post 是可选的。
func main(){
	sum := 1
	for ;sum <1000 ;{
		sum += sum
	}
	fmt.Println(sum)
}