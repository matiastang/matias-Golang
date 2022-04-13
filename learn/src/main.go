/*
 * @Author: matiastang
 * @Date: 2022-04-12 19:40:03
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-13 19:22:21
 * @FilePath: /matias-Golang/learn/src/main.go
 * @Description: main
 */
package main

import (
	"fmt"
	test "learn/test/data"
)

func init() {
	fmt.Println("learn main init")
}

func main() {
	fmt.Println("learn main main")
	test.SliceTest()
	test.ArrayTest()
	test.MapTest()
}
