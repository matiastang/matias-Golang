/*
 * @Author: matiastang
 * @Date: 2022-04-18 16:35:58
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-18 17:33:51
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
