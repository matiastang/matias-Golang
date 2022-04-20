/*
 * @Author: matiastang
 * @Date: 2022-04-20 10:24:35
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-20 11:11:03
 * @FilePath: /matias-Golang/learn/src/test/data/array_test.go
 * @Description: array 测试
 */
package main

import (
	"fmt"
	"testing"
)

func Benchmark_arr_len_4_test(b *testing.B) {
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

func Benchmark_arr_len_8_test(b *testing.B) {
	var c int
	for i := 0; i < b.N; i++ {
		arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		// arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
		for i := 0; i < len(arr); i++ {
			c = arr[i]
		}
	}
	fmt.Println(c)
}

/*
go test -v -bench=. array_test.go
goos: darwin
goarch: arm64
Benchmark_arr_len_4_test
Benchmark_arr_len_4_test-8      100000000               10.36 ns/op
Benchmark_arr_len_8_test
Benchmark_arr_len_8_test-8      136603740                8.671 ns/op
PASS
ok      command-line-arguments  3.663s
*/
/*
go test -v -bench=. array_test.go
goos: darwin
goarch: arm64
Benchmark_arr_len_4_test
Benchmark_arr_len_4_test-8      200647575                5.915 ns/op
Benchmark_arr_len_8_test
Benchmark_arr_len_8_test-8      162815412                7.364 ns/op
PASS
ok      command-line-arguments  4.296s
*/
