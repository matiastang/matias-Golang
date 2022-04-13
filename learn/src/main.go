/*
 * @Author: matiastang
 * @Date: 2022-04-12 19:40:03
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-13 15:12:59
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
}
