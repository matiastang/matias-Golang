/*
 * @Author: matiastang
 * @Date: 2022-04-18 17:01:55
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-18 17:04:11
 * @FilePath: /matias-Golang/learn/src/app/get/router.go
 * @Description: router
 */
package get

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/userinfo/:id", GetUserInfo)
}
