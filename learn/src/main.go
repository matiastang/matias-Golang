/*
 * @Author: matiastang
 * @Date: 2022-04-12 19:40:03
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-27 13:51:11
 * @FilePath: /matias-Golang/learn/src/main.go
 * @Description: main
 */
package main

import (
	"fmt"
	"learn/app/get"
	"learn/routers"
)

func init() {
	fmt.Println("learn main init")
}

//go:generate go run main.go
//go:generate go version
func main() {
	// pk1 := painkiller.Aspirin
	// fmt.Println(pk1)
	fmt.Println("learn main main")
	// 加载多个APP的路由配置
	routers.Include(get.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		//...
	}

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello, Gin")
	// })
	// r.GET("/:id", get.GetUserInfo)
	// type msg struct {
	// 	Name    string
	// 	Message string
	// 	Age     int
	// }
	// r.GET("/json", func(c *gin.Context) {
	// 	data := msg{"cl", "消息", 19}
	// 	c.JSON(http.StatusOK, data)
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080

	// test.SliceTest()
	// test.ArrayTest()
	// test.MapTest()
}
