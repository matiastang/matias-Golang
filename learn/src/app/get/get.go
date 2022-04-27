/*
 * @Author: matiastang
 * @Date: 2022-04-18 16:35:58
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-27 17:57:02
 * @FilePath: /matias-Golang/learn/src/app/get/get.go
 * @Description: get
 */
package get

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Id   string `json:"userId"`
	Name string
}

func GetUserInfo(c *gin.Context) {
	userId := c.Param("id")
	info := UserInfo{
		userId,
		"name_" + userId,
	}
	fmt.Println(info)
	c.JSON(http.StatusOK, info)
}

func Get_Info(c *gin.Context) {
	c.String(http.StatusOK, `“_”是特殊标识符，用来忽略结果`)
}

func GetIotaInfo(c *gin.Context) {
	c.String(http.StatusOK, `iota是go语言的常量计数器，只能在常量的表达式中使用。 iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。`)
}

// 遍历字符串
func traversalString(s string) []string {
	var arr []string
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
		arr = append(arr, fmt.Sprintln(s[i], s[i]))
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
		arr = append(arr, fmt.Sprintln(r, r))
	}
	fmt.Println()
	return arr
}

func GetString(c *gin.Context) {
	s := c.Param("str")
	arr := traversalString(s)
	c.JSON(http.StatusOK, arr)
}
