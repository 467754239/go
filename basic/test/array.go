package main

import (
	"fmt"
)

func main() {

	var a [2]int // 长度为2的int型数组
	var b [1]int // 长度为1的int型数组

	c := [2]int{1, 1}
	d := [2]int{1}

	e := [20]int{19: 1}             //将索引为19的元素的值赋值为1
	f := [...]int{1, 2, 3, 4, 5}    //用三个点来代替数组的长度，自行判断.
	g := [...]int{0: 1, 1: 2, 2: 3} //通过数组索引赋值的方式.
	h := [...]int{19: 1}

	// 指向数组的指针
	j := [...]int{99: 1} // 定义一个数组j
	var p *[100]int = &j // 取j的地址
	fmt.Println(p)       //表示取数组的地址

	// 使用new关键之返回的是指向数组的指针
	pp := new([10]int)
	fmt.Println("pp >>>", pp)

	// 指针数组
	x, y := 1, 2
	m := [...]*int{&x, &y}
	fmt.Println(m) // 内存地址

	fmt.Println(a, b, c, d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	// 数组比较
	// 数组的比较一定是要相同类型才能比较，比如长度不同就无法比较.
	// 长度也是数组类型的一部分
	aa := [2]int{1, 2}
	bb := [2]int{1, 2}
	cc := [2]int{1, 3}
	fmt.Println(aa == bb)
	fmt.Println(aa == cc)

	/*
	   不管是数组本身还是指向数组的指针都可以使用variable[index] = value的方式操作
	*/
	aaa := [10]int{} //长度为10的int型数组
	aaa[1] = 100     //数组aaa中索引为9的值赋值为100
	bbb := new([10]int)
	bbb[1] = 2
	fmt.Println(aaa, bbb)

	// 多维数组
	// 数组的元素有2个，每一个元素都是长度为3的int型数组.
	aaaa := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}

	aaab := [2][3]int{
		{1: 1},
		{2: 3}}
	fmt.Println(aaaa)
	fmt.Println(aaab)

	// 冒泡排序
	abc := [...]int{5, 2, 6, 3, 9}
	fmt.Println(abc)

	num := len(abc)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if abc[i] < abc[j] {
				temp := abc[i]
				abc[i] = abc[j]
				abc[j] = temp
			}
		}
	}
	fmt.Println(abc)

}
