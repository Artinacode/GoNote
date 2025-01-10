package main

import "fmt"
//返回值可以命名
//不带参数的return语句返回命名的返回值，称为裸返回
func split(sum int) (x, y int) {
	x =sum *4 /9
	y =sum -x
	return
}

func main (){
	fmt.Println(split(17))
}