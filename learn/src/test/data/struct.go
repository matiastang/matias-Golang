/*
 * @Author: matiastang
 * @Date: 2022-04-13 19:23:56
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-13 19:23:56
 * @FilePath: /matias-Golang/learn/src/test/data/struct.go
 * @Description: struct
 */
package data

import "fmt"

func init() {
	fmt.Println("map init")
}

type user struct {
	name string
	age  int
}

// 结构体user Read方法
func (u *user) Read() string {
	return fmt.Sprintf("%s 已经 %d 岁了", u.name, u.age)
}

func StructTest() {
	// 初始化
	u := &user{
		name: "迈莫coding",
		age:  1,
	}
	fmt.Println(u.name, "已经", u.age, "岁了")
	// 调用结构体user的 Read方法
	fmt.Println(u.Read()) // 迈莫coding 已经 1 岁
}
