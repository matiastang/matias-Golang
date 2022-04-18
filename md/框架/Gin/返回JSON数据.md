<!--
 * @Author: matiastang
 * @Date: 2022-04-18 17:13:02
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-18 17:19:01
 * @FilePath: /matias-Golang/md/框架/Gin/返回JSON数据.md
 * @Description: 返回JSON数据
-->
# 返回JSON数据

## 用map作序列化
```go
package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func main() {
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context) {
		//自定义map
		data := map[string]interface{}{
			"name":    "小王子",
			"message": "helloworld",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run()
}
```
```go
package main

import(
	"github.com/gin-gonic/gin"
)
func main(){
	r := gin.Default()
	r.GET("/json",func(c *gin.Context){
		c.JSON(200,gin.H{
			"name":"cl",
			"message":"kkk",
			"age":"19",
		})
	})
	r.Run()
}
```
其中`gin.H`是一种生成`map[string]interface{}`的快捷方式

## 返回struct作序列化
```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main(){
	r := gin.Default()

	type msg struct{
		Name string
		Message string
		Age int
	}
	r.GET("/struct_json",func(c *gin.Context){
		data := msg{"cl","消息",19 }
		c.JSON(http.StatusOK,data )
	})
	r.Run()
}
```
**注意**结构体的字段必须是以大写字母开头(对外暴露),否则接收不到
如果实在想返回首字母为小写,可以使用标签来解决
```go
type msg struct{
	Name string `json:"name"`
	Message string
	Age int
}
```