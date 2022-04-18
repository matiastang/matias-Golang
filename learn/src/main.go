/*
 * @Author: matiastang
 * @Date: 2022-04-12 19:40:03
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-18 16:12:38
 * @FilePath: /matias-Golang/learn/src/main.go
 * @Description: main
 */
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// test "learn/test/data"
)

// 在 RESTful API 中，使用的主要是以下五种HTTP方法：
// 	GET，表示读取服务器上的资源
// 	POST，表示在服务器上创建资源
// 	PUT,表示更新或者替换服务器上的资源
// 	DELETE，表示删除服务器上的资源
// 	PATCH，表示更新/修改资源的一部分

func init() {
	fmt.Println("learn main init")
}

func main() {
	fmt.Println("learn main main")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	// test.SliceTest()
	// test.ArrayTest()
	// test.MapTest()
}
