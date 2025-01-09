package main

import "fmt"
//类型在变量名之后
func add(x int,y int)int {
	return x + y
}

//两个或多个连续命名函数参数共享一个类型时，除了最后一个，前面的省略类型
func add2(x ,y int)int {
	return x + y
}

func main()	{
	fmt.Println(add(42,13))
	fmt.Println(add2(42,13))

}