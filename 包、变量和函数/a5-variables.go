package main

import "fmt"

//var 声明一个变量列表
//和函数的参数列表一样，类型在最后
//var 可以出现在包级别或者函数级别
var c,python,java bool = true,false,true
var kk = "no"

var a,b int =1,2
func main()	{
	var i int =2
	fmt.Println(i,c,python,java)
	fmt.Println(i,a,b,kk)

}