package main

import "fmt"
//函数内部。可以使用：=短赋值语句来代替具有隐式类型的var声明
func main() {
	var i,j int = 1,2
	k :=3
	c,python,java := true,false,"no"
	fmt.Println(i,j,k,c,python,java)
}