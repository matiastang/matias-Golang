/*
 * @Author: matiastang
 * @Date: 2022-04-12 19:35:48
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-13 14:29:11
 * @FilePath: /matias-Golang/learn/src/test/main.go
 * @Description:
 */
package main

import (
	"fmt"
	data "test/data"
)

func init() {
	fmt.Println("test main init")
}

func main() {
	fmt.Println("test main main")
	data.SliceTest()
}
