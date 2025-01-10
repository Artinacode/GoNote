package main

import "fmt"

//左侧不指定显示类型，变量从右侧推断出来
//当右侧包含无类型的数字常量，新的变量取决于常量的精度。
func main()	{
	var i int 
	j := i
	v := 42 //
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("%T %T\n",i,j)
}