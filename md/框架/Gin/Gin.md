<!--
 * @Author: matiastang
 * @Date: 2022-04-18 15:45:34
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-18 17:29:56
 * @FilePath: /matias-Golang/md/框架/Gin/Gin.md
 * @Description: Gin
-->
# Gin

[Go Gin 入门教程](https://zhuanlan.zhihu.com/p/375688504)

## 介绍

Gin 是使用 Go/golang 语言实现的 HTTP Web 框架。接口简洁，性能极高。截止 1.4.0 版本，包含测试代码，仅14K，其中测试代码 9K 左右，也就是说框架源码仅 5K 左右。

$ find . -name "*_test.go" | xargs cat | wc -l
8657
$ find . -name "*.go" | xargs cat | wc -l
14115
## Gin 特性
快速：路由不使用反射，基于Radix树，内存占用少。
中间件：HTTP请求，可先经过一系列中间件处理，例如：Logger，Authorization，GZIP等。这个特性和 NodeJs 的 Koa 框架很像。中间件机制也极大地提高了框架的可扩展性。
异常处理：服务始终可用，不会宕机。Gin 可以捕获 panic，并恢复。而且有极为便利的机制处理HTTP请求过程中发生的错误。
JSON：Gin可以解析并验证请求的JSON。这个特性对Restful API的开发尤其有用。
路由分组：例如将需要授权和不需要授权的API分组，不同版本的API分组。而且分组可嵌套，且性能不受影响。
渲染内置：原生支持JSON，XML和HTML的渲染。


## 安装

```
go get -u -v github.com/gin-gonic/gin
```
-v：打印出被构建的代码包的名字
-u：已存在相关的代码包，强行更新代码包及其依赖包

## 使用

```
# go.mod添加require github.com/gin-gonic/gin v1.7.7
go mod edit -require github.com/gin-gonic/gin@latest
# 更新依赖（生成go.sum）
go mod tidy
```