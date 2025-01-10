package main

import (
	"fmt"
)
//init 语句
// 条件表达式 
//post语句 

// {}始终需要
func main() {
	sum:=0
	for i := 0;i < 10;i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
}