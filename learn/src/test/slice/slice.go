/*
 * @Author: matiastang
 * @Date: 2022-04-12 19:50:14
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-20 16:18:11
 * @FilePath: /matias-Golang/learn/src/test/slice/slice.go
 * @Description: slice 切片
 */
package main

/*
切片(slice)本身不是动态数组或动态指针。只是它内部采用数组存储数据，当数组长度达到数组容量时，会进行动态扩容。
切片(slice)可以动态改变长度。数组(array)长度是固定的。

切片的很多功能都是由运行时实现的，无论是初始化切片，还是对切片进行追加或扩容都需要运行时的支持，需要注意的是在遇到大切片扩容或者复制时可能会发生大规模的内存拷贝，一定要减少类似操作避免影响程序的性能。
*/

import "fmt"

func init() {
	fmt.Println("slice init")
}

func main() {
	var arr []int64
	// arr = append(arr, 1, 2, 3, 4, 5)
	arr = append(arr, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
}

/*
5
6
*/

// func SliceTest() {
// 	// slice初始化方式一
// 	a := make([]int, 5)
// 	fmt.Println(a)
// 	a = append(a, 6)
// 	fmt.Println(a)
// 	// slice初始化方式二
// 	var b []int
// 	fmt.Println(b)
// 	b = append(b, 1)
// 	fmt.Println(b)
// }
