/*
 * @Author: matiastang
 * @Date: 2022-04-18 17:00:17
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-18 17:04:58
 * @FilePath: /matias-Golang/learn/src/routers/routers.go
 * @Description: routers
 */
package routers

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	r := gin.New()
	for _, opt := range options {
		opt(r)
	}
	return r
}
