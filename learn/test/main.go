/*
 * @Author: matiastang
 * @Date: 2022-04-19 19:46:34
 * @LastEditors: matiastang
 * @LastEditTime: 2024-09-04 17:51:49
 * @FilePath: /matias-Golang/learn/test/main.go
 * @Description:
 */
/*
* go tool compile -m ./main.go
* 编译成可执行文件，生成main.o
 */
package main

// import "fmt"

// func init() {
// 	fmt.Println("learn main init")
// }

func main() {
	// fmt.Println("learn main main")

	data := make(map[string][]string)
	data["key"] = []string{"value1"}
}
