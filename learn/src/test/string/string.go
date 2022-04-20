/*
 * @Author: matiastang
 * @Date: 2022-04-20 16:27:27
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-20 17:25:36
 * @FilePath: /matias-Golang/learn/src/test/string/string.go
 * @Description: string
 */
package main

/*
Go 语言只是不支持直接修改 string 类型变量的内存空间，我们仍然可以通过在 string 和 []byte 类型之间反复转换实现修改这一目的：

先将这段内存拷贝到堆或者栈上；
将变量的类型转换成 []byte 后并修改字节数据；
将修改后的字节数组转换回 string；
Java、Python 以及很多编程语言的字符串也都是不可变的，这种不可变的特性可以保证我们不会引用到意外发生改变的值，而因为 Go 语言的字符串可以作为哈希的键，所以如果哈希的键是可变的，不仅会增加哈希实现的复杂度，还可能会影响哈希的比较。
因为字符串作为只读的类型，我们并不会直接向字符串直接追加元素改变其本身的内存空间，所有在字符串上的写入操作都是通过拷贝实现的。

字符串的拼接和拷贝

但是在正常情况下，运行时会调用 copy 将输入的多个字符串拷贝到目标字符串所在的内存空间。新的字符串是一片新的内存空间，与原来的字符串也没有任何关联，一旦需要拼接的字符串非常大，拷贝带来的性能损失是无法忽略的。

字符串和字节数组的转换

字符串和 []byte 中的内容虽然一样，但是字符串的内容是只读的，我们不能通过下标或者其他形式改变其中的数据，而 []byte 中的内容是可以读写的。不过无论从哪种类型转换到另一种都需要拷贝数据，而内存拷贝的性能损耗会随着字符串和 []byte 长度的增长而增长。
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
