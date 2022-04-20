/*
 * @Author: matiastang
 * @Date: 2022-04-20 15:24:47
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-20 17:41:56
 * @FilePath: /matias-Golang/learn/src/test/data/slice_test.go
 * @Description: slice
 */
package main

import (
	"fmt"
	"testing"
)

func Benchmark_ar_len_4_test(b *testing.B) {

	// len 和 cap
	var c int
	for i := 0; i < b.N; i++ {
		arr := [5]int{1, 2, 3, 4, 5}
		// arr := [4]int{1, 2, 3, 4}
		for i := 0; i < len(arr); i++ {
			c = arr[i]
		}
		arr = [5]int{6, 7, 8, 9, 10}
		// arr = [4]int{5, 6, 7, 8}
		for i := 0; i < len(arr); i++ {
			c = arr[i]
		}
	}
	fmt.Println(c)
}
