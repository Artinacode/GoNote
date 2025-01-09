package main

import (
	"fmt"
	"math/cmplx"
)
// 类似导入语句，变量声明可以用括号 
var (
	ToBe bool	= false
	MaxInt uint64 = 1<<64-1
	z	complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T value: %v\n",ToBe,ToBe)
	fmt.Printf("Type: %T Value: %v\n",MaxInt,MaxInt)
	fmt.Printf("Type: %T Value: %v\n",z ,z)

}