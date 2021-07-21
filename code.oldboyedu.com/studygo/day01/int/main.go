package main

import "fmt"

func main() {
	var i1 = 101

	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1)
	fmt.Printf("%o\n", i1)
	fmt.Printf("%x\n", i1)

	i3 := 0x3141312
	fmt.Printf("%d\n", i3)
	//查看变量类型
	fmt.Printf("%T\n", i3)

	i4 := int8(9)

	fmt.Printf("%T\n", i4)
}
