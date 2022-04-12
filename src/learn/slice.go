/*
 * @Author: matiastang
 * @Date: 2022-04-12 19:50:14
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-12 20:00:49
 * @FilePath: /matias-Golang/src/learn/slice.go
 * @Description: slice切片
 */
package slice

/*
切片(slice)本身不是动态数组或动态指针。只是它内部采用数组存储数据，当数组长度达到数组容量时，会进行动态扩容。
切片(slice)可以动态改变长度。数组(array)长度是固定的。
*/

import "fmt"

func test() {
	// slice初始化方式一
	a := make([]int, 5)
	fmt.Println(a)
	a = append(a, 6)
	fmt.Println(a)
	// slice初始化方式二
	var b []int
	fmt.Println(b)
	b = append(b, 1)
	fmt.Println(b)
}
